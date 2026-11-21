<template>
	<q-page class="row bg-grey-2">
		<q-scroll-area class="col" style="overflow-x: hidden">
			<div class="q-px-lg q-pb-lg q-pt-md">
				<div class="text-h5 q-mb-lg text-weight-bold text-dark">
					Экономическая эффективность
				</div>

				<div class="row q-col-gutter-lg">
					<!-- ЛЕВАЯ КОЛОНКА: Ввод данных -->
					<div class="col-12 col-md-5">
						<q-card
							style="
								border-radius: 16px;
								box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
							"
						>
							<q-card-section>
								<div
									class="text-subtitle1 text-weight-bold text-grey-8 q-mb-md flex items-center"
								>
									<q-icon name="tune" class="q-mr-sm" color="primary" />
									Исходные параметры
								</div>

								<q-form @submit="onCalculate" class="q-gutter-y-md">
									<q-input
										v-model.number="form.downtime_hours_per_month"
										type="number"
										label="Простои оборудования (ч/мес)"
										outlined
										dense
										suffix="ч"
									/>
									<q-input
										v-model.number="form.cost_per_downtime_hour"
										type="number"
										label="Стоимость 1 часа простоя"
										outlined
										dense
										suffix="₽"
									/>
									<q-input
										v-model.number="form.expected_reduction_percent"
										type="number"
										label="Ожидаемое сокращение простоев"
										outlined
										dense
										suffix="%"
									/>

									<q-separator class="q-my-md" />

									<q-input
										v-model.number="form.implementation_cost"
										type="number"
										label="Затраты на внедрение (разово)"
										outlined
										dense
										suffix="₽"
									/>
									<q-input
										v-model.number="form.monthly_operating_cost"
										type="number"
										label="Операционные расходы (в мес)"
										outlined
										dense
										suffix="₽"
									/>

									<div class="q-mt-lg">
										<q-btn
											label="Рассчитать эффективность"
											color="primary"
											icon="calculate"
											type="submit"
											:loading="isCalculating"
											class="full-width"
											size="lg"
											style="border-radius: 8px"
										/>
									</div>
								</q-form>
							</q-card-section>
						</q-card>
					</div>

					<!-- ПРАВАЯ КОЛОНКА: Результаты -->
					<div class="col-12 col-md-7">
						<div
							v-if="!result && !isCalculating"
							class="flex flex-center text-grey-5"
							style="
								height: 100%;
								min-height: 400px;
								border: 2px dashed #ccc;
								border-radius: 16px;
							"
						>
							<div class="text-center">
								<q-icon name="savings" size="56px" class="q-mb-sm" />
								<div class="text-h6">Ожидание расчёта</div>
								<div class="text-body2">Введите параметры и нажмите кнопку</div>
							</div>
						</div>

						<div
							v-if="isCalculating"
							class="flex flex-center"
							style="height: 100%; min-height: 400px"
						>
							<q-spinner-dots color="primary" size="4em" />
						</div>

						<div v-if="result" class="column q-gutter-y-md">
							<!-- Главная карточка: Окупаемость -->
							<q-card
								style="
									border-radius: 16px;
									background: linear-gradient(
										135deg,
										var(--q-primary) 0%,
										#312e81 100%
									);
									color: white;
									box-shadow: 0 8px 24px rgba(79, 70, 229, 0.3);
								"
							>
								<q-card-section class="text-center q-pa-lg">
									<q-icon
										name="schedule"
										size="40px"
										class="q-mb-sm opacity-80"
									/>
									<div class="text-subtitle1 opacity-80">Срок окупаемости</div>
									<div class="text-h1 text-weight-bolder">
										{{
											result.payback_months > 0 ? result.payback_months : '∞'
										}}
										<span class="text-h4">мес.</span>
									</div>
								</q-card-section>
							</q-card>

							<!-- Сетка с цифрами -->
							<div class="row q-col-gutter-md">
								<div class="col-12 col-sm-6">
									<q-card
										class="bg-white"
										style="border-radius: 12px; height: 100%"
									>
										<q-card-section>
											<div class="text-caption text-grey-7">
												Экономия в месяц (чистая)
											</div>
											<div
												class="text-h4 text-positive text-weight-bold q-mt-xs"
											>
												+ {{ formatMoney(result.net_monthly_savings) }}
											</div>
										</q-card-section>
									</q-card>
								</div>
								<div class="col-12 col-sm-6">
									<q-card
										class="bg-white"
										style="border-radius: 12px; height: 100%"
									>
										<q-card-section>
											<div class="text-caption text-grey-7">
												Прибыль за 1-й год
											</div>
											<div
												class="text-h4"
												:class="
													result.first_year_net_benefit > 0
														? 'text-positive'
														: 'text-negative'
												"
											>
												{{ result.first_year_net_benefit > 0 ? '+' : '' }}
												{{ formatMoney(result.first_year_net_benefit) }}
											</div>
										</q-card-section>
									</q-card>
								</div>
							</div>

							<!-- Детализация -->
							<q-card class="bg-white" style="border-radius: 12px">
								<q-card-section>
									<div
										class="text-subtitle2 text-weight-bold text-grey-8 q-mb-md"
									>
										Детализация расчёта
									</div>

									<div class="row justify-between q-py-xs border-bottom">
										<div class="text-grey-7">Валовая экономия на простоях:</div>
										<div class="text-weight-medium">
											{{ formatMoney(result.gross_monthly_savings) }}
										</div>
									</div>
									<div class="row justify-between q-py-xs border-bottom">
										<div class="text-grey-7">Операционные расходы:</div>
										<div class="text-weight-medium text-negative">
											- {{ formatMoney(form.monthly_operating_cost) }}
										</div>
									</div>
									<div class="row justify-between q-py-sm">
										<div class="text-weight-bold text-grey-9">
											Чистая экономия в месяц:
										</div>
										<div class="text-weight-bold text-positive">
											{{ formatMoney(result.net_monthly_savings) }}
										</div>
									</div>
								</q-card-section>
							</q-card>

							<!-- Предупреждение -->
							<q-banner
								class="bg-orange-1 text-orange-9"
								style="border-radius: 8px"
							>
								<template v-slot:avatar>
									<q-icon name="warning" color="orange-8" />
								</template>
								<div class="text-body2" style="line-height: 1.4">
									{{ result.note }}
								</div>
							</q-banner>
						</div>
					</div>
				</div>
			</div>
		</q-scroll-area>
	</q-page>
</template>

<script setup>
import { ref } from 'vue'
import { CalculateEconomics } from '../../bindings/DefectAnalyzer/desktop/core/handlers/economicshandler'

const isCalculating = ref(false)
const result = ref(null)

// Данные по умолчанию (из примера API)
const form = ref({
	downtime_hours_per_month: 16,
	cost_per_downtime_hour: 5000,
	expected_reduction_percent: 25,
	implementation_cost: 60000,
	monthly_operating_cost: 2000,
})

const onCalculate = async () => {
	isCalculating.value = true
	result.value = null
	try {
		// Вызов Go-метода, который делает математику
		result.value = await CalculateEconomics(form.value)
	} catch (error) {
		console.error('Ошибка при расчёте:', error)
	} finally {
		isCalculating.value = false
	}
}

// Форматирование денег (например, 18000 -> 18 000 ₽)
const formatMoney = value => {
	if (value === undefined || value === null) return '0 ₽'
	return new Intl.NumberFormat('ru-RU', {
		style: 'currency',
		currency: 'RUB',
		maximumFractionDigits: 0,
	}).format(value)
}
</script>

<style scoped>
.border-bottom {
	border-bottom: 1px solid #e0e0e0;
}
</style>

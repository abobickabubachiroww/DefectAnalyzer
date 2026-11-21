<template>
	<q-page class="row bg-grey-2">
		<q-scroll-area class="col" style="overflow-x: hidden">
			<div class="q-px-lg q-pb-lg q-pt-md">
				<!-- Заголовок страницы -->
				<div class="text-h6 q-mb-lg text-weight-bold text-grey-9">
					Прогноз партии
				</div>

				<div class="row q-col-gutter-lg">
					<!-- ЛЕВАЯ КОЛОНКА: Форма -->
					<div class="col-12 col-md-8">
						<q-card
							style="
								border-radius: 16px;
								box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
							"
						>
							<q-card-section>
								<div
									class="text-subtitle1 text-weight-bold text-grey-8 q-mb-md"
								>
									Параметры производства
								</div>

								<q-form @submit="onPredict" class="row q-col-gutter-md">
									<!-- Производство -->
									<div class="col-12">
										<div class="text-caption text-grey-6 text-weight-bold">
											ПРОИЗВОДСТВО
										</div>
									</div>

									<div class="col-12 col-sm-4">
										<q-input
											v-model.number="form.ProductionVolume"
											type="number"
											label="Объём производства"
											outlined
											dense
										/>
									</div>
									<div class="col-12 col-sm-4">
										<q-input
											v-model.number="form.ProductionCost"
											type="number"
											label="Себестоимость"
											outlined
											dense
										/>
									</div>
									<div class="col-12 col-sm-4">
										<q-input
											v-model.number="form.WorkerProductivity"
											type="number"
											label="Производительность труда"
											outlined
											dense
										/>
									</div>

									<!-- Поставки -->
									<div class="col-12 q-mt-md">
										<div class="text-caption text-grey-6 text-weight-bold">
											ПОСТАВКИ
										</div>
									</div>

									<div class="col-12 col-sm-4">
										<q-input
											v-model.number="form.SupplierQuality"
											type="number"
											label="Качество поставщиков"
											outlined
											dense
										/>
									</div>
									<div class="col-12 col-sm-4">
										<q-input
											v-model.number="form.DeliveryDelay"
											type="number"
											label="Задержка поставки"
											outlined
											dense
										/>
									</div>
									<div class="col-12 col-sm-4">
										<q-input
											v-model.number="form.InventoryTurnover"
											type="number"
											label="Оборачиваемость запасов"
											outlined
											dense
										/>
									</div>
									<div class="col-12 col-sm-4">
										<q-input
											v-model.number="form.StockoutRate"
											type="number"
											label="Доля дефицита"
											outlined
											dense
										/>
									</div>

									<!-- Обслуживание и ресурсы -->
									<div class="col-12 q-mt-md">
										<div class="text-caption text-grey-6 text-weight-bold">
											ОБСЛУЖИВАНИЕ И РЕСУРСЫ
										</div>
									</div>

									<div class="col-12 col-sm-4">
										<q-input
											v-model.number="form.MaintenanceHours"
											type="number"
											label="Время обслуживания"
											outlined
											dense
										/>
									</div>
									<div class="col-12 col-sm-4">
										<q-input
											v-model.number="form.DowntimePercentage"
											type="number"
											label="Простой оборудования (%)"
											outlined
											dense
										/>
									</div>
									<div class="col-12 col-sm-4">
										<q-input
											v-model.number="form.SafetyIncidents"
											type="number"
											label="Инциденты безопасности"
											outlined
											dense
										/>
									</div>
									<div class="col-12 col-sm-4">
										<q-input
											v-model.number="form.EnergyConsumption"
											type="number"
											label="Потребление энергии"
											outlined
											dense
										/>
									</div>
									<div class="col-12 col-sm-4">
										<q-input
											v-model.number="form.EnergyEfficiency"
											type="number"
											label="Энергоэффективность"
											outlined
											dense
										/>
									</div>

									<!-- Аддитивные процессы -->
									<div class="col-12 q-mt-md">
										<div class="text-caption text-grey-6 text-weight-bold">
											АДДИТИВНЫЕ ПРОЦЕССЫ
										</div>
									</div>

									<div class="col-12 col-sm-4">
										<q-input
											v-model.number="form.AdditiveProcessTime"
											type="number"
											label="Время аддитивного процесса"
											outlined
											dense
										/>
									</div>
									<div class="col-12 col-sm-4">
										<q-input
											v-model.number="form.AdditiveMaterialCost"
											type="number"
											label="Стоимость материала"
											outlined
											dense
										/>
									</div>

									<!-- Кнопка прогноза -->
									<div class="col-12 q-mt-lg flex justify-end">
										<q-btn
											label="Сделать прогноз"
											color="primary"
											icon="insights"
											type="submit"
											:loading="isPredicting"
											class="full-width"
											size="lg"
											style="border-radius: 8px"
										/>
									</div>
								</q-form>
							</q-card-section>
						</q-card>
					</div>

					<!-- ПРАВАЯ КОЛОНКА: Результат -->
					<div class="col-12 col-md-4">
						<!-- Состояние до получения прогноза -->
						<div
							v-if="!result"
							class="flex flex-center text-grey-5"
							style="
								height: 100%;
								min-height: 300px;
								border: 2px dashed #ccc;
								border-radius: 16px;
							"
						>
							<div class="text-center">
								<q-icon name="batch_prediction" size="48px" class="q-mb-sm" />
								<div>
									Заполните данные и нажмите<br />
									«Сделать прогноз»
								</div>
							</div>
						</div>

						<!-- Карточка результата -->
						<q-card
							v-else
							style="
								border-radius: 16px;
								box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
								overflow: hidden;
							"
						>
							<!-- Заголовок результата -->
							<div
								:class="
									result.predicted_class === 1 ? 'bg-negative' : 'bg-positive'
								"
								class="q-pa-lg text-white text-center"
							>
								<q-icon
									:name="
										result.predicted_class === 1 ? 'warning' : 'check_circle'
									"
									size="48px"
									class="q-mb-sm"
								/>
								<div class="text-h5 text-weight-bold">
									{{ result.label }}
								</div>
								<div class="text-subtitle2 q-mt-sm opacity-80">
									Модель: {{ result.model }}
								</div>
							</div>

							<q-card-section class="q-pa-lg">
								<!-- Оценка модели -->
								<div class="text-center q-mb-lg">
									<div class="text-caption text-grey-7">
										Оценка модели (Score)
									</div>
									<div
										class="text-h2 text-weight-bolder"
										:class="
											result.predicted_class === 1
												? 'text-negative'
												: 'text-positive'
										"
									>
										{{ (result.score * 100).toFixed(1) }}%
									</div>
									<q-linear-progress
										:value="result.score"
										:color="
											result.predicted_class === 1 ? 'negative' : 'positive'
										"
										track-color="grey-3"
										size="12px"
										class="q-mt-md"
										style="border-radius: 6px"
									/>
									<div
										class="row justify-between text-caption text-grey-6 q-mt-xs"
									>
										<div>0%</div>
										<div>Порог: {{ (result.threshold * 100).toFixed(1) }}%</div>
										<div>100%</div>
									</div>
								</div>

								<!-- Пояснение -->
								<q-banner
									class="bg-grey-2 text-grey-8"
									style="border-radius: 8px"
								>
									<template v-slot:avatar>
										<q-icon name="info" color="grey-6" />
									</template>
									<div class="text-caption" style="line-height: 1.4">
										{{ result.note }}
									</div>
								</q-banner>
							</q-card-section>
						</q-card>
					</div>
				</div>

				<!-- Нижний отступ -->
				<div style="height: 30px"></div>
			</div>
		</q-scroll-area>
	</q-page>
</template>

<script setup>
import { ref } from 'vue'
import { Predict } from '../../bindings/DefectAnalyzer/desktop/core/handlers/predicthandler'

const isPredicting = ref(false)
const result = ref(null)

// Начальные данные (из примера API)
const form = ref({
	ProductionVolume: 850,
	ProductionCost: 12000,
	SupplierQuality: 85,
	DeliveryDelay: 3,
	MaintenanceHours: 12,
	DowntimePercentage: 8,
	InventoryTurnover: 5,
	StockoutRate: 0.05,
	WorkerProductivity: 80,
	SafetyIncidents: 1,
	EnergyConsumption: 1500,
	EnergyEfficiency: 0.8,
	AdditiveProcessTime: 4,
	AdditiveMaterialCost: 200,
})

const onPredict = async () => {
	isPredicting.value = true
	result.value = null
	try {
		const response = await Predict(form.value)
		result.value = response
	} catch (error) {
		console.error('Ошибка при прогнозе:', error)
	} finally {
		isPredicting.value = false
	}
}
</script>

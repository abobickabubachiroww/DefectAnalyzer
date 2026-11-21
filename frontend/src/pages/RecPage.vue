<template>
	<q-page class="row bg-grey-2">
		<q-scroll-area class="col" style="overflow-x: hidden">
			<div class="q-px-lg q-pb-lg q-pt-md">
				<div class="text-h6 q-mb-lg text-weight-bold text-grey-9">
					Рекомендации по улучшению качества
				</div>

				<q-tabs
					v-model="tab"
					class="text-grey-7 q-mb-lg"
					active-color="primary"
					indicator-color="primary"
					align="left"
				>
					<q-tab name="general" icon="lightbulb" label="Общие рекомендации" />

					<q-tab
						name="personalized"
						icon="person_search"
						label="Для конкретной партии"
					/>
				</q-tabs>

				<q-tab-panels v-model="tab" animated class="bg-transparent">
					<!-- ВКЛАДКА 1: ОБЩИЕ РЕКОМЕНДАЦИИ -->

					<q-tab-panel name="general" class="q-pa-none">
						<div v-if="isLoadingGeneral" class="flex flex-center q-pa-xl">
							<q-spinner-dots color="primary" size="3em" />
						</div>

						<div v-else-if="generalRecs" class="row q-col-gutter-lg">
							<div
								class="col-12"
								v-for="rec in generalRecs.recommendations"
								:key="rec.id"
							>
								<q-card
									style="
										border-radius: 16px;

										box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
									"
								>
									<q-card-section class="row items-start q-pb-none">
										<q-icon
											name="build_circle"
											size="32px"
											color="primary"
											class="q-mr-md"
										/>

										<div>
											<div class="text-h6 text-weight-bold text-grey-9">
												{{ rec.title }}
											</div>

											<div class="text-body2 text-grey-7 q-mt-xs">
												{{ rec.basis }}
											</div>
										</div>
									</q-card-section>

									<q-card-section>
										<div
											class="text-caption text-grey-6 text-weight-bold q-mb-sm text-uppercase"
										>
											Рекомендуемые действия:
										</div>

										<q-list dense>
											<q-item
												v-for="(action, idx) in rec.actions"
												:key="idx"
												class="q-pl-none"
											>
												<q-item-section avatar style="min-width: 24px">
													<q-icon
														name="check_circle_outline"
														color="positive"
														size="sm"
													/>
												</q-item-section>

												<q-item-section class="text-body2 text-grey-8">{{
													action
												}}</q-item-section>
											</q-item>
										</q-list>
									</q-card-section>
								</q-card>
							</div>

							<div class="col-12">
								<q-banner
									class="bg-grey-3 text-grey-8 q-mt-md"
									style="border-radius: 16px"
								>
									<template v-slot:avatar><q-icon name="info" /></template>

									{{ generalRecs.disclaimer }}
								</q-banner>
							</div>
						</div>
					</q-tab-panel>

					<!-- ВКЛАДКА 2: ПЕРСОНАЛЬНЫЕ РЕКОМЕНДАЦИИ -->

					<q-tab-panel name="personalized" class="q-pa-none">
						<div class="row q-col-gutter-lg">
							<!-- Форма ввода -->

							<div class="col-12 col-md-5">
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
											Параметры партии
										</div>

										<q-form @submit="onPersonalize" class="row q-col-gutter-sm">
											<!-- Упрощенная форма для компактности -->

											<div class="col-6">
												<q-input
													v-model.number="form.ProductionVolume"
													type="number"
													label="Объём"
													outlined
													dense
												/>
											</div>

											<div class="col-6">
												<q-input
													v-model.number="form.MaintenanceHours"
													type="number"
													label="Обслуживание"
													outlined
													dense
												/>
											</div>

											<div class="col-6">
												<q-input
													v-model.number="form.SupplierQuality"
													type="number"
													label="Кач-во поставщика"
													outlined
													dense
												/>
											</div>

											<div class="col-6">
												<q-input
													v-model.number="form.ProductionCost"
													type="number"
													label="Себестоимость"
													outlined
													dense
												/>
											</div>

											<q-expansion-item
												label="Дополнительные параметры"
												class="col-12 q-mt-sm"
												dense
											>
												<div class="row q-col-gutter-sm q-pt-sm">
													<div class="col-6">
														<q-input
															v-model.number="form.DeliveryDelay"
															type="number"
															label="Задержка"
															outlined
															dense
														/>
													</div>

													<div class="col-6">
														<q-input
															v-model.number="form.DowntimePercentage"
															type="number"
															label="Простой %"
															outlined
															dense
														/>
													</div>

													<div class="col-6">
														<q-input
															v-model.number="form.WorkerProductivity"
															type="number"
															label="Производит."
															outlined
															dense
														/>
													</div>

													<div class="col-6">
														<q-input
															v-model.number="form.EnergyEfficiency"
															type="number"
															label="Энергоэфф."
															outlined
															dense
														/>
													</div>
												</div>
											</q-expansion-item>

											<div class="col-12 q-mt-md">
												<q-btn
													label="Получить рекомендации"
													color="primary"
													type="submit"
													:loading="isLoadingPersonal"
													class="full-width"
													style="border-radius: 16px"
												/>
											</div>
										</q-form>
									</q-card-section>
								</q-card>
							</div>

							<!-- Результаты SHAP -->

							<div class="col-12 col-md-7">
								<div
									v-if="!personalRecs && !isLoadingPersonal"
									class="flex flex-center text-grey-5"
									style="
										height: 100%;

										min-height: 300px;

										border: 2px dashed #ccc;

										border-radius: 16px;
									"
								>
									<div class="text-center">
										<q-icon name="psychology" size="48px" class="q-mb-sm" />

										<div>
											Введите данные партии, чтобы ИИ<br />проанализировал
											факторы риска
										</div>
									</div>
								</div>

								<div
									v-if="isLoadingPersonal"
									class="flex flex-center"
									style="height: 100%; min-height: 300px"
								>
									<q-spinner-dots color="primary" size="3em" />
								</div>

								<div v-if="personalRecs" class="column q-gutter-y-md">
									<!-- Баннер с итогом -->

									<q-card
										:class="
											personalRecs.predicted_class === 1
												? 'bg-negative'
												: 'bg-positive'
										"
										class="text-white"
										style="border-radius: 16px"
									>
										<q-card-section class="row items-center">
											<q-icon
												:name="
													personalRecs.predicted_class === 1
														? 'warning'
														: 'check_circle'
												"
												size="32px"
												class="q-mr-md"
											/>

											<div>
												<div class="text-h6 text-weight-bold">
													{{ personalRecs.label }}
												</div>

												<div class="text-caption opacity-80">
													Оценка модели:

													{{ (personalRecs.score * 100).toFixed(1) }}%
												</div>
											</div>
										</q-card-section>
									</q-card>

									<!-- Карточки факторов -->

									<q-card
										v-for="(rec, idx) in personalRecs.recommendations"
										:key="idx"
										style="
											border-radius: 16px;

											box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
										"
									>
										<q-card-section>
											<div class="row items-center justify-between q-mb-sm">
												<div
													class="text-subtitle1 text-weight-bold text-grey-9"
												>
													{{ rec.title }}
												</div>

												<!-- Бейдж приоритета -->

												<q-chip
													:color="
														rec.priority === 'high'
															? 'negative'
															: rec.priority === 'medium'
																? 'warning'
																: 'grey-5'
													"
													text-color="white"
													dense
												>
													{{
														rec.priority === 'high'
															? 'Высокий риск'
															: rec.priority === 'medium'
																? 'Средний риск'
																: 'Низкий риск'
													}}
												</q-chip>
											</div>

											<div class="text-body2 text-grey-7 q-mb-md">
												{{ rec.description }}
											</div>

											<!-- Влияние SHAP -->

											<div
												class="bg-grey-2 q-pa-sm flex items-center q-mb-md"
												style="border-radius: 16px"
											>
												<q-icon
													name="insights"
													color="primary"
													class="q-mr-sm"
												/>

												<div class="text-caption text-grey-8">
													Признак <b>{{ rec.feature }}</b> (значение:

													{{ rec.value }}) повышает риск на

													<b>{{ (rec.shap_value * 100).toFixed(1) }}%</b>
												</div>
											</div>

											<q-list dense>
												<q-item
													v-for="(action, i) in rec.actions"
													:key="i"
													class="q-pl-none"
												>
													<q-item-section avatar style="min-width: 24px">
														<q-icon
															name="arrow_right"
															color="primary"
															size="sm"
														/>
													</q-item-section>

													<q-item-section class="text-body2 text-grey-8">{{
														action
													}}</q-item-section>
												</q-item>
											</q-list>
										</q-card-section>
									</q-card>
								</div>
							</div>
						</div>
					</q-tab-panel>
				</q-tab-panels>

				<div style="height: 30px"></div>
			</div>
		</q-scroll-area>
	</q-page>
</template>

<script setup>
import { ref, onMounted } from 'vue'

import {
	GetGeneralRecommendations,
	GetPersonalizedRecommendations,
} from '../../bindings/DefectAnalyzer/desktop/core/handlers/recommendationshandler'

const tab = ref('general')

// Состояния

const generalRecs = ref(null)

const personalRecs = ref(null)

const isLoadingGeneral = ref(true)

const isLoadingPersonal = ref(false)

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

onMounted(async () => {
	try {
		generalRecs.value = await GetGeneralRecommendations()
	} catch (error) {
		console.error('Ошибка при загрузке общих рекомендаций:', error)
	} finally {
		isLoadingGeneral.value = false
	}
})

const onPersonalize = async () => {
	isLoadingPersonal.value = true

	personalRecs.value = null

	try {
		personalRecs.value = await GetPersonalizedRecommendations(form.value)
	} catch (error) {
		console.error('Ошибка при генерации персональных рекомендаций:', error)
	} finally {
		isLoadingPersonal.value = false
	}
}
</script>

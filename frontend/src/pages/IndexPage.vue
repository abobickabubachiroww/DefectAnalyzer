<template>
	<q-page class="row bg-grey-2">
		<q-scroll-area class="col" style="overflow-x: hidden">
			<div class="q-px-lg q-pb-lg q-pt-md">
				<div v-if="isLoading" class="flex flex-center" style="height: 50vh">
					<q-spinner-dots color="primary" size="4em" />
					<div class="text-h6 q-ml-md">Загрузка данных...</div>
				</div>

				<div v-else-if="dashboardData && modelData">
					<!-- МЕТРИКИ МОДЕЛИ -->
					<div class="text-h6 q-mb-md text-weight-bold text-grey-9">
						Качество ИИ-модели ({{ modelData.model }})
					</div>

					<div class="row q-col-gutter-lg q-mb-lg">
						<!-- Accuracy -->
						<div class="col-6 col-md-3">
							<q-card
								class="bg-white text-center q-pa-md"
								style="
									border-radius: 16px;
									box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
								"
							>
								<div class="text-subtitle2 text-grey-7 flex flex-center">
									Accuracy
									<q-icon
										name="info_outline"
										size="xs"
										class="q-ml-xs cursor-pointer text-grey-5"
									>
										<q-menu>
											<div
												class="q-pa-md text-body2 bg-white text-black shadow-2"
												style="max-width: 250px; border-radius: 8px"
											>
												Общая доля правильных ответов модели (и по норме, и по
												браку).
											</div>
										</q-menu>
									</q-icon>
								</div>
								<div class="text-h4 text-primary text-weight-bold">
									{{ (modelData.metrics.accuracy * 100).toFixed(1) }}%
								</div>
							</q-card>
						</div>

						<!-- F1-Score -->
						<div class="col-6 col-md-3">
							<q-card
								class="bg-white text-center q-pa-md"
								style="
									border-radius: 16px;
									box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
								"
							>
								<div class="text-subtitle2 text-grey-7 flex flex-center">
									F1-Score
									<q-icon
										name="info_outline"
										size="xs"
										class="q-ml-xs cursor-pointer text-grey-5"
									>
										<q-menu>
											<div
												class="q-pa-md text-body2 bg-white text-black shadow-2"
												style="max-width: 250px; border-radius: 8px"
											>
												Баланс между точностью и полнотой. Показывает, насколько
												хорошо модель находит брак, не поднимая ложную панику.
											</div>
										</q-menu>
									</q-icon>
								</div>
								<div class="text-h4 text-primary text-weight-bold">
									{{ (modelData.metrics.f1_class_1 * 100).toFixed(1) }}%
								</div>
							</q-card>
						</div>

						<!-- ROC-AUC -->
						<div class="col-6 col-md-3">
							<q-card
								class="bg-white text-center q-pa-md"
								style="
									border-radius: 16px;
									box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
								"
							>
								<div class="text-subtitle2 text-grey-7 flex flex-center">
									ROC-AUC
									<q-icon
										name="info_outline"
										size="xs"
										class="q-ml-xs cursor-pointer text-grey-5"
									>
										<q-menu>
											<div
												class="q-pa-md text-body2 bg-white text-black shadow-2"
												style="max-width: 250px; border-radius: 8px"
											>
												Способность модели отличать норму от брака. 100% —
												идеал, 50% — случайное угадывание.
											</div>
										</q-menu>
									</q-icon>
								</div>
								<div class="text-h4 text-primary text-weight-bold">
									{{ (modelData.metrics.roc_auc * 100).toFixed(1) }}%
								</div>
							</q-card>
						</div>

						<!-- Balanced Acc -->
						<div class="col-6 col-md-3">
							<q-card
								class="bg-white text-center q-pa-md"
								style="
									border-radius: 16px;
									box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
								"
							>
								<div class="text-subtitle2 text-grey-7 flex flex-center">
									Balanced Acc
									<q-icon
										name="info_outline"
										size="xs"
										class="q-ml-xs cursor-pointer text-grey-5"
									>
										<q-menu>
											<div
												class="q-pa-md text-body2 bg-white text-black shadow-2"
												style="max-width: 250px; border-radius: 8px"
											>
												Средняя точность по обоим классам (норма и брак) по
												отдельности. Полезна, когда классов разное количество.
											</div>
										</q-menu>
									</q-icon>
								</div>
								<div class="text-h4 text-primary text-weight-bold">
									{{ (modelData.metrics.balanced_accuracy * 100).toFixed(1) }}%
								</div>
							</q-card>
						</div>
					</div>

					<!-- МАТРИЦА ОШИБОК И СТАТИСТИКА -->
					<div class="row q-col-gutter-lg">
						<!-- Матрица ошибок -->
						<div class="col-12 col-md-4">
							<q-card
								class="bg-white text-dark"
								style="
									height: 100%;
									border-radius: 16px;
									box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
								"
							>
								<q-card-section>
									<div class="text-subtitle1 text-weight-bold text-grey-8">
										Матрица ошибок (Test: {{ modelData.test_size }})
									</div>
								</q-card-section>
								<q-card-section class="q-pt-none">
									<div class="confusion-matrix">
										<!-- TN -->
										<div class="cm-cell bg-positive text-white">
											<div class="text-h4 text-weight-bold">
												{{ modelData.confusion_matrix[0][0] }}
											</div>
											<div
												class="text-subtitle2 text-weight-medium q-mt-xs flex items-center justify-center"
											>
												TN
												<q-icon
													name="info_outline"
													size="xs"
													class="q-ml-xs opacity-80"
												/>
											</div>
											<q-menu>
												<div
													class="q-pa-md text-body2 bg-white text-black shadow-2"
													style="max-width: 200px; border-radius: 8px"
												>
													<b>Верная норма (TN).</b> Модель сказала «всё ок», и
													деталь действительно без дефектов.
												</div>
											</q-menu>
										</div>
										<!-- FP -->
										<div class="cm-cell bg-negative text-white">
											<div class="text-h4 text-weight-bold">
												{{ modelData.confusion_matrix[0][1] }}
											</div>
											<div
												class="text-subtitle2 text-weight-medium q-mt-xs flex items-center justify-center"
											>
												FP
												<q-icon
													name="info_outline"
													size="xs"
													class="q-ml-xs opacity-80"
												/>
											</div>
											<q-menu>
												<div
													class="q-pa-md text-body2 bg-white text-black shadow-2"
													style="max-width: 200px; border-radius: 8px"
												>
													<b>Ложная тревога (FP).</b> Модель сказала «брак», но
													деталь на самом деле нормальная.
												</div>
											</q-menu>
										</div>
										<!-- FN -->
										<div class="cm-cell bg-negative text-white">
											<div class="text-h4 text-weight-bold">
												{{ modelData.confusion_matrix[1][0] }}
											</div>
											<div
												class="text-subtitle2 text-weight-medium q-mt-xs flex items-center justify-center"
											>
												FN
												<q-icon
													name="info_outline"
													size="xs"
													class="q-ml-xs opacity-80"
												/>
											</div>
											<q-menu>
												<div
													class="q-pa-md text-body2 bg-white text-black shadow-2"
													style="max-width: 200px; border-radius: 8px"
												>
													<b>Пропуск брака (FN).</b> Модель пропустила дефект
													(самая опасная ошибка).
												</div>
											</q-menu>
										</div>
										<!-- TP -->
										<div class="cm-cell bg-positive text-white">
											<div class="text-h4 text-weight-bold">
												{{ modelData.confusion_matrix[1][1] }}
											</div>
											<div
												class="text-subtitle2 text-weight-medium q-mt-xs flex items-center justify-center"
											>
												TP
												<q-icon
													name="info_outline"
													size="xs"
													class="q-ml-xs opacity-80"
												/>
											</div>
											<q-menu>
												<div
													class="q-pa-md text-body2 bg-white text-black shadow-2"
													style="max-width: 200px; border-radius: 8px"
												>
													<b>Верный брак (TP).</b> Модель нашла дефект, и он
													действительно есть.
												</div>
											</q-menu>
										</div>
									</div>
								</q-card-section>
							</q-card>
						</div>

						<!-- Карточка Всего + Пончик -->
						<div class="col-12 col-md-8">
							<div class="row q-col-gutter-lg" style="height: 100%">
								<div class="col-12 col-sm-5">
									<q-card
										class="bg-primary text-white flex flex-center"
										style="
											height: 100%;
											border-radius: 16px;
											box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
										"
									>
										<div class="text-center q-pa-md">
											<q-icon
												name="dataset"
												size="48px"
												class="q-mb-sm opacity-80"
											/>
											<div class="text-subtitle1 opacity-80">
												Всего наблюдений
											</div>
											<div class="text-h2 text-weight-bolder q-mb-sm">
												{{ dashboardData.total }}
											</div>
											<div
												class="row justify-center text-caption bg-white text-primary q-px-sm q-py-xs"
												style="border-radius: 8px; opacity: 0.9"
											>
												<div class="q-mr-md">
													<q-icon
														name="view_column"
														size="14px"
														class="q-mr-xs"
													/>Колонок: {{ dashboardData.columns }}
												</div>
												<div>
													<q-icon
														name="error_outline"
														size="14px"
														class="q-mr-xs"
													/>Пропусков: {{ dashboardData.missing_values }}
												</div>
											</div>
										</div>
									</q-card>
								</div>
								<div class="col-12 col-sm-7">
									<q-card
										style="
											height: 100%;
											border-radius: 16px;
											box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
										"
									>
										<q-card-section>
											<div class="text-subtitle1 text-weight-bold text-grey-8">
												Распределение по классам
											</div>
										</q-card-section>
										<q-card-section class="q-pt-none">
											<v-chart class="chart" :option="pieOption" autoresize />
										</q-card-section>
									</q-card>
								</div>
							</div>
						</div>

						<!-- Столбцы (Обслуживание) -->
						<div class="col-12 col-md-6">
							<q-card
								style="
									border-radius: 16px;
									box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
								"
							>
								<q-card-section>
									<div class="text-subtitle1 text-weight-bold text-grey-8">
										Среднее время обслуживания (ч)
									</div>
								</q-card-section>
								<q-card-section class="q-pt-none">
									<v-chart
										class="chart"
										:option="maintenanceOption"
										autoresize
									/>
								</q-card-section>
							</q-card>
						</div>

						<!-- Столбцы (Объем) -->
						<div class="col-12 col-md-6">
							<q-card
								style="
									border-radius: 16px;
									box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
								"
							>
								<q-card-section>
									<div class="text-subtitle1 text-weight-bold text-grey-8">
										Средний объём производства
									</div>
								</q-card-section>
								<q-card-section class="q-pt-none">
									<v-chart class="chart" :option="volumeOption" autoresize />
								</q-card-section>
							</q-card>
						</div>
					</div>
					<div style="height: 30px"></div>
				</div>
			</div>
		</q-scroll-area>
	</q-page>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { PieChart, BarChart } from 'echarts/charts'
import {
	TooltipComponent,
	LegendComponent,
	GridComponent,
} from 'echarts/components'
import VChart from 'vue-echarts'
import { getCssVar } from 'quasar'

import {
	GetDashboardStats,
	GetModelMetrics,
} from '../../bindings/DefectAnalyzer/desktop/core/handlers/statshandler'

use([
	CanvasRenderer,
	PieChart,
	BarChart,
	TooltipComponent,
	LegendComponent,
	GridComponent,
])

const isLoading = ref(true)
const dashboardData = ref(null)
const modelData = ref(null)

onMounted(async () => {
	try {
		const [dash, model] = await Promise.all([
			GetDashboardStats(),
			GetModelMetrics(),
		])
		dashboardData.value = dash
		modelData.value = model
	} catch (error) {
		console.error('Ошибка при получении данных:', error)
	} finally {
		isLoading.value = false
	}
})

const colorLow = getCssVar('positive')
const colorHigh = getCssVar('negative')

const pieOption = computed(() => {
	if (!dashboardData.value) return {}
	return {
		tooltip: { trigger: 'item', formatter: '{b}: {c} ({d}%)' },
		legend: { bottom: '0%', left: 'center' },
		series: [
			{
				type: 'pie',
				radius: ['50%', '70%'],
				itemStyle: { borderRadius: 8, borderColor: '#ffffff', borderWidth: 2 },
				data: [
					{
						value: dashboardData.value.class_distribution?.low || 0,
						name: 'Low (Норма)',
						itemStyle: { color: colorLow },
					},
					{
						value: dashboardData.value.class_distribution?.high || 0,
						name: 'High (Брак)',
						itemStyle: { color: colorHigh },
					},
				],
			},
		],
	}
})

const maintenanceOption = computed(() => {
	if (!dashboardData.value) return {}
	return {
		tooltip: { trigger: 'axis', axisPointer: { type: 'shadow' } },
		grid: { left: '3%', right: '4%', bottom: '3%', containLabel: true },
		xAxis: { type: 'category', data: ['Low (Норма)', 'High (Брак)'] },
		yAxis: { type: 'value' },
		series: [
			{
				type: 'bar',
				barWidth: '50%',
				data: [
					{
						value: dashboardData.value.maintenance_by_class?.low || 0,
						itemStyle: { color: colorLow },
					},
					{
						value: dashboardData.value.maintenance_by_class?.high || 0,
						itemStyle: { color: colorHigh },
					},
				],
				itemStyle: { borderRadius: [4, 4, 0, 0] },
				label: { show: true, position: 'top' },
			},
		],
	}
})

const volumeOption = computed(() => {
	if (!dashboardData.value) return {}
	return {
		tooltip: { trigger: 'axis', axisPointer: { type: 'shadow' } },
		grid: { left: '3%', right: '4%', bottom: '3%', containLabel: true },
		xAxis: { type: 'category', data: ['Low (Норма)', 'High (Брак)'] },
		yAxis: { type: 'value' },
		series: [
			{
				type: 'bar',
				barWidth: '50%',
				data: [
					{
						value: dashboardData.value.production_volume_by_class?.low || 0,
						itemStyle: { color: colorLow },
					},
					{
						value: dashboardData.value.production_volume_by_class?.high || 0,
						itemStyle: { color: colorHigh },
					},
				],
				itemStyle: { borderRadius: [4, 4, 0, 0] },
				label: { show: true, position: 'top' },
			},
		],
	}
})
</script>

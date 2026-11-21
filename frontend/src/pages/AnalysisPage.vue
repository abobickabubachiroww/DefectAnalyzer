<template>
	<q-page class="row bg-grey-2">
		<!-- 👇 Весь контент обернут в скролл-зону, которая занимает все свободное место (класс col) -->
		<q-scroll-area class="col q-pa-md">
			<div v-if="isLoading" class="flex flex-center" style="height: 50vh">
				<q-spinner-dots color="primary" size="4em" />
				<div class="text-h6 q-ml-md">Загрузка данных...</div>
			</div>

			<div v-else-if="featuresList">
				<!-- БЛОК 1: Распределение показателя -->
				<q-card style="border-radius: 12px" class="q-mb-md">
					<q-card-section class="row items-center q-col-gutter-md">
						<div class="col-12 col-md-6">
							<div class="text-h6 text-weight-bold text-grey-9">
								Распределение показателя
							</div>
						</div>
						<div class="col-12 col-md-6">
							<q-select
								v-model="selectedFeature"
								:options="featuresList"
								option-label="label"
								option-value="name"
								emit-value
								map-options
								outlined
								dense
								label="Выберите показатель"
								@update:model-value="fetchDistribution"
								popup-content-class="bg-white text-dark"
							>
								<template v-slot:option="scope">
									<q-item v-bind="scope.itemProps" class="text-dark">
										<q-item-section>
											<q-item-label class="text-weight-medium">{{
												scope.opt.label
											}}</q-item-label>
											<q-item-label caption class="text-grey-7">{{
												scope.opt.unit
											}}</q-item-label>
										</q-item-section>
									</q-item>
								</template>
							</q-select>
						</div>
					</q-card-section>
					<q-card-section class="q-pt-none">
						<div
							v-if="isDistributionLoading"
							class="flex flex-center"
							style="height: 250px"
						>
							<q-spinner color="primary" size="2em" />
						</div>
						<v-chart
							v-else-if="distributionData"
							style="height: 250px; width: 100%"
							:option="distributionOption"
							autoresize
						/>
					</q-card-section>
				</q-card>

				<!-- БЛОК 2: Время обслуживания -->
				<div v-if="maintenanceGroups" class="q-mb-md">
					<div class="text-h6 text-weight-bold text-grey-9 q-mb-sm">
						Время обслуживания
					</div>
					<div class="row q-col-gutter-md">
						<div class="col-12 col-md-3">
							<q-card
								class="bg-white text-dark full-width flex flex-center"
								style="
									border-radius: 12px;
									border-left: 4px solid var(--q-positive);
								"
							>
								<div class="q-pa-sm text-center">
									<div class="text-caption text-grey-7">
										Меньше {{ maintenanceGroups.threshold_hours_per_week }} ч
									</div>
									<div class="text-h5 text-weight-bold q-mt-xs">
										{{ maintenanceGroups.below_threshold.total }}
									</div>
									<div
										class="text-negative text-weight-bold q-mt-xs text-caption"
									>
										Брак:
										{{
											(
												maintenanceGroups.below_threshold.high_defect_share *
												100
											).toFixed(1)
										}}%
									</div>
								</div>
							</q-card>
						</div>
						<div class="col-12 col-md-3">
							<q-card
								class="bg-white text-dark full-width flex flex-center"
								style="
									border-radius: 12px;
									border-left: 4px solid var(--q-accent);
								"
							>
								<div class="q-pa-sm text-center">
									<div class="text-caption text-grey-7">
										От {{ maintenanceGroups.threshold_hours_per_week }} ч и выше
									</div>
									<div class="text-h5 text-weight-bold q-mt-xs">
										{{ maintenanceGroups.at_or_above_threshold.total }}
									</div>
									<div
										class="text-negative text-weight-bold q-mt-xs text-caption"
									>
										Брак:
										{{
											(
												maintenanceGroups.at_or_above_threshold
													.high_defect_share * 100
											).toFixed(1)
										}}%
									</div>
								</div>
							</q-card>
						</div>
						<div class="col-12 col-md-6" style="min-width: 0">
							<q-card
								style="border-radius: 12px; height: 100%; overflow: hidden"
							>
								<v-chart
									style="height: 200px; width: 100%"
									:option="maintenanceGroupsOption"
									autoresize
								/>
							</q-card>
						</div>
					</div>
				</div>

				<!-- БЛОК 3: Объём производства -->
				<div v-if="productionGroups" class="q-mb-md">
					<div class="text-h6 text-weight-bold text-grey-9 q-mb-sm">
						Объём производства
					</div>
					<div class="row q-col-gutter-md">
						<div class="col-12 col-md-3">
							<q-card
								class="bg-white text-dark full-width flex flex-center"
								style="
									border-radius: 12px;
									border-left: 4px solid var(--q-positive);
								"
							>
								<div class="q-pa-sm text-center">
									<div class="text-caption text-grey-7">
										Меньше {{ productionGroups.threshold_units_per_day }} ед
									</div>
									<div class="text-h5 text-weight-bold q-mt-xs">
										{{ productionGroups.below_threshold.total }}
									</div>
									<div
										class="text-negative text-weight-bold q-mt-xs text-caption"
									>
										Брак:
										{{
											(
												productionGroups.below_threshold.high_defect_share * 100
											).toFixed(1)
										}}%
									</div>
								</div>
							</q-card>
						</div>
						<div class="col-12 col-md-3">
							<q-card
								class="bg-white text-dark full-width flex flex-center"
								style="
									border-radius: 12px;
									border-left: 4px solid var(--q-accent);
								"
							>
								<div class="q-pa-sm text-center">
									<div class="text-caption text-grey-7">
										От {{ productionGroups.threshold_units_per_day }} ед и выше
									</div>
									<div class="text-h5 text-weight-bold q-mt-xs">
										{{ productionGroups.at_or_above_threshold.total }}
									</div>
									<div
										class="text-negative text-weight-bold q-mt-xs text-caption"
									>
										Брак:
										{{
											(
												productionGroups.at_or_above_threshold
													.high_defect_share * 100
											).toFixed(1)
										}}%
									</div>
								</div>
							</q-card>
						</div>
						<div class="col-12 col-md-6" style="min-width: 0">
							<q-card
								style="border-radius: 12px; height: 100%; overflow: hidden"
							>
								<v-chart
									style="height: 200px; width: 100%"
									:option="productionGroupsOption"
									autoresize
								/>
							</q-card>
						</div>
					</div>
				</div>

				<!-- БЛОК 4: Совместный риск -->
				<div v-if="riskGroups" class="q-mb-md">
					<div
						class="text-h6 text-weight-bold text-grey-9 q-mb-sm flex items-center"
					>
						Совместный риск
						<q-icon
							name="help_outline"
							size="sm"
							class="q-ml-sm text-grey-6 cursor-pointer"
						>
							<q-menu>
								<div
									class="q-pa-md text-body2 bg-white text-black shadow-2"
									style="max-width: 250px"
								>
									{{ riskGroups.note }}
								</div>
							</q-menu>
						</q-icon>
					</div>
					<div class="row q-col-gutter-md">
						<div class="col-12" style="min-width: 0">
							<q-card style="border-radius: 12px; overflow: hidden">
								<q-card-section>
									<div class="text-caption text-grey-7 q-mb-sm">
										Пороги: Обслуживание &ge;
										{{ riskGroups.thresholds.maintenance_hours_per_week }} ч и
										Объем &ge;
										{{ riskGroups.thresholds.production_units_per_day }} ед
									</div>
									<v-chart
										style="height: 300px; width: 100%"
										:option="riskGroupsOption"
										autoresize
									/>
								</q-card-section>
							</q-card>
						</div>
					</div>
				</div>

				<!-- Отступ снизу, чтобы последний блок не прилипал к краю экрана -->
				<div style="height: 30px"></div>
			</div>
		</q-scroll-area>
	</q-page>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { BarChart } from 'echarts/charts'
import {
	TooltipComponent,
	LegendComponent,
	GridComponent,
} from 'echarts/components'
import VChart from 'vue-echarts'
import { getCssVar } from 'quasar'

import {
	GetFeatures,
	GetFeatureDistribution,
	GetMaintenanceGroups,
	GetProductionGroups,
	GetRiskGroups,
} from '../../bindings/DefectAnalyzer/desktop/core/handlers/analysishandler'

use([
	CanvasRenderer,
	BarChart,
	TooltipComponent,
	LegendComponent,
	GridComponent,
])

const isLoading = ref(true)
const isDistributionLoading = ref(false)

const featuresList = ref([])
const selectedFeature = ref(null)
const distributionData = ref(null)

const maintenanceGroups = ref(null)
const productionGroups = ref(null)
const riskGroups = ref(null)

const colorLow = getCssVar('positive') || '#16a34a'
const colorHigh = getCssVar('negative') || '#dc2626'

onMounted(async () => {
	try {
		const [featuresRes, maintRes, prodRes, riskRes] = await Promise.all([
			GetFeatures(),
			GetMaintenanceGroups(),
			GetProductionGroups(),
			GetRiskGroups(),
		])

		featuresList.value = featuresRes.features
		maintenanceGroups.value = maintRes
		productionGroups.value = prodRes
		riskGroups.value = riskRes

		if (featuresList.value.length > 0) {
			selectedFeature.value = featuresList.value[0].name
			await fetchDistribution()
		}
	} catch (error) {
		console.error('Ошибка при загрузке данных:', error)
	} finally {
		isLoading.value = false
	}
})

const fetchDistribution = async () => {
	if (!selectedFeature.value) return
	isDistributionLoading.value = true
	try {
		distributionData.value = await GetFeatureDistribution(
			selectedFeature.value,
			15,
		)
	} catch (error) {
		console.error('Ошибка при загрузке распределения:', error)
	} finally {
		isDistributionLoading.value = false
	}
}

// --- НАСТРОЙКИ ГРАФИКОВ ---

const distributionOption = computed(() => {
	if (!distributionData.value) return {}
	const data = distributionData.value
	const labels = []
	for (let i = 0; i < data.bin_edges.length - 1; i++) {
		labels.push(`${data.bin_edges[i]} - ${data.bin_edges[i + 1]}`)
	}
	const lowCounts = data.series.find(s => s.class === 0)?.counts || []
	const highCounts = data.series.find(s => s.class === 1)?.counts || []

	return {
		tooltip: { trigger: 'axis', axisPointer: { type: 'shadow' } },
		legend: { bottom: '0%', left: 'center' },
		grid: { left: '3%', right: '4%', bottom: '10%', containLabel: true },
		xAxis: {
			type: 'category',
			data: labels,
			name: data.unit,
			nameLocation: 'middle',
			nameGap: 30,
		},
		yAxis: { type: 'value', name: 'Кол-во деталей' },
		series: [
			{
				name: 'Low (Норма)',
				type: 'bar',
				itemStyle: { color: colorLow },
				barGap: '10%',
				data: lowCounts,
			},
			{
				name: 'High (Брак)',
				type: 'bar',
				itemStyle: { color: colorHigh },
				data: highCounts,
			},
		],
	}
})

const maintenanceGroupsOption = computed(() => {
	if (!maintenanceGroups.value) return {}
	const data = maintenanceGroups.value
	return {
		tooltip: { trigger: 'axis', axisPointer: { type: 'shadow' } },
		grid: { left: '3%', right: '4%', bottom: '10%', containLabel: true },
		xAxis: {
			type: 'category',
			data: [
				`< ${data.threshold_hours_per_week}ч`,
				`>= ${data.threshold_hours_per_week}ч`,
			],
			axisLabel: { fontSize: 11 },
		},
		yAxis: { type: 'value', show: false },
		series: [
			{
				name: 'Low',
				type: 'bar',
				stack: 'total',
				itemStyle: { color: colorLow },
				data: [
					data.below_threshold.low_defect,
					data.at_or_above_threshold.low_defect,
				],
			},
			{
				name: 'High',
				type: 'bar',
				stack: 'total',
				itemStyle: { color: colorHigh },
				data: [
					data.below_threshold.high_defect,
					data.at_or_above_threshold.high_defect,
				],
			},
		],
	}
})

const productionGroupsOption = computed(() => {
	if (!productionGroups.value) return {}
	const data = productionGroups.value
	return {
		tooltip: { trigger: 'axis', axisPointer: { type: 'shadow' } },
		grid: { left: '3%', right: '4%', bottom: '10%', containLabel: true },
		xAxis: {
			type: 'category',
			data: [
				`< ${data.threshold_units_per_day}ед`,
				`>= ${data.threshold_units_per_day}ед`,
			],
			axisLabel: { fontSize: 11 },
		},
		yAxis: { type: 'value', show: false },
		series: [
			{
				name: 'Low',
				type: 'bar',
				stack: 'total',
				itemStyle: { color: colorLow },
				data: [
					data.below_threshold.low_defect,
					data.at_or_above_threshold.low_defect,
				],
			},
			{
				name: 'High',
				type: 'bar',
				stack: 'total',
				itemStyle: { color: colorHigh },
				data: [
					data.below_threshold.high_defect,
					data.at_or_above_threshold.high_defect,
				],
			},
		],
	}
})

const riskGroupsOption = computed(() => {
	if (!riskGroups.value) return {}
	const groups = riskGroups.value.groups

	// Формируем подписи для 4 столбцов
	const labels = groups.map(g => {
		const m = g.maintenance_ge_11 ? 'Обсл. ≥ 11' : 'Обсл. < 11'
		const v = g.volume_ge_800 ? 'Объем ≥ 800' : 'Объем < 800'
		return `${m}\n${v}`
	})

	const lowData = groups.map(g => g.low_defect)
	const highData = groups.map(g => g.high_defect)

	return {
		tooltip: {
			trigger: 'axis',
			axisPointer: { type: 'shadow' },
			formatter: function (params) {
				const idx = params[0].dataIndex
				const g = groups[idx]
				return `<b>${labels[idx].replace('\n', ', ')}</b><br/>
						Всего: ${g.total}<br/>
						Брак: <b>${g.high_defect}</b> (${(g.high_defect_share * 100).toFixed(1)}%)<br/>
						Норма: ${g.low_defect}`
			},
		},
		legend: { bottom: '0%', left: 'center' },
		grid: { left: '3%', right: '4%', bottom: '15%', containLabel: true },
		xAxis: {
			type: 'category',
			data: labels,
			axisLabel: { interval: 0, fontSize: 11, lineHeight: 16, color: '#555' },
		},
		yAxis: { type: 'value', name: 'Кол-во деталей' },
		series: [
			{
				name: 'Low (Норма)',
				type: 'bar',
				stack: 'total',
				itemStyle: { color: colorLow },
				data: lowData,
			},
			{
				name: 'High (Брак)',
				type: 'bar',
				stack: 'total',
				itemStyle: { color: colorHigh },
				label: {
					show: true,
					position: 'top',
					formatter: function (params) {
						return `Брак: ${(groups[params.dataIndex].high_defect_share * 100).toFixed(1)}%`
					},
					fontWeight: 'bold',
					color: colorHigh,
				},
				data: highData,
			},
		],
	}
})
</script>

<template>
	<q-layout view="hHh Lpr lFf">
		>
		<q-header
			elevated
			class="bg-primary text-white"
			style="--wails-draggable: drag"
		>
			<q-toolbar>
				<q-btn
					flat
					dense
					round
					icon="menu"
					@click="toggleLeftDrawer"
					style="--wails-draggable: no-drag"
				/>
				<q-toolbar-title> DefectAnalyzer </q-toolbar-title>
				<div>DefectAnalyzer v0.8</div>
			</q-toolbar>
		</q-header>

		<q-drawer v-model="leftDrawerOpen" show-if-above bordered class="bg-grey-1">
			<q-list padding class="text-dark">
				<q-item-label header class="text-grey-8">Меню</q-item-label>

				<q-item clickable to="/" exact class="nav-item">
					<q-item-section avatar><q-icon name="engineering" /></q-item-section>
					<q-item-section>Главная</q-item-section>
				</q-item>

				<q-item clickable to="/analysis" class="nav-item">
					<q-item-section avatar><q-icon name="query_stats" /></q-item-section>
					<q-item-section>Анализ факторов</q-item-section>
				</q-item>

				<q-item clickable to="/predict" class="nav-item">
					<q-item-section avatar
						><q-icon name="batch_prediction"
					/></q-item-section>
					<q-item-section>Контроль брака</q-item-section>
				</q-item>

				<q-item clickable to="/recomendations" class="nav-item">
					<q-item-section avatar><q-icon name="trending_up" /></q-item-section>
					<q-item-section>Рекомендации</q-item-section>
				</q-item>

				<q-item clickable to="/economic" class="nav-item">
					<q-item-section avatar>
						<q-icon name="paid" />
					</q-item-section>
					<q-item-section>Экономика</q-item-section>
				</q-item>
			</q-list>
		</q-drawer>

		<q-page-container class="bg-grey-2">
			<router-view />
		</q-page-container>
	</q-layout>
</template>

<script setup>
import { ref } from 'vue'

const leftDrawerOpen = ref(false)
function toggleLeftDrawer() {
	leftDrawerOpen.value = !leftDrawerOpen.value
}
</script>

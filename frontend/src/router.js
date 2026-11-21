import { createRouter, createWebHashHistory } from 'vue-router'
import IndexPage from './pages/IndexPage.vue'
import AnalysisPage from './pages/AnalysisPage.vue'

import PredictPage from './pages/PredictPage.vue'
import RecPage from './pages/RecPage.vue'
import EconomicPage from './pages/EconomicPage.vue'

const routes = [
	{ path: '/', component: IndexPage },
	{ path: '/analysis', component: AnalysisPage },
	{ path: '/predict', component: PredictPage },
	{ path: '/recomendations', component: RecPage },
	{ path: '/economic', component: EconomicPage },
]

const router = createRouter({
	history: createWebHashHistory(),
	routes,
})

export default router

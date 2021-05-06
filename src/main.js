import Vue from 'vue'
import App from './App.vue'
import FocusVisible from 'vue-focus-visible'

Vue.config.productionTip = false
Vue.use(FocusVisible)

new Vue({
	render: h => h(App),
}).$mount('#app')

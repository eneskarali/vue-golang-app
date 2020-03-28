import Vue from 'vue'
import App from './App.vue'
import router from './router'
import scroolView from 'vue-scrollview'

Vue.use(scroolView)


Vue.config.productionTip = false

new Vue({
  render: h => h(App), router: router
}).$mount('#app')

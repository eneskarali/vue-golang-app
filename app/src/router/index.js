import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import loginPage from '@/components/loginPage'
import dashboard from '@/components/dashboard'


Vue.use(Router)

let router = new Router({
    mode: 'history',
    routes: [
        {
            path: '/',
            name: 'HelloWorld',
            component: HelloWorld,
        },
        {
            path: '/login',
            name:'loginPage',
            component: loginPage,
            meta: {
                guest:true
            }
        },
        {
            path: '/dashboard',
            name: 'dashboard',
            component: dashboard,
            meta:{
                requiresAuth: true
            }
        }
    ]
})


router.beforeEach((to, from, next) => {
   console.log("sa")
   next()    

})

export default router

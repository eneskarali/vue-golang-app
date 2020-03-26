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
            path: '/hello',
            name: 'HelloWorld',
            component: HelloWorld,
            meta: {
                guest: true
            }
        },
        {
            path: '/login',
            name: 'loginPage',
            component: loginPage,
        },
        {
            path: '/',
            name: 'dashboard',
            component: dashboard,
            meta: {
                requiresAuth: true
            }
        }
    ]
})

function getCookie(cname) {
    var name = cname + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(';');
    for (var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return null
}

router.beforeEach((to, from, next) => {
    if (to.matched.some(record => record.meta.requiresAuth)) {
        if (getCookie('token') == null) {
            next({
                path: '/login',
                params: { nextUrl: to.fullPath }
            })
        } else {
            let token = getCookie('token')
            console.log(token)
            next()
        }
    } else {
        next()
    }

})

export default router

import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/helloWorld'
import loginPage from '@/components/loginPage'
import dashboard from '@/components/dashboard'
import axios from 'axios'

Vue.use(Router)

let router = new Router({
    mode: 'history',
    routes: [
        {
            path: '/hello',
            name: 'helloWorld',
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
            axios
                .get(
                    "http://localhost:8000/youValid",
                    { withCredentials: true }
                )
                .then(response => {
                    if (response.status == 200) {
                        localStorage.setItem('name',response.data.name)
                        localStorage.setItem('username',response.data.username)
                        next()
                    }
                })
                .catch(err => {
                    console.log("validation service unable: " + err);
                    next({
                        path: '/login',
                        params: { nextUrl: to.fullPath }
                    })
                });
        }
    } else {
        next()
    }

})

export default router

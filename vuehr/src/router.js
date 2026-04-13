import Vue from 'vue'
import Router from 'vue-router'
import Login from './views/Login.vue'
import Register from './views/Register.vue'
import Home from './views/Home.vue'
import FriendChat from './views/chat/FriendChat.vue'
import PerOnboarding from './views/per/PerOnboarding.vue'
import PerOffboarding from './views/per/PerOffboarding.vue'

Vue.use(Router)

export default new Router({
    routes: [
        {
            path: '/',
            name: 'Login',
            component: Login,
            hidden:true
        }, {
            path: '/register',
            name: 'Register',
            component: Register,
            hidden:true
        }, {
            path: '/home',
            name: 'Home',
            component: Home,
            hidden:true,
            meta:{
                roles:['admin','user']
            },
            children:[
                {
                    path: '/chat',
                    name: '在线聊天',
                    component: FriendChat,
                    hidden:true
                },
                {
                    path: '/per/onboarding',
                    name: '入职管理',
                    component: PerOnboarding,
                    hidden:true
                },
                {
                    path: '/per/offboarding',
                    name: '离职管理',
                    component: PerOffboarding,
                    hidden:true
                }
            ]
        }
    ]
})
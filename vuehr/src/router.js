import Vue from 'vue'
import Router from 'vue-router'
import Login from './views/Login.vue'
import Register from './views/Register.vue'
import Home from './views/Home.vue'
import FriendChat from './views/chat/FriendChat.vue'
import OnboardingProcess from './views/process/OnboardingProcess.vue'
import ResignationProcess from './views/process/ResignationProcess.vue'

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
                    path: '/onboarding',
                    name: '入职流程',
                    component: OnboardingProcess,
                    hidden:true
                },
                {
                    path: '/resignation',
                    name: '离职流程',
                    component: ResignationProcess,
                    hidden:true
                }
            ]
        }
    ]
})
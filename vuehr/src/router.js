import Vue from 'vue'
import Router from 'vue-router'
import Login from './views/Login.vue'
import Register from './views/Register.vue'
import Home from './views/Home.vue'
import FriendChat from './views/chat/FriendChat.vue'
import FlowDashboard from './views/flow/FlowDashboard.vue'
import OnboardingFlow from './views/flow/OnboardingFlow.vue'
import OffboardingFlow from './views/flow/OffboardingFlow.vue'

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
                    path: '/flow/dashboard',
                    name: '流程概览',
                    component: FlowDashboard,
                    iconCls: 'fa fa-random',
                    meta:{
                        roles:['admin','hr']
                    }
                },
                {
                    path: '/onboarding',
                    name: '入职流程管理',
                    component: OnboardingFlow,
                    iconCls: 'fa fa-user-plus',
                    meta:{
                        roles:['admin','hr']
                    }
                },
                {
                    path: '/offboarding',
                    name: '离职流程管理',
                    component: OffboardingFlow,
                    iconCls: 'fa fa-user-times',
                    meta:{
                        roles:['admin','hr']
                    }
                }
            ]
        }
    ]
})
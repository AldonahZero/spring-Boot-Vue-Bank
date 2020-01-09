import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

import {postRequest} from "./utils/api";
import {postKeyValueRequest} from "./utils/api";
import {putRequest} from "./utils/api";
import {deleteRequest} from "./utils/api";
import {getRequest} from "./utils/api";
import {initMenu} from "./utils/menus";
import VueParticles from 'vue-particles';
import 'font-awesome/css/font-awesome.min.css';

Vue.prototype.postRequest = postRequest;
Vue.prototype.postKeyValueRequest = postKeyValueRequest;
Vue.prototype.putRequest = putRequest;
Vue.prototype.deleteRequest = deleteRequest;
Vue.prototype.getRequest = getRequest;

Vue.config.productionTip = false

Vue.use(ElementUI,{size:'large'});
Vue.use(VueParticles);
router.beforeEach((to, from, next) => {
    if (to.path == '/'||to.path == '/register') {
        next();
    }else {
        if (window.sessionStorage.getItem("user")) {
            initMenu(router, store);
            next();
        }else{
            next('/');
        }
    }
})

new Vue({
    router,
    store,
    render: h => h(App)
}).$mount('#app')

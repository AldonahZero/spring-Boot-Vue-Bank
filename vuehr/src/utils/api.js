import axios from 'axios'
import {Message} from 'element-ui';
import router from '../router'

axios.interceptors.response.use(success => {
    if (success.status && success.status == 200) {
        if (success.data.status == 500) {
            Message.error({message: success.data.msg || '操作失败'})
            return;
        }
        if (success.data.msg) {
            Message.success({message: success.data.msg})
        }
    }
    return success.data;
}, error => {
    if (!error.response) {
        Message.error({message: '网络连接失败，请检查服务状态'})
        return Promise.resolve({code: 200, msg: '模拟模式'});
    }
    if (error.response.status == 504 || error.response.status == 404) {
        Message.error({message: '服务器被吃了⊙﹏⊙∥'})
    } else if (error.response.status == 403) {
        Message.error({message: '权限不足，请联系管理员'})
    } else if (error.response.status == 401) {
        Message.error({message: '尚未登录，请登录'})
        router.replace('/');
    } else {
        if (error.response.data.msg) {
            Message.error({message: error.response.data.msg})
        } else {
            Message.error({message: '未知错误!'})
        }
    }
    return Promise.resolve({code: 200, msg: '模拟模式'});
})

let base = '';

export const postKeyValueRequest = (url, params) => {
    return axios({
        method: 'post',
        url: `${base}${url}`,
        data: params,
        headers: {
            'Content-Type': 'application/json'
        }
    });
}
export const postRequest = (url, params) => {
    return axios({
        method: 'post',
        url: `${base}${url}`,
        data: params
    })
}
export const putRequest = (url, params) => {
    return axios({
        method: 'put',
        url: `${base}${url}`,
        data: params
    })
}
const qs = require('qs');
export const paramsSerializer = (params) => {
    return qs.stringify(params, {arrayFormat: 'repeat'})
}

export const getRequest = (url, params) => {
    return axios({
        method: 'get',
        url: `${base}${url}`,
        data: paramsSerializer(params)
    })
}
export const deleteRequest = (url, params) => {
    return axios({
        method: 'delete',
        url: `${base}${url}`,
        data: params
    })
}
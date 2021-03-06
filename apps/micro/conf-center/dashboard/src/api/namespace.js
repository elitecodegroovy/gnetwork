import request from '@/utils/request'

export function getNamespaces(params) {
    return request({
        url: '/namespaces',
        method: 'get',
        params
    })
}

export function createNamespace(data) {
    return request({
        url: '/namespace',
        method: 'post',
        data
    })
}

export function deleteNamespace(data) {
    return request({
        url: '/namespace',
        method: 'delete',
        data
    })
}


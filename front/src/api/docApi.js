import request from '@/utils/request'
import api from './API_LIST.js'

export function getContent(path, params) {
  var url = api.HOST + '/api/' + path.key + '/content/' + (path.id == undefined ? '' : path.id)
  return request({
    url: url,
    method: 'get',
    params
  })
}

export function getTree(path, params) {
  var url = api.HOST + '/api/' + path.key + '/tree'
  return request({
    url: url,
    method: 'get',
    params
  })
}

import request from '../config/request'

export const signinAjax = (p) => request('post', '/signin', p)

export const addBookAjax = (p) => request('post', '/addbook', p)
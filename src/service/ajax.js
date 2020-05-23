import request from '../config/request'

export const signinAjax = p => request('post', '/signin', p)

export const addBookAjax = p => request('post', '/addbook', p)

export const getBookbyIdAjax = p => request('get', '/getbookbyid', p)

export const getBookListAjax = p => request('get', '/getbooklist', p)

export const updateBookAjax = p => request('post', '/updatebook', p)

export const deleteBookAjax = p => request('post', '/deletebook', p)



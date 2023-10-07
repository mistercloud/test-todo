import axios from 'axios'

const apiClient = axios.create({
    baseURL: 'http://localhost:8080',
    withCredentials: false,
    headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json'
    }
})

export default {
    getList() {
        return apiClient.get('/list')
    },
    addItem(title) {
        return apiClient.post('/add',{title: title})
    },
    removeItem(id) {
        return apiClient.delete('/remove/' + id)
    }
}
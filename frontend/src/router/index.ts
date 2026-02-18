import { createRouter, createWebHistory } from 'vue-router'
import QuizList from '../views/QuizList.vue'
import QuizForm from '../views/QuizForm.vue'

const routes = [
    {
        path: '/',
        name: 'QuizList',
        component: QuizList,
    },
    {
        path: '/create',
        name: 'QuizForm',
        component: QuizForm,
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router

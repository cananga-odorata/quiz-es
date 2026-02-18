import axios from 'axios'
import type { Quiz, CreateQuizRequest, ApiResponse } from '../types/quiz'

const api = axios.create({
    baseURL: '/api/v1',
    headers: {
        'Content-Type': 'application/json',
    },
})

export async function getQuizzes(): Promise<Quiz[]> {
    const { data } = await api.get<ApiResponse<Quiz[]>>('/quizzes')
    return data.data
}

export async function createQuiz(req: CreateQuizRequest): Promise<Quiz> {
    const { data } = await api.post<ApiResponse<Quiz>>('/quizzes', req)
    return data.data
}

export async function deleteQuiz(id: string): Promise<void> {
    await api.delete(`/quizzes/${id}`)
}

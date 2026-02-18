export interface Quiz {
    id: string
    question: string
    choice1: string
    choice2: string
    choice3: string
    choice4: string
    display_order: number
}

export interface CreateQuizRequest {
    question: string
    choice1: string
    choice2: string
    choice3: string
    choice4: string
}

export interface ApiResponse<T> {
    success: boolean
    data: T
    error?: {
        code: string
        message: string
    }
}

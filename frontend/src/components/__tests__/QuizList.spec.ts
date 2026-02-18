import { describe, it, expect, vi } from 'vitest'
import { mount } from '@vue/test-utils'
import QuizList from '../../views/QuizList.vue'

// Mock axios since QuizList uses it

// Mock the router
const mockRouterPush = vi.fn()
vi.mock('vue-router', () => ({
    useRouter: () => ({
        push: mockRouterPush
    })
}))

// Mock API (named exports)
const mockGetQuizzes = vi.fn().mockResolvedValue([])
const mockDeleteQuiz = vi.fn()

vi.mock('../../api/quiz', () => ({
    getQuizzes: () => mockGetQuizzes(),
    deleteQuiz: () => mockDeleteQuiz()
}))

describe('QuizList', () => {
    it('renders correctly', () => {
        const wrapper = mount(QuizList, {
            global: {
                stubs: ['router-link']
            }
        })
        expect(wrapper.exists()).toBe(true)
        expect(wrapper.find('h1').text()).toContain('IT 08-1')
    })

    it('displays quizzes when API returns data', async () => {
        const quizzes = [
            { id: '1', question: 'Q1', choices: {}, display_order: 1 },
            { id: '2', question: 'Q2', choices: {}, display_order: 2 }
        ]

        // Mock the implementation to return data
        mockGetQuizzes.mockResolvedValue(quizzes)

        const wrapper = mount(QuizList)

        // Wait for API call and DOM update
        await new Promise(resolve => setTimeout(resolve, 0))
        await wrapper.vm.$nextTick()
        await wrapper.vm.$nextTick() // Extra tick for v-for rendering

        // Check if quizzes are rendered
        expect(wrapper.text()).toContain('Q1')
        expect(wrapper.text()).toContain('Q2')
    })
})

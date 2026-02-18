<template>
  <div class="page-container">
    <!-- Header -->
    <header class="page-header">
      <h1>IT 08-1</h1>
    </header>

    <!-- Content -->
    <main class="page-content">
      <!-- Add Quiz Button -->
      <div class="action-bar">
        <button class="btn btn-add" @click="goToCreate">เพิ่มข้อสอบ</button>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="loading">กำลังโหลด...</div>

      <!-- Empty State -->
      <div v-else-if="quizzes.length === 0" class="empty-state">
        <p>ยังไม่มีข้อสอบ กดปุ่ม "เพิ่มข้อสอบ" เพื่อเริ่มเพิ่มข้อสอบ</p>
      </div>

      <!-- Quiz List -->
      <div v-else class="quiz-list">
        <div v-for="quiz in quizzes" :key="quiz.id" class="quiz-card">
          <div class="quiz-header">
            <span class="quiz-number">{{ quiz.display_order }}. {{ quiz.question }}</span>
            <button class="btn btn-delete" @click="handleDelete(quiz.id)">ลบ</button>
          </div>
          <div class="quiz-choices">
            <label v-for="(choice, index) in getChoices(quiz)" :key="index" class="choice-item">
              <input type="radio" :name="'quiz-' + quiz.id" disabled />
              <span>{{ choice }}</span>
            </label>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { getQuizzes, deleteQuiz } from '../api/quiz'
import type { Quiz } from '../types/quiz'

const router = useRouter()
const quizzes = ref<Quiz[]>([])
const loading = ref(true)

const fetchQuizzes = async () => {
  loading.value = true
  try {
    quizzes.value = await getQuizzes()
  } catch (error) {
    console.error('Failed to fetch quizzes:', error)
  } finally {
    loading.value = false
  }
}

const goToCreate = () => {
  router.push('/create')
}

const handleDelete = async (id: string) => {
  try {
    await deleteQuiz(id)
    await fetchQuizzes()
  } catch (error) {
    console.error('Failed to delete quiz:', error)
  }
}

const getChoices = (quiz: Quiz): string[] => {
  return [quiz.choice1, quiz.choice2, quiz.choice3, quiz.choice4]
}

onMounted(fetchQuizzes)
</script>

<style scoped>
.page-container {
  max-width: 800px;
  margin: 0 auto;
  background: #fff;
  min-height: 100vh;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.page-header {
  background-color: #4CAF50;
  color: white;
  padding: 12px 20px;
  text-align: center;
}

.page-header h1 {
  margin: 0;
  font-size: 1.2rem;
  font-weight: 600;
}

.page-content {
  padding: 20px;
}

.action-bar {
  margin-bottom: 20px;
}

.btn {
  border: none;
  padding: 8px 20px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.9rem;
  font-weight: 500;
  transition: opacity 0.2s;
}

.btn:hover {
  opacity: 0.85;
}

.btn-add {
  background-color: #4CAF50;
  color: white;
}

.btn-delete {
  background-color: #f44336;
  color: white;
  padding: 4px 16px;
  font-size: 0.85rem;
}

.loading {
  text-align: center;
  color: #666;
  padding: 40px 0;
}

.empty-state {
  text-align: center;
  color: #999;
  padding: 40px 0;
  font-size: 0.95rem;
}

.quiz-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.quiz-card {
  border: 1px solid #e0e0e0;
  border-radius: 6px;
  padding: 16px;
  background: #fafafa;
}

.quiz-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 12px;
}

.quiz-number {
  font-weight: 600;
  color: #d32f2f;
  font-size: 0.95rem;
}

.quiz-choices {
  display: flex;
  flex-direction: column;
  gap: 6px;
  padding-left: 8px;
}

.choice-item {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 0.9rem;
  color: #333;
  cursor: default;
}

.choice-item input[type="radio"] {
  margin: 0;
}
</style>

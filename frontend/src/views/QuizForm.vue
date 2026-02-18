<template>
  <div class="page-container">
    <!-- Header -->
    <header class="page-header">
      <h1>IT 08-2</h1>
    </header>

    <!-- Content -->
    <main class="page-content">
      <form @submit.prevent="handleSubmit" class="quiz-form">
        <div class="form-group">
          <label for="question">คำถาม</label>
          <input
            id="question"
            v-model="form.question"
            type="text"
            placeholder="กรอกคำถาม"
            required
          />
        </div>

        <div class="form-group">
          <label for="choice1">คำตอบ 1</label>
          <input
            id="choice1"
            v-model="form.choice1"
            type="text"
            placeholder="กรอกคำตอบที่ 1"
            required
          />
        </div>

        <div class="form-group">
          <label for="choice2">คำตอบ 2</label>
          <input
            id="choice2"
            v-model="form.choice2"
            type="text"
            placeholder="กรอกคำตอบที่ 2"
            required
          />
        </div>

        <div class="form-group">
          <label for="choice3">คำตอบ 3</label>
          <input
            id="choice3"
            v-model="form.choice3"
            type="text"
            placeholder="กรอกคำตอบที่ 3"
            required
          />
        </div>

        <div class="form-group">
          <label for="choice4">คำตอบ 4</label>
          <input
            id="choice4"
            v-model="form.choice4"
            type="text"
            placeholder="กรอกคำตอบที่ 4"
            required
          />
        </div>

        <div class="form-actions">
          <button type="submit" class="btn btn-save" :disabled="submitting">
            {{ submitting ? 'กำลังบันทึก...' : 'บันทึก' }}
          </button>
          <button type="button" class="btn btn-cancel" @click="handleCancel">ยกเลิก</button>
        </div>
      </form>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { createQuiz } from '../api/quiz'
import type { CreateQuizRequest } from '../types/quiz'

const router = useRouter()
const submitting = ref(false)

const form = reactive<CreateQuizRequest>({
  question: '',
  choice1: '',
  choice2: '',
  choice3: '',
  choice4: '',
})

const handleSubmit = async () => {
  if (submitting.value) return

  submitting.value = true
  try {
    await createQuiz(form)
    router.push('/')
  } catch (error) {
    console.error('Failed to create quiz:', error)
    alert('ไม่สามารถบันทึกข้อสอบได้ กรุณาลองใหม่อีกครั้ง')
  } finally {
    submitting.value = false
  }
}

const handleCancel = () => {
  router.push('/')
}
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
  padding: 30px 40px;
}

.quiz-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-group {
  display: flex;
  align-items: center;
  gap: 16px;
}

.form-group label {
  min-width: 80px;
  font-size: 0.95rem;
  font-weight: 500;
  color: #333;
  text-align: right;
}

.form-group input {
  flex: 1;
  padding: 8px 12px;
  border: 1px solid #ccc;
  border-radius: 4px;
  font-size: 0.9rem;
  outline: none;
  transition: border-color 0.2s;
}

.form-group input:focus {
  border-color: #4CAF50;
}

.form-actions {
  display: flex;
  justify-content: center;
  gap: 12px;
  margin-top: 20px;
}

.btn {
  border: none;
  padding: 8px 28px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.9rem;
  font-weight: 500;
  transition: opacity 0.2s;
}

.btn:hover {
  opacity: 0.85;
}

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.btn-save {
  background-color: #2196F3;
  color: white;
}

.btn-cancel {
  background-color: #f44336;
  color: white;
}
</style>

# 🌟 Quiz Management System (ThaiBev Assignment No.8)

A modern, full-stack application for managing quizzes, built with **Go** (Backend) and **Vue 3** (Frontend).
Designed with Clean Architecture, high-performance infrastructure, and production-ready deployment workflows.

---

## ✨ Key Features

- **Quiz Management**: Create, List, and Delete quiz questions.
- **Automated Reordering**: System automatically re-numbers questions when one is deleted.
- **Rate Limiting**: Protects API with **Token Bucket** algorithm (20 req/s, Burst 5).
- **Modern UI**: Responsive and interactive interface using Vue 3 + Tailwind/CSS.
- **Performance**: Optimized Docker images (~15MB for backend) with Multi-stage builds.

---

## 🛠 Technology Stack

| Layer       | Technology                         | Details |
|-------------|------------------------------------|---------|
| **Backend** | Go 1.24, Chi Router, sqlx          | Clean Architecture (Domain, Repo, Service, Handler) |
| **Frontend** | Vue 3, TypeScript, Vite           | Composition API, Pinia, Tailwind CSS |
| **Database** | PostgreSQL 15                     | Supabase or Local Docker container |
| **Infra**   | Docker Compose, Makefile           | Distroless (Backend), Nginx Alpine (Frontend) |

---

## 🚀 Getting Started

### 🐳 Run with Docker (Recommended)

Quickly start the application using `make` commands:

1.  **Development Mode** (Hot-Reload):
    ```bash
    make up-full
    # Frontend: http://localhost:5173
    # Backend:  http://localhost:8080
    ```

2.  **Production Mode** (Distroless / Nginx):
    ```bash
    make up-prod
    # Web App: http://localhost:80
    # API:     http://localhost:8080
    ```

3.  **Stop Containers**:
    ```bash
    make down       # Dev
    make down-prod  # Prod
    ```

---

## 📡 API Endpoints

- `GET /api/v1/quizzes`: List all quizzes
- `POST /api/v1/quizzes`: Create a new quiz
- `DELETE /api/v1/quizzes/{id}`: Delete a quiz (auto-renumber)

Example `curl` to create a quiz:
```bash
curl -X POST http://localhost:8080/api/v1/quizzes \
  -H "Content-Type: application/json" \
  -d '{
    "question": "What is 2+2?",
    "choice1": "3", "choice2": "4", "choice3": "5", "choice4": "6",
    "answer": 2
  }'
```

---

## 🧪 Testing

The system includes comprehensive tests:

- **Unit Tests (Go)**: `make test`
- **Integration Tests**: `sh tests/integration_test.sh`
- **Load Tests**: `sh tests/load_test.sh` (Tests Rate Limiting)

---

> **Note**: Full architecture documentation and deployment scripts are available internally.
> Created for ThaiBev Assignment.

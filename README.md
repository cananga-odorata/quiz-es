# 🌟 Quiz Management System (ThaiBev Assignment No.8)

A modern, full-stack application for managing quizzes, built with **Go** (Backend) and **Vue 3** (Frontend).
Designed with Clean Architecture, high-performance infrastructure, and production-ready deployment workflows.

---

## 📚 Documentation

- **[Setup & Deployment Guide](./README-Setup.md)**:
    -   Detailed instructions for Docker, Local Development, and Production Deployment.
    -   Includes `setup.sh` usage for fresh servers.
- **[Architecture & Technology Stack](./document.md)**:
    -   Deep dive into Clean Architecture implementation.
    -   Explanation of technology choices (Chi, Sqlx, Composition API, etc.).
    -   Infrastructure details (Multi-stage builds, Distroless images).

---

## ✨ Key Features

- **Quiz Management**: Create, List, and Delete quiz questions.
- **Automated Reordering**: System automatically re-numbers questions when one is deleted.
- **Rate Limiting**: Protects API with Token Bucket algorithm.
- **Modern UI**: Responsive and interactive interface using Vue 3 + Tailwind/CSS.
- **Performance**: Optimized Docker images (~15MB for backend) and fast build times.

## 🛠 Tech Stack

- **Backend**: Go 1.24, Chi Router, PostgreSQL, Clean Architecture
- **Frontend**: Vue 3, TypeScript, Vite, Vitest
- **Infrastructure**: Docker Compose, Makefile, Nginx, Distroless Images

---

> Created for ThaiBev Assignment.

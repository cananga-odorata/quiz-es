# ==========================================
# Quiz Management System — Root Makefile
# ==========================================
# ใช้สั่งงาน Docker Compose จาก root ของโปรเจกต์

COMPOSE_DEV = docker compose -f infra/docker-compose.yml
COMPOSE_PROD = docker compose -f infra/docker-compose.prod.yml
# Default alias
COMPOSE = $(COMPOSE_DEV)

.PHONY: help up-back up-front up-full up-prod down down-prod logs \
        build test migrate clean

# ─────────────────────────────────────────
# Help
# ─────────────────────────────────────────
help:
	@echo ""
	@echo "╔══════════════════════════════════════════╗"
	@echo "║   Quiz Management System — Commands      ║"
	@echo "╚══════════════════════════════════════════╝"
	@echo ""
	@echo "  Docker (Development - Hot Reload):"
	@echo "    make up-back      Backend (Air) + DB"
	@echo "    make up-front     Frontend (Vite)"
	@echo "    make up-full      All Services (Dev)"
	@echo "    make down         Stop Dev containers"
	@echo ""
	@echo "  Docker (Production - Distroless/Nginx):"
	@echo "    make up-prod      Run Full Stack (Prod)"
	@echo "    make down-prod    Stop Prod containers"
	@echo ""
	@echo "  Logs:"
	@echo "    make logs         View logs (Dev)"
	@echo "    make logs-prod    View logs (Prod)"
	@echo ""
	@echo "  Local & Utils:"
	@echo "    make test         Run Unit Tests"
	@echo "    make migrate      Run DB Migration"
	@echo ""
	@echo "  Docker Hub:"
	@echo "    make push-prod    Tag & Push to Docker Hub (Set DOCKER_USER=...)"
	@echo "    make pull-prod    Pull from Docker Hub"
	@echo ""

# ─────────────────────────────────────────
# Development (Hot Reload)
# ─────────────────────────────────────────
up-back:
	$(COMPOSE_DEV) --profile back up -d
	@echo "✅ [DEV] Backend started"
	@echo "   API: http://localhost:8080"

up-front:
	$(COMPOSE_DEV) --profile front up -d
	@echo "✅ [DEV] Frontend started"
	@echo "   App: http://localhost:5173"

up-full:
	$(COMPOSE_DEV) --profile full up -d
	@echo "✅ [DEV] All services started"
	@echo "   API:      http://localhost:8080"
	@echo "   Frontend: http://localhost:5173"

down:
	$(COMPOSE_DEV) --profile full down
	@echo "🛑 [DEV] Containers stopped"

logs:
	$(COMPOSE_DEV) --profile full logs -f

logs-back:
	$(COMPOSE_DEV) --profile back logs -f backend

logs-front:
	$(COMPOSE_DEV) --profile front logs -f frontend

# ─────────────────────────────────────────
# Production (Distroless + Nginx)
# ─────────────────────────────────────────
up-prod:
	$(COMPOSE_PROD) --profile full up -d --build
	@echo "🚀 [PROD] All services started (Distroless/Nginx)"
	@echo "   API:      http://localhost:8080"
	@echo "   Frontend: http://localhost:80"

down-prod:
	$(COMPOSE_PROD) --profile full down
	@echo "🛑 [PROD] Containers stopped"

logs-prod:
	$(COMPOSE_PROD) --profile full logs -f

# ─────────────────────────────────────────
# Docker Build (Manual)
# ─────────────────────────────────────────
build-back:
	docker build -t quiz-api:latest backend/

build-front:
	docker build -t quiz-frontend:latest frontend/

build-all:
	make build-back
	make build-front

# ─────────────────────────────────────────
# Docker Hub (Push/Pull)
# ─────────────────────────────────────────
# Usage: make push-prod DOCKER_USER=myuser VERSION=v1.0
DOCKER_USER ?= yourusername
VERSION ?= latest

tag-prod:
	docker tag quiz-api:latest $(DOCKER_USER)/quiz-api:$(VERSION)
	docker tag quiz-frontend:latest $(DOCKER_USER)/quiz-frontend:$(VERSION)
	@echo "🏷️  Tagged images as $(DOCKER_USER)/...:$(VERSION)"

push-prod: tag-prod
	docker push $(DOCKER_USER)/quiz-api:$(VERSION)
	docker push $(DOCKER_USER)/quiz-frontend:$(VERSION)
	@echo "🚀 Pushed images to Docker Hub"

pull-prod:
	docker pull $(DOCKER_USER)/quiz-api:$(VERSION)
	docker pull $(DOCKER_USER)/quiz-frontend:$(VERSION)
	@echo "📥 Pulled images from Docker Hub"

# ─────────────────────────────────────────
# Local Development (No Docker)
# ─────────────────────────────────────────
run-back:
	cd backend && go run cmd/api/main.go

run-front:
	cd frontend && pnpm dev

test:
	cd backend && go test -v -race -coverprofile=coverage.out ./internal/modules/quiz/... ./internal/shared/middleware/...
	@echo "✅ Unit Tests passed"

migrate:
	cd backend && go run cmd/migrate/main.go

# ─────────────────────────────────────────
# Cleanup
# ─────────────────────────────────────────
clean:
	$(COMPOSE_DEV) --profile full down -v --rmi local
	$(COMPOSE_PROD) --profile full down -v --rmi local
	rm -rf backend/tmp backend/bin backend/coverage.*
	@echo "🧹 Cleaned up everything"

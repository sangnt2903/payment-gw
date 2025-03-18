# Default values
PORT ?= 8080
REGISTRY ?= ghcr.io
IMAGE_NAME ?= payment-gw
ENV ?= development
VERSION ?= latest
GIT_SHA = $(shell git rev-parse --short HEAD)

# Docker image tags
IMAGE_LATEST = $(REGISTRY)/$(IMAGE_NAME):latest
IMAGE_SHA = $(REGISTRY)/$(IMAGE_NAME):$(GIT_SHA)

.PHONY: build
build:
	docker build \
		--build-arg PORT=$(PORT) \
		--build-arg APP_ENV=$(ENV) \
		-t $(IMAGE_LATEST) \
		-t $(IMAGE_SHA) \
		.

.PHONY: push
push: build
	docker push $(IMAGE_LATEST)
	docker push $(IMAGE_SHA)

.PHONY: run
run: build
	docker run -p $(PORT):$(PORT) \
		-e APP_ENV=$(ENV) \
		$(IMAGE_LATEST)

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	docker rmi $(IMAGE_LATEST) $(IMAGE_SHA) || true

# Setup development environment
.PHONY: bootstrap
bootstrap: bootstrap-db bootstrap-redis run-migrations

.PHONY: bootstrap-redis
bootstrap-redis:
	docker-compose up -d redis
	@echo "Waiting for Redis to be ready..."
	@sleep 5

.PHONY: bootstrap-db
bootstrap-db:
	@if [ ! -f docker-compose.yml ]; then echo "docker-compose.yml not found"; exit 1; fi
	docker-compose up -d postgres
	@echo "Waiting for PostgreSQL to be ready..."
	@sleep 5
	@until docker-compose exec -T postgres pg_isready -U postgres; do sleep 2; done

.PHONY: run-migrations
run-migrations:
	go run main.go migrate

.PHONY: clean-db
clean-db:
	docker-compose down -v

.PHONY: reset-db
reset-db: clean-db bootstrap

.PHONY: dev swag
dev:
	air

# swagger docs
.PHONY: swag
swag:
	swag init -g cmd/serve.go

# live reloading
.PHONY: install-tools
install-tools:
	go install github.com/cosmtrek/air@latest

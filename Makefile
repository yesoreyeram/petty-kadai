.PHONY: up down logs test test-all compose-cmd

COMPOSE_FILE := infra/docker-compose.yml

# Detect docker compose binary: prefer plugin (`docker compose`), fallback to legacy `docker-compose`.
DC := $(shell if docker compose version >/dev/null 2>&1; then echo "docker compose"; \
             elif command -v docker-compose >/dev/null 2>&1; then echo docker-compose; \
             else echo ""; fi)

compose-cmd:
	@if [ -z "$(DC)" ]; then \
		echo "ERROR: Neither 'docker compose' nor 'docker-compose' is available in PATH."; \
		echo "Please install Docker Desktop or docker-compose."; \
		exit 127; \
	fi

up: compose-cmd
	COMPOSE_FILE=$(COMPOSE_FILE) $(DC) up -d --build

down: compose-cmd
	COMPOSE_FILE=$(COMPOSE_FILE) $(DC) down -v

logs: compose-cmd
	COMPOSE_FILE=$(COMPOSE_FILE) $(DC) logs -f --tail=100

# Run tests for all Go services that define a go.mod
test:
	@set -e; \
	for d in services/*; do \
		if [ -f $$d/go.mod ]; then \
			echo "Running tests in $$d"; \
			( cd $$d && go test ./... ); \
		fi; \
	done

test-all: test

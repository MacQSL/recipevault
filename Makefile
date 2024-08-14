#!make

# ------------------------------------------------------------------------------
# Makefile -- RecipeVault
# ------------------------------------------------------------------------------

-include .env

# Apply the contents of the .env to the terminal, so that the docker-compose file can use them in its builds
export $(shell sed 's/=.*//' .env)

## ------------------------------------------------------------------------------
## Alias Commands
## - Performs logical groups of commands for your convenience
## ------------------------------------------------------------------------------

# Running the docker build
# 1. Run `make env`
# 2. Edit the `.env` file as needed to update variables and secrets
# 3. Run `make web`

## Performs all commands necessary to run all backend docker services (db, migrations, api)
backend: | close build-backend run-backend
web: | close build-app run-app build-backend run-backend

## ------------------------------------------------------------------------------
## Setup/Cleanup Commands
## ------------------------------------------------------------------------------

env: ## Prepares the environment variables used by all project docker containers
	@echo "==============================================="
	@echo "Make: setup - copying env.docker to .env"
	@echo "==============================================="
	@cp -i .docker/env.docker .env

close: ## Closes all project containers
	@echo "==============================================="
	@echo "Make: close - closing Docker containers"
	@echo "==============================================="
	@docker compose -f docker-compose.yml down

clean: ## Closes and cleans (removes) all project containers
	@echo "==============================================="
	@echo "Make: clean - closing and cleaning Docker containers"
	@echo "==============================================="
	@docker compose -f docker-compose.yml down -v --rmi all --remove-orphans

## ------------------------------------------------------------------------------
## Build/Run Backend Commands
## - Builds all backend services (db, api)
## ------------------------------------------------------------------------------

build-backend: ## Builds all backend containers
	@echo "==============================================="
	@echo "Make: build-backend - building backend images"
	@echo "==============================================="
	@docker compose -f docker-compose.yml up -d --build recipevault-db recipevault-migrations recipevault-api

run-backend: ## Runs all backend containers
	@echo "==============================================="
	@echo "Make: run-backend - running backend images"
	@echo "==============================================="
	@docker compose -f docker-compose.yml up -d recipevault-db recipevault-migrations recipevault-api

## ------------------------------------------------------------------------------
## Build/Run App Commands
## - Builds all frontend services (app)
## ------------------------------------------------------------------------------

build-app: ## Builds all frontend containers
	@echo "==============================================="
	@echo "Make: build-frontend - building frontend images"
	@echo "==============================================="
	@docker compose -f docker-compose.yml up -d --build recipevault-app

run-app: ## Runs all frontend containers
	@echo "==============================================="
	@echo "Make: run-frontend - running frontend images"
	@echo "==============================================="
	@docker compose -f docker-compose.yml up -d recipevault-app

## ------------------------------------------------------------------------------
## Test Commands
## ------------------------------------------------------------------------------

test-api: ## Test api
	@echo "==============================================="
	@echo "Make: test-api - running api test suite"
	@echo "==============================================="
	@cd ./api && go test -v ./...

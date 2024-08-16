.PHONY: build run auth deploy destroy

IMAGE_NAME=scrape-line-bot
IMAGE_VERSION=v1.0
FLY_APP_NAME=scrape-line-bot

build:
	@echo "Building Docker image..."
	docker build -t $(IMAGE_NAME):$(IMAGE_VERSION) .

run:
	@echo "Running Docker image..."
	docker compose up --build

format:
	@echo "Formatting code..."
	go fmt ./...

test:
	@echo "Testing Docker image..."
	docker compose up --build -d
# testing command execution
	docker compose down

auth:
	@echo "Authenticating to fly.io..."
	fly auth login

setup:
	@echo "Creating fly.io app..."
	fly apps create $(FLY_APP_NAME)

setenv:
	@echo "Setting environment variables..."
	./scripts/flyio_set_env.sh

deploy:
	@echo "Deploying to fly.io..."
	fly deploy -a $(FLY_APP_NAME)

app_start: setup setenv deploy

destroy:
	@echo "Destroying fly.io app..."
	fly destroy $(FLY_APP_NAME) -y

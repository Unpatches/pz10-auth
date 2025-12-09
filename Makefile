APP_NAME       = pz10-auth
APP_DIR        = ./cmd/server
BINARY_DIR     = ./bin
BINARY         = $(BINARY_DIR)/$(APP_NAME)
DOCKER_IMAGE   = pz10-auth:latest
CONTAINER_NAME = pz10-auth

.PHONY: run build test clean docker-build docker-run docker-stop docker-logs

## Локальный запуск приложения
run:
	go run $(APP_DIR)

## Сборка бинарного файла
build:
	mkdir -p $(BINARY_DIR)
	go build -o $(BINARY) $(APP_DIR)

## Прогон модульных тестов
test:
	go test ./...

## Очистка артефактов сборки
clean:
	rm -rf $(BINARY_DIR)

## Сборка Docker-образа
docker-build:
	docker build -t $(DOCKER_IMAGE) .

## Запуск контейнера на порту 8080 с использованием .env
docker-run:
	docker run -d --name $(CONTAINER_NAME) \
		--env-file .env \
		-p 8082:8080 \
		$(DOCKER_IMAGE)

## Остановка и удаление контейнера
docker-stop:
	-docker stop $(CONTAINER_NAME)
	-docker rm $(CONTAINER_NAME)

## Просмотр логов контейнера
docker-logs:
	docker logs -f $(CONTAINER_NAME)

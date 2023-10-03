.PHONY: clean test security build run

APP_NAME = go-fiber-htmx-boilerplate
BUILD_DIR = ./build
DB_NAME = postgres
DB_USER = postgres
DB_PASS = postgres
DATABASE_URL = postgres://$(DB_USER):$(DB_PASS)@localhost/$(DB_NAME)?sslmode=disable

clean:
	rm -rf $(BUILD_DIR)/*
	rm -rf *.out

swag:
	swag init

build: swag clean
	CGO_ENABLED=0  go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: build
	$(BUILD_DIR)/$(APP_NAME)
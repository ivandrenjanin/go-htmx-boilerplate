.PHONY: clean test security build run

APP_NAME = go-echo-htmx-boilerplate
BUILD_DIR = ./bin

clean:
	rm -rf $(BUILD_DIR)/*
	rm -rf *.out

tailwind:
	npm run build

build: tailwind
	CGO_ENABLED=0  go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) ./cmd/main/main.go

run: build
	$(BUILD_DIR)/$(APP_NAME)
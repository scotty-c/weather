.PHONY: build clean test release

# Set the project name variable
PROJECT_NAME = weather

build:
	@echo "Building $(PROJECT_NAME)..."
	@go build -o $(PROJECT_NAME) main.go

clean:
	@echo "Cleaning up..."
	@rm -f $(PROJECT_NAME)

test-local:
	@echo "Running tests..."

	go test -v ./...

test:
	@echo "Running tests..."
	@if [ "$(GOARCH)" = "amd64" ] && [ "$(GOOS)" = "linux" ]; then \
		echo "Running tests..."; \
		go test -v ./...; \
	else \
		echo "Skipping tests for cross-compilation..."; \
	fi

release: test
	@echo "Building release..."
	@VERSION=$$(git tag --sort=-v:refname | head -n 1); \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/$(PROJECT_NAME)-$$VERSION-linux-amd64 main.go; \
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./bin/$(PROJECT_NAME)-$$VERSION-linux-arm64 main.go; \
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/$(PROJECT_NAME)-$$VERSION-darwin-amd64 main.go; \
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ./bin/$(PROJECT_NAME)-$$VERSION-darwin-amd64 main.go	
.PHONY: build clean release

build:
	@echo "Building weather..."
	@go build -o weather main.go

clean:
	@echo "Cleaning up..."
	@rm -f weather

release:
	@echo "Building release..."
	@VERSION=$$(git tag --sort=-v:refname | head -n 1  ); \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/weather-$$VERSION-linux-amd64 main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./bin/weather-$$VERSION-linux-arm64 main.go
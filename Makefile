.PHONY: build clean

build:
	@echo "Building weather..."
	@go build -o weather main.go


clean:
	@echo "Cleaning up..."
	@rm -f weather
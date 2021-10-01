.PHONY = all clean
all: build

build: main.go
	@echo "Building binary..."
	go build -o eKYC-service-gin.o

run: 
	docker-compose up -d
	go run .

clean:
	@echo "Cleaning up..."
	rm eKYC-service-gin
	go clean
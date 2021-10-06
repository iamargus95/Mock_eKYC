.PHONY = all clean
all: build

build: main.go
	@echo "Building binary..."
	go build -o eKYC-service-gin.o

run: 
	docker-compose up -d --remove-orphans
	go run .

clean:
	@echo "Cleaning up..."
	docker-compose down
	rm eKYC-service-gin.o
	go clean
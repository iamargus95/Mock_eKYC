.PHONY = all clean
all: build

build: main.go
	@echo "Building binary..."
	go build -o eKYC-service-gin.o

run: 
	make build
	docker-compose up -d --remove-orphans
	./eKYC-service-gin.o

clean:
	@echo "Cleaning up..."
	docker-compose down
	rm eKYC-service-gin.o
	go clean
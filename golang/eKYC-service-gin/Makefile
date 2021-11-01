.PHONY = all clean
all: run

build: main.go
	@echo "Building binary..."
	go build -o eKYC-service-gin.o
	clear

run: 
	make build
	docker-compose up -d --remove-orphans
	./eKYC-service-gin.o

test:
	docker-compose up -d
	go test -v ./...

clean:
	@echo "Cleaning up..."
	docker-compose down
	rm eKYC-service-gin.o
	go clean
	clear
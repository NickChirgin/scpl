build: 
		go build -o bin/main cmd/main.go

run: 
		go mod tidy 
		go run cmd/main.go
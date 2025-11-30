run:
	go run cmd/server/main.go

run-compiled:
	./build/main

compile:
	go build -o build cmd/server/main.go
default: run
run:
	go run main.go
build:
	GOOS=js GOARCH=wasm go build -o main.wasm main.go
build:
	@go build -o fiber_app.exe

run:
	@go run main.go

test:
	@go test -v ./tests
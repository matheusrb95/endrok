client:
	@go run cmd/client/*.go

server:
	@go run cmd/server/*.go

build/exe:
	@CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./bin ./cmd/client/*.go

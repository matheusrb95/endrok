run:
	@go run cmd/client/*.go

client:
	@go run cmd/client/*.go

server:
	@go run cmd/server/*.go

build:
	@CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/game.exe ./cmd/client/*.go
	@zip -j game.zip bin/game.exe
	@zip -r game.zip assets

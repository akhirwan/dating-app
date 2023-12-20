run:
	@go run cmd/main.go

build:
	cd cmd && env GOARCH=amd64 GOOS=linux go build -v -ldflags="-s -w" -o dating-app main.go && upx dating-app
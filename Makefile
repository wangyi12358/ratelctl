build:
	go build -ldflags="-s -w" ratelctl.go

test:
	go test ./...
.PHONY: build

build:
	GOOS=darwin GOARCH=amd64 go build -o dspipe ./cmd/pipe

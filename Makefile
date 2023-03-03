.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux cd cmd && go build -a -installsuffix cgo -ldflags '-w' -o gpt-chat main.go

.PHONY: docker
docker:
	docker build . -t gpt-chat:latest

build: go-build docker-build

docker-build:
		docker-compose build

go-build: go-fmt go-get
		go build -o bin/challenge

go-fmt:
		go fmt ./...

go-get:
		go get ./...

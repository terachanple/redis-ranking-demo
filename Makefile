deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

deps-update:
	rm -rf ./vendor
	dep ensure -update

build:
	GOOS=linux GOARCH=amd64 go build -o bin/redis-ranking-demo main.go

clean:
	rm -rf ./bin
	rm -rf ./vendor

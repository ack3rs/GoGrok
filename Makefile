PROJECT = GoGrok
IMAGE_NAME = GoGrok
VERSION := $(shell git describe --tag --abbrev=0)
NEXT_VERSION:=$(shell git describe --tags --abbrev=0 | awk -F . '{OFS="."; $$NF+=1; print}')
SHA1 := $(shell git rev-parse HEAD)
NOW := $(shell date -u +'%Y%m%d-%H%M%S')


release: fmt
	@git tag -a $(NEXT_VERSION) -m "Release $(NEXT_VERSION)"
	@git push --all
	@git push --tags

run:
	@go run app/server.go

build:
	go build -o bin/$(PROJECT) main.go

fmt:
	@go mod tidy
	@goimports -w .
	@gofmt -w -s .
	@go clean ./...

test:
	@go test  -v -coverprofile=profile.cov ./...
	@go tool cover -func profile.cov

commit: fmt
	@git add .
	@git commit -a -m "$(m)"
	@git pull
	@git push

watch: fmt
	@gow run main.go

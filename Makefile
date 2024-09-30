include .env

run:
	@go run app/*.go

build:
	@go build -o bin/app app/*.go
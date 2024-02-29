.PHONY: dev-http-service build-proto

build-proto:
	protoc -I=./generated --go_out=./generated ./generated/.proto

dev-http-service:
	cd http-service && air


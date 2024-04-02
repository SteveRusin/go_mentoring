.PHONY: http-service build-proto user-service

build-proto:
	 protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		./generated/users_rpc.proto

http-service:
	cd http-service && air

user-service:
	cd user-management-service && air

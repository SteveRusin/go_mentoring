.PHONY: http-service build-proto user-service

build-proto:
	 protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		./generated/users_rpc.proto

http-service:
	cd http-service && air

user-service:
	cd user-management-service && air
benchmark:
	ab -n 100 -c 10 -p /dev/null http://127.0.0.1:8080/user/image

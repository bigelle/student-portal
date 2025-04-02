.PHONY: generate auth gateway all

generate:
	cd ./proto && \
	protoc --go_out=auth --go_opt=paths=source_relative --go-grpc_out=auth --go-grpc_opt=paths=source_relative auth.proto

auth:
	cd ./auth-service && \
	go run ./cmd/main.go -env="../dev.env" &
	
gateway:
	cd ./api-gateway && \
	go run ./cmd/main.go -env="../dev.env" &
	
build:
	protoc proto/user/user.proto --go_out=plugins=grpc:./
	GOOS=linux GOARCH=amd64 go build
	docker build -t user-service .

run:
	docker run -p 50053:50051 user-service
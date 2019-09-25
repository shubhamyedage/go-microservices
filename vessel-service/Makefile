build:
	protoc proto/vessel/vessel.proto --go_out=plugins=grpc:./
	GOOS=linux GOARCH=amd64 go build
	docker build -t vessel-service .

run:
	docker run -p 50052:50051 vessel-service
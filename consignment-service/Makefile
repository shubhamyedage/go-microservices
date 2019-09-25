build:
	protoc proto/consignment/consignment.proto --go_out=plugins=grpc:./
	# protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/microservices/consignment-service/proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-service .

run:
	docker run -p 50051:50051 consignment-service

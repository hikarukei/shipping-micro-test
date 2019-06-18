build:
	protoc -I. --go_out=plugins=grpc:. \
	  proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build
	sudo docker build -t shipping-service-consignment .
run:
    sudo docker run -p 50051:50051 shipping-service-consignment
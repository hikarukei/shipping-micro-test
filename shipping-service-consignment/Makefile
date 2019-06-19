build:
	# Compile protobuf into Golang code for data interface, go-micro is being used
	protoc -I. --go_out=plugins=micro:. \
        proto/consignment/consignment.proto
    # Build the Golang code for Linux OS and amd64 architecture
	GOOS=linux GOARCH=amd64 go build
	# Build the app into docker image
	sudo docker build -t shipping-service-consignment .

run:
	# run the docker image of this app in detached mode, expose at port 50051
	sudo docker run -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		shipping-service-consignment

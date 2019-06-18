package main

import (
	"context"
	_ "go/build"
	"log"
	"net"
	"sync"

	// Import the generated protobuf code
	pb "github.com/hikarukei/shipping-micro-test/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}

// Repository - Dummy repository to simulate the use of datastore
// of some kind. Will be replaced with real implementation later on
type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

// Create new consignment
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	return consignment, nil
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	repo repository
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, handled by gRPC server
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	// save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	// Return matching the 'Response' message we created in our protobuf definition
	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func main() {
	repo := &Repository{}

	// Set-up our gRPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Register our service with gRPC server, this will tie our
	// Implementation into the auto-generated interface code for our
	// protobuf definition
	pb.RegisterShippingServiceServer(s, &service{repo})

	// Register reflection service on gRPC server
	reflection.Register(s)

	log.Println("Running on port: ", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

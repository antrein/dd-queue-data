package grpc

import (
	"antrein/dd-queue-data/application/common/repository"
	"antrein/dd-queue-data/internal/handler/analytic"
	"antrein/dd-queue-data/model/config"
	"context"

	pb "github.com/antrein/proto-repository/pb/dd"
	"google.golang.org/grpc"
)

type helloServer struct {
	pb.UnimplementedGreeterServer
}

func (s *helloServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func ApplicationDelegate(cfg *config.Config, repo *repository.CommonRepository) (*grpc.Server, error) {
	grpcServer := grpc.NewServer()

	// Hello service
	helloServer := &helloServer{}
	pb.RegisterGreeterServer(grpcServer, helloServer)

	// Analytic service
	analyticServer := analytic.New(repo)
	pb.RegisterAnalyticServiceServer(grpcServer, analyticServer)

	return grpcServer, nil
}

package main

import (
	"context"
	"log"
	"net"

	pb "github.com/sungho-cho/covid-spread.viz/backend/proto"
	"github.com/sungho-cho/covid-spread.viz/backend/utils"
	"google.golang.org/grpc"
)

const port = "9000"

type covidDataServer struct {
	pb.CovidDataServer
}

func (s *covidDataServer) GetActiveCases(ctx context.Context, req *pb.GetActiveCasesRequest) (*pb.GetActiveCasesResponse, error) {
	return &pb.GetActiveCasesResponse{
		NumCases: 5,
	}, nil
}

func main() {
	go utils.Fetch()

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCovidDataServer(grpcServer, &covidDataServer{})

	log.Printf("start gRPC server on %s port", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	log.Println("end of main")
}

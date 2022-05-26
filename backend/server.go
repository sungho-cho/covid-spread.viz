package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"time"

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
func updateData() {
	fmt.Println("Updating data!")
}

func waiter() {
	for {
		// fmt.Println("Waiting...")
		time.Sleep(1 * time.Second)
		updateData()
	}
}

func main() {
	queryDate := time.Date(2022, 5, 22, 0, 0, 0, 0, time.UTC)
	fmt.Printf("%v\n", queryDate)
	fmt.Printf("%v\n", queryDate.AddDate(0, 0, 1))
	go utils.Fetch(queryDate)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCovidDataServer(grpcServer, &covidDataServer{})

	go waiter()
	log.Printf("start gRPC server on %s port", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	log.Println("end of main")
}

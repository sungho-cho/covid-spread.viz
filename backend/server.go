package main

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
	pb "github.com/sungho-cho/covid-spread.viz/backend/protos"
	"github.com/sungho-cho/covid-spread.viz/backend/utils"
	"google.golang.org/grpc"
)

const port = "9090"

type covidDataServer struct {
	pb.CovidDataServer
}

func (s *covidDataServer) GetAllData(ctx context.Context, empty_req *pb.Empty) (*pb.GetAllDataResponse, error) {
	var data []*pb.CountriesData
	lastDate := utils.GetLastDate()
	for date := utils.FirstDate; date.Before(lastDate) || date == lastDate; date = utils.NextDay(date) {
		filePath := utils.GetFilePath(date)
		in, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Println("Error reading file for date", date, ":", err)
			continue
		}
		countriesData := &pb.CountriesData{}
		if err := proto.Unmarshal(in, countriesData); err != nil {
			log.Println("Failed to parse countries data for date", date, ":", err)
			continue
		}
		data = append(data, countriesData)
	}
<<<<<<< HEAD
	log.Println("GetAllData successfully sending proto for:", utils.FirstDate, "~", lastDate)
	return &pb.GetAllDataResponse{
		FirstDate: utils.DateToProto(utils.FirstDate),
		LastDate:  utils.DateToProto(lastDate),
		Data:      data,
=======
	return &pb.GetAllDataResponse{
		Data: data,
>>>>>>> 92508c07b6c65410724d0b2c2c63c691b0fb552e
	}, nil
}

func (s *covidDataServer) GetCountriesData(ctx context.Context, req *pb.GetCountriesDataRequest) (*pb.GetCountriesDataResponse, error) {
	date := time.Date(int(req.Date.Year), time.Month(req.Date.Month), int(req.Date.Day), 0, 0, 0, 0, time.UTC)
	filePath := utils.GetFilePath(date)
	in, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("Error reading file:", err)
		return &pb.GetCountriesDataResponse{}, nil
	}
	countriesData := &pb.CountriesData{}
	if err := proto.Unmarshal(in, countriesData); err != nil {
		log.Println("Failed to parse countries data:", err)
		return &pb.GetCountriesDataResponse{}, nil
	}

	log.Println("GetCountriesData successfully sending proto for:", date)
	return &pb.GetCountriesDataResponse{
		CountriesData: countriesData,
	}, nil
}

func (s *covidDataServer) GetMostRecentDate(ctx context.Context, empty_req *pb.Empty) (*pb.Date, error) {
	lastDate := utils.GetLastDate()
	log.Println("GetMostRecentDate successfully sending proto:", lastDate)
	return utils.DateToProto(lastDate), nil
}

func main() {
	// Fetch covid data for all countries and save the data locally
	go utils.Fetch()

	// Set up a gRPC server
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
<<<<<<< HEAD
	grpcServer := grpc.NewServer(grpc.MaxRecvMsgSize(1024*1024*10), grpc.MaxSendMsgSize(1024*1024*10))
=======
	grpcServer := grpc.NewServer(grpc.MaxRecvMsgSize(1024*1024*20), grpc.MaxSendMsgSize(1024*1024*20))
>>>>>>> 92508c07b6c65410724d0b2c2c63c691b0fb552e
	pb.RegisterCovidDataServer(grpcServer, &covidDataServer{})

	// Start the gRPC server
	log.Printf("start gRPC server on %s port", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

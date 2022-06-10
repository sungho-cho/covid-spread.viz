package main

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	"github.com/golang/protobuf/proto"
	pb "github.com/sungho-cho/covid-spread.viz/backend/protos"
	"github.com/sungho-cho/covid-spread.viz/backend/utils"
	"google.golang.org/grpc"
)

type covidDataServer struct {
	pb.CovidDataServer
}

func (s *covidDataServer) GetAllData(ctx context.Context, empty_req *pb.Empty) (*pb.GetAllDataResponse, error) {
	var data []*pb.CountriesData
	lastDate := utils.GetLastDate()
	firstDate := utils.PreviousDay(utils.FirstDate)
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
		// add empty data for the day before DB's first date
		if len(data) == 0 {
			data = append(data, generateEmptyData(firstDate, countriesData))
		}
		data = append(data, countriesData)
	}
	log.Println("GetAllData successfully sending proto for:", firstDate, "~", lastDate)
	return &pb.GetAllDataResponse{
		FirstDate: utils.DateToProto(firstDate),
		LastDate:  utils.DateToProto(lastDate),
		Data:      data,
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

func generateEmptyData(firstDate time.Time, fullData *pb.CountriesData) *pb.CountriesData {
	var countries []*pb.CountryData
	for _, country := range fullData.Countries {
		emptyCountry := &pb.CountryData{
			Country:   country.Country,
			Iso3S:     country.Iso3S,
			Date:      utils.DateToProto(firstDate),
			Confirmed: 0,
			Recovered: 0,
			Deaths:    0,
		}
		countries = append(countries, emptyCountry)
	}
	return &pb.CountriesData{Date: utils.DateToProto(firstDate), Countries: countries}
}

func main() {
	// Fetch covid data for all countries and save the data locally
	go utils.Fetch()

	// Set up a gRPC server
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
		log.Printf("Defaulting to port %s", port)
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.MaxRecvMsgSize(1024*1024*10), grpc.MaxSendMsgSize(1024*1024*10))
	pb.RegisterCovidDataServer(grpcServer, &covidDataServer{})

	// Start the gRPC server
	log.Printf("Starting gRPC server on %s port", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}

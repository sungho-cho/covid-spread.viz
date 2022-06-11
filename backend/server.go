package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	pb "github.com/sungho-cho/covid-spread.viz/backend/protos"
	"github.com/sungho-cho/covid-spread.viz/backend/utils"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
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

func allowCors(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	resp.Header().Set("Access-Control-Expose-Headers", "grpc-status, grpc-message")
	resp.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, XMLHttpRequest, x-user-agent, x-grpc-web, grpc-status, grpc-message")
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

	grpcServer := grpc.NewServer(grpc.MaxRecvMsgSize(1024*1024*10), grpc.MaxSendMsgSize(1024*1024*10))
	pb.RegisterCovidDataServer(grpcServer, &covidDataServer{})

	// Start the gRPC server
	wrapServer := grpcweb.WrapServer(grpcServer)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		allowCors(resp, req)
		if wrapServer.IsGrpcWebRequest(req) || wrapServer.IsAcceptableGrpcCorsRequest(req) {
			wrapServer.ServeHTTP(resp, req)
		}
	})
	grpcWebServer := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}
	// Start the gRPC-web server
	log.Printf("Starting gRPC server on %s port", port)
	if err := grpcWebServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to serve gRPC web server: %s", err)
	}
}

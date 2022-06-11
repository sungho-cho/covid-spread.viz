package main

import (
	"context"
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
	in, err := utils.ReadObject(utils.GCSObjName)
	if err != nil {
		log.Fatalf("Failed to read all data: %s", err)
	}
	allData := &pb.GetAllDataResponse{}
	if err := proto.Unmarshal(in, allData); err != nil {
		log.Printf("Failed to parse all data: %s", err)
	}
	return allData, nil
}

func (s *covidDataServer) GetCountriesData(ctx context.Context, req *pb.GetCountriesDataRequest) (*pb.GetCountriesDataResponse, error) {
	date := time.Date(int(req.Date.Year), time.Month(req.Date.Month), int(req.Date.Day), 0, 0, 0, 0, time.UTC)
	dateStr := date.Format("2006-01-02")
	if !utils.DoesExist(dateStr) {
		log.Printf("Countries data for %s does not exist", dateStr)
		return &pb.GetCountriesDataResponse{}, nil
	}
	in, err := utils.ReadObject(dateStr)
	if err != nil {
		log.Println("Error reading GCS object for date", dateStr, ":", err)
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

	// Wrap the gRPC server into a gRPC-web server
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

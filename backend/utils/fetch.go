package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/sungho-cho/covid-spread.viz/backend/protos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/proto"
)

// CountriesSummary represents the document structure of documents in the
// 'countries_summary' collection.
type CountriesSummary struct {
	ID   primitive.ObjectID `bson:"_id"`
	UIDS []int32

	// Location:
	CombinedNames []string `bson:"combined_names"`
	County        string
	States        []string
	Country       string
	CountryCodes  []int32  `bson:"country_codes"`
	CountryISO2S  []string `bson:"country_iso2s"`
	CountryISO3S  []string `bson:"country_iso3s"`

	Date time.Time

	// Statistics:
	Confirmed  int32
	Deaths     int32
	Population int32
	Recovered  int32
}

func (cs *CountriesSummary) toProto() *pb.CountryData {
	return &pb.CountryData{
		Country:   cs.Country,
		Iso3S:     cs.CountryISO3S,
		Date:      DateToProto(cs.Date),
		Confirmed: cs.Confirmed,
		Recovered: cs.Recovered,
		Deaths:    cs.Deaths,
	}
}

// Metadata represents (a subset of) the data stored in the metadata
// collection in a single document.
type Metadata struct {
	LastDate time.Time `bson:"last_date"`
}

const mdbURL = "mongodb+srv://readonly:readonly@covid-19.hip2i.mongodb.net/covid19"

var FirstDate time.Time = time.Date(2020, 1, 22, 0, 0, 0, 0, time.UTC)
var FinalDate time.Time = FirstDate // gets incremented

const GCSObjName = "covid_data"

// MongoDB Collections
var countries_summary *mongo.Collection
var metadata *mongo.Collection

func init() {
	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mdbURL))
	if err != nil {
		fmt.Printf("%T\n", err)
		panic(fmt.Sprintf("Error initializing MongoDB Client: %v", err))
	}

	// Get references to the main collections:
	database := client.Database("covid19")
	countries_summary = database.Collection("countries_summary")
	metadata = database.Collection("metadata")
}

// Fetch covid cases data from MongoDB, convert the data to protobuf
// and save them to local directory, to be used by the gRPC service
func Fetch() {
	var countriesDataList []*pb.CountriesData

	// If there's pre-existing data on GCS, read in the data so that we can append to it
	if DoesExist(GCSObjName) {
		in, err := ReadObject(GCSObjName)
		if err != nil {
			log.Fatalf("Failed to read all data: %s", err)
		}
		allData := &pb.GetAllDataResponse{}
		if err := proto.Unmarshal(in, allData); err != nil {
			log.Printf("Failed to parse all data: %s", err)
		}
		countriesDataList = allData.GetData()
		FinalDate = NextDay(DateFromProto(allData.GetLastDate()))
	}

	// Fetch all countries data that is up on MongoDB but not included in GCS-stored object
	// Store the data after the data fetching process is over
	for dbDate := GetLastDate(); FinalDate.Before(dbDate) || FinalDate.Equal(dbDate); FinalDate = NextDay(FinalDate) {
		log.Println("Fetching countries data for:", FinalDate.Format("2006-01-02"))
		countriesData := fetchCountriesSummary(FinalDate)
		countriesDataList = append(countriesDataList, countriesData)
	}
	storeAllData(countriesDataList)

	// Wait for MongoDB to contain new data for the next date
	// Once it's updated, fetch and store the data and continue waiting
	// TODO: Replace this functionality with CI
	for {
		// TODO: Timeout
		dbDate := GetLastDate()
		if FinalDate.Equal(dbDate) {
			log.Println("Fetching countries data for:", FinalDate.Format("2006-01-02"))
			countriesData := fetchCountriesSummary(FinalDate)
			countriesDataList = append(countriesDataList, countriesData)
			storeAllData(countriesDataList)
			FinalDate = NextDay(FinalDate)
		}
	}
}

// Get the date of the most recent metadata on the MongoDB
func GetLastDate() time.Time {
	var meta Metadata
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := metadata.FindOne(ctx, bson.D{}).Decode(&meta); err != nil {
		panic(fmt.Sprintf("Error loading metadata document: %v", err))
	}
	return meta.LastDate
}

func DateToProto(date time.Time) *pb.Date {
	return &pb.Date{
		Year:  int32(date.Year()),
		Month: int32(date.Month()),
		Day:   int32(date.Day()),
	}
}

func DateFromProto(protoDate *pb.Date) time.Time {
	return time.Date(
		int(protoDate.GetYear()),
		time.Month(protoDate.GetMonth()),
		int(protoDate.GetDay()),
		0, 0, 0, 0, time.UTC,
	)
}

func PreviousDay(date time.Time) time.Time {
	return date.AddDate(0, 0, -1)
}

func NextDay(date time.Time) time.Time {
	return date.AddDate(0, 0, 1)
}

// Store marshalled GetAllDataResponse proto to GCS bucket using the given data
func storeAllData(countiresDataList []*pb.CountriesData) {
	// Prepend empty data to countiresDataList
	firstProtoDate := countiresDataList[0].GetDate()
	firstDate := DateFromProto(firstProtoDate)
	emptyData := generateEmptyData(PreviousDay(firstDate), countiresDataList[0])
	countiresDataList = append([]*pb.CountriesData{emptyData}, countiresDataList...)

	// Initialize GetAllDataResponse proto instance
	lastProtoDate := countiresDataList[len(countiresDataList)-1].GetDate()
	allData := &pb.GetAllDataResponse{
		FirstDate: firstProtoDate,
		LastDate:  lastProtoDate,
		Data:      countiresDataList,
	}
	// Save the proto to GCS Bucket
	out, err := proto.Marshal(allData)
	if err != nil {
		log.Fatalf("Failed to encode all data: %s", err)
	}
	if err := WriteObject(GCSObjName, out); err != nil {
		log.Fatalf("Failed to store all data: %s", err)
	} else {
		log.Println("Successfully saved all data")
	}
}

func generateEmptyData(date time.Time, fullData *pb.CountriesData) *pb.CountriesData {
	var countries []*pb.CountryData
	for _, country := range fullData.Countries {
		emptyCountry := &pb.CountryData{
			Country:   country.Country,
			Iso3S:     country.Iso3S,
			Date:      DateToProto(date),
			Confirmed: 0,
			Recovered: 0,
			Deaths:    0,
		}
		countries = append(countries, emptyCountry)
	}
	return &pb.CountriesData{Date: DateToProto(date), Countries: countries}
}

// Fetch data from MongoDB's 'countries_summary' collection
func fetchCountriesSummary(date time.Time) *pb.CountriesData {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	opts := options.Find().SetSort(bson.D{{"country", 1}})
	cur, err := countries_summary.Find(ctx, bson.D{{"date", date}}, opts)
	if err != nil {
		panic(err)
	}
	defer cur.Close(ctx)

	var countries []*pb.CountryData
	for cur.Next(ctx) {
		var cs CountriesSummary
		err := cur.Decode(&cs)
		if err != nil {
			panic(err)
		}
		countries = append(countries, cs.toProto())
	}
	countriesData := &pb.CountriesData{
		Date:      DateToProto(date),
		Countries: countries,
	}
	return countriesData
}

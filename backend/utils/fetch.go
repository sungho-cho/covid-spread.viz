package utils

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/golang/protobuf/proto"
	pb "github.com/sungho-cho/covid-spread.viz/backend/protos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mdbURL = "mongodb+srv://readonly:readonly@covid-19.hip2i.mongodb.net/covid19"

var FirstDate time.Time = time.Date(2020, 1, 22, 0, 0, 0, 0, time.UTC)
var FinalDate time.Time = FirstDate // gets incremented

const protoDir string = "/covid-spread.viz/data/" // is prepended by $HOME

// MongoDB Collections
var countries_summary *mongo.Collection
var metadata *mongo.Collection

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

// Metadata represents (a subset of) the data stored in the metadata
// collection in a single document.
type Metadata struct {
	LastDate time.Time `bson:"last_date"`
}

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
	// Download data for up until three days ago
	for ; FinalDate.Before(threeDaysAgo(time.Now().UTC())); FinalDate = NextDay(FinalDate) {
		fetchAndStoreDataForDate(FinalDate)
	}

	// Check if MongoDB owns updated data for the next date
	// This is run separately from the first loop to avoid
	// accumulative network cost of getLastDate()
	for dbDate := GetLastDate(); FinalDate.Before(dbDate) || FinalDate.Equal(dbDate); FinalDate = NextDay(FinalDate) {
		fetchAndStoreDataForDate(FinalDate)
	}

	// Wait for MongoDB to contain new data for the next date
	// Onece it's updated, fetch and store the data and continue waiting
	// TODO: Replace this functionality with CI
	for {
		dbDate := GetLastDate()
		if FinalDate.Equal(dbDate) {
			fetchAndStoreDataForDate(FinalDate)
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

// Get a corresponding file path for the given date
func GetFilePath(date time.Time) string {
	// Get/greate a directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	fullDir := homeDir + protoDir
	os.MkdirAll(fullDir, os.ModePerm)
	// Get file path
	filePath := fullDir + date.Format("2006-01-02")
	return filePath
}

func PreviousDay(date time.Time) time.Time {
	return date.AddDate(0, 0, -1)
}

func NextDay(date time.Time) time.Time {
	return date.AddDate(0, 0, 1)
}

// Fetch and store country data locally for a certain date
func fetchAndStoreDataForDate(date time.Time) {
	// No need to fetch new data if local copy already exists
	if doesCountriesDataExist(date) {
		log.Println("Skipping download for pre-existing countries data for:", date.Format("2006-01-02"))
		return
	}
	countriesData := fetchCountriesSummary(date)
	storeCountriesData(countriesData, date)
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

// Return true if local countries data exist for the given date; false otherwise
func doesCountriesDataExist(date time.Time) bool {
	filePath := GetFilePath(date)
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

// Save the protobuf instance to a local directory
func storeCountriesData(countriesData *pb.CountriesData, date time.Time) {
	// Save a marshalled countries proto to a file
	filePath := GetFilePath(date)
	out, err := proto.Marshal(countriesData)
	if err != nil {
		log.Fatalln("Failed to encode countries data:", err)
	}
	if err := ioutil.WriteFile(filePath, out, 0644); err != nil {
		log.Fatalln("Failed to write countries data:", err)
	} else {
		log.Println("Successfully saved countries data for:", date.Format("2006-01-02"))
	}
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

func threeDaysAgo(date time.Time) time.Time {
	return date.AddDate(0, 0, -3)
}

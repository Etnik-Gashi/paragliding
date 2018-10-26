package main

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// *** DB METHODS *** //

func mongoConnect() *mongo.Client {
	// Connect to MongoDB
	conn, err := mongo.Connect(context.Background(), "mongodb://localhost:27017", nil)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return conn
}

// Check if the track already exists in the database
func urlInMongo(url string, trackColl *mongo.Collection) bool {

	// Read the documents where the trackurl field is equal to url parameter
	cursor, err := trackColl.Find(context.Background(),
		bson.NewDocument(bson.EC.String("url", url)))
	if err != nil {
		log.Fatal(err)
	}

	// 'Close' the (cursor A pointer to the result set of a query. Clients can iterate through a cursor to retrieve results).
	defer cursor.Close(context.Background())

	track := tracks{}

	// Point the cursor at whatever is found
	for cursor.Next(context.Background()) {
		err = cursor.Decode(&track)
		if err != nil {
			log.Fatal(err)
		}
	}

	if track.Url == "" { // If there is an empty field, in this case, `url`, it means the track is not on the database
		return false
	}
	return true
}

// Get track
func getTrack(client *mongo.Client, url string) tracks {
	db := client.Database("igcFiles")     // `paragliding` Database
	collection := db.Collection("tracks") // `track` Collection

	cursor, err := collection.Find(context.Background(), bson.NewDocument(bson.EC.String("url", url)))

	if err != nil {
		log.Fatal(err)
	}

	resTrack := tracks{}

	for cursor.Next(context.Background()) {
		err := cursor.Decode(&resTrack)
		if err != nil {
			log.Fatal(err)
		}
	}

	return resTrack

}

func getTrack1(client *mongo.Client, id string, w http.ResponseWriter) tracks {
	db := client.Database("igcFiles")     // `paragliding` Database
	collection := db.Collection("tracks") // `track` Collection
	filter := bson.NewDocument(bson.EC.String("uniqueid", ""+id+""))
	resTrack := tracks{}
	err := collection.FindOne(context.Background(), filter).Decode(&resTrack)
	if err != nil {
		http.Error(w, "File not found!", 404)
	}
	return resTrack

}

//FUnction that returns ids from database

func getTrackID(client *mongo.Client) string {

	db := client.Database("igcFiles")     // `paragliding` Database
	collection := db.Collection("tracks") // `track` Collection

	cursor, err := collection.Find(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	resTrack := tracks{}
	length, error := collection.Count(context.Background(), nil)
	if error != nil {
		log.Fatal(error)
	}
	ids := "["
	i := int64(0)
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&resTrack)
		if err != nil {
			log.Fatal(err)
		}
		ids += resTrack.UniqueID
		if i == length-1 {
			break
		}
		ids += ","
		i++
	}
	ids += "]"
	return ids
}

// Delete all tracks
func deleteAllTracks(client *mongo.Client) {
	db := client.Database("igcFiles")
	collection := db.Collection("tracks")

	// Delete the tracks
	collection.DeleteMany(context.Background(), bson.NewDocument())
}

// *** URANI MONGO *** //

// *** URANI MONGO *** //

// *** URANI MONGO *** //
// Count all tracks
func countAllTracks(client *mongo.Client) int64 {
	db := client.Database("igcFiles")
	collection := db.Collection("tracks")

	// Count the tracks
	count, _ := collection.Count(context.Background(), nil, nil)

	return count
}

// Return track names
// And also t_stop track

func returnTracks(n int) (string, time.Time) {
	var response string
	var tStop time.Time

	conn := mongoConnect()

	resultTracks := getAllTracks(conn)

	for key, val := range resultTracks { // Go through the slice
		response += `"` + val.UniqueID + `",`
		if key == n-1 || key == len(resultTracks)-1 {
			tStop = val.TimeRecorded
			break
		}
	}

	// Get rid of that last `,` of JSON will freak out
	response = strings.TrimRight(response, ",")

	return response, tStop
}

// ObjectID used in MongoDB
type ObjectID [12]byte

// Counter struct
type Counter struct {
	ID      objectid.ObjectID `bson:"_id"`
	Counter int               `bson:"counter"`
}

// Get trackName from URL
func trackNameFromURL(url string, trackColl *mongo.Collection) string {
	// Get the trackName
	cursor, err := trackColl.Find(context.Background(),
		bson.NewDocument(bson.EC.String("trackurl", url)))

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	dbResult := tracks{}

	for cursor.Next(context.Background()) {
		err = cursor.Decode(&dbResult)
		if err != nil {
			log.Fatal(err)
		}
	}

	return dbResult.Url
}

// Get track counter from DB
func getTrackCounter(db *mongo.Database) int {
	counter := db.Collection("counter") // `counter` Collection

	cursor, err := counter.Find(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	resCounter := Counter{}

	for cursor.Next(context.Background()) {
		err := cursor.Decode(&resCounter)
		if err != nil {
			log.Fatal(err)
		}
	}
	return resCounter.Counter
}

// Increase the track counter
func increaseTrackCounter(cnt int32, db *mongo.Database) {
	collection := db.Collection("counter") // `counter` Collection

	// This is the way to update the counter field in the document
	// Which is storen in the counter collection
	_, err := collection.UpdateOne(context.Background(), nil,
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$set", bson.EC.Int32("counter", cnt+1)), // Increase the counter by one

		),
	)
	if err != nil {
		log.Fatal(err)
	}
}

// Get all tracks
func getAllTracks(client *mongo.Client) []tracks {
	db := client.Database("igcFiles")     // `paragliding` Database
	collection := db.Collection("tracks") // `track` Collection

	var cursor mongo.Cursor
	var err error

	// If points boolean is true
	// Get the points for the track also
	// Otherwise don't

	cursor, err = collection.Find(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	resTracks := []tracks{}
	resTrack := tracks{}

	for cursor.Next(context.Background()) {
		err := cursor.Decode(&resTrack)
		if err != nil {
			log.Fatal(err)
		}
		resTracks = append(resTracks, resTrack) // Append each resTrack to resTracks slice
	}

	return resTracks
}

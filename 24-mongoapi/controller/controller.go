package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Dibyendu-dev/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017/"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

// connect with mongodb
func init() {
	//client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.Background(),clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb connection success")
	collection = client.Database(dbName).Collection(colName)

	fmt.Println("collection instance is ready to explore..")
}


// mongodb helpers

// insert 1 record

func insertOneMovie(movie model.Netflix){
	inserted , err := collection.InsertOne(context.Background(),movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted 1 movie in db with id:",inserted.InsertedID)

}

// update 1 record
func updateOneMovie(movieId string){
	id , _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	update := bson.M{"$set":bson.M{"watched":true}}

	updated , err :=collection.UpdateOne(context.Background(),filter,update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("updated count",updated.ModifiedCount)
}

// delete 1 record
func deleteOneMovie(movieId string){
	id , _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	deleteCount,err :=collection.DeleteOne(context.Background(),filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted with delete count",deleteCount)
}

// delete all record
func deleteAllMovie() int64{
	deleteResult , err := collection.DeleteMany(context.Background(),bson.D{{}},nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies delete:",deleteResult.DeletedCount)
	return  deleteResult.DeletedCount
}

// get all movies
func getAllMovies() []primitive.M {
	cursor , err := collection.Find(context.Background(),bson.D{{}}) 
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M
	for cursor.Next(context.Background()){
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
		log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return  movies
}

// actual controller

func GetAllMovies(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	allmovies := getAllMovies()
	json.NewEncoder(w).Encode(allmovies)

}

func CreateMovies(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Origin","POST")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Origin","PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAMovie(w http.ResponseWriter , r *http.Request) {
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Origin","DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAllMovies(w http.ResponseWriter , r *http.Request) {
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Origin","DELETE")
	
	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)

}





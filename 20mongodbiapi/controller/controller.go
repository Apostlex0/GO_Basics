package controller

import (
	"context"
	"fmt"

	"fo.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "here,comes,the,mongodb,string,you,get"
const dbName = "netflix"
const collectionName = "watchlist"

var collection *mongo.Collection

//connect with mongodb

func init() {
	// client options
	clientOptions := options.Client().ApplyURI(connectionString)

	//connecet to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	//whenever something is happening on another machine we always have to pass a context , as we are not sure what to pass , we pass TODO() context
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Made")
	collection = client.Database(dbname).collection(collectionName)
}

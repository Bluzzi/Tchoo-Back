package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var databaseConnection *mongo.Client

var Context context.Context
var AuthenticationCollection *mongo.Collection

func Load() {
	//xaname4183@ehstock.com
	//swourire123

	//DB
	//Usr: tchoo
	//Pwd: metafriendtchoo
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://tchoo:metafriendtchoo@cluster0.kmiox.mongodb.net/Tchoo?retryWrites=true&w=majority")
	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	databaseConnection = client
	Context = ctx
	AuthenticationCollection = databaseConnection.Database("Tchoo").Collection("Authentication")

	fmt.Println("[LOAD] - Database loaded.")
}
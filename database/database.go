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
var InvitesCollection *mongo.Collection
var NftDataCollection *mongo.Collection
var PromisesCollection *mongo.Collection

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
	NftDataCollection = databaseConnection.Database("Tchoo").Collection("NftData")
	PromisesCollection = databaseConnection.Database("Tchoo").Collection("Promises")
	InvitesCollection = databaseConnection.Database("Tchoo").Collection("Invites")

	fmt.Println("[LOAD] - Database loaded.")
}
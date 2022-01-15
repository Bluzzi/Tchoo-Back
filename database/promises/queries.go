package promises

import (
	"MetaFriend/database"
	"MetaFriend/database/nft"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func GetExecutablePromises() []DatabaseEntry {
	cursor, _ := database.PromisesCollection.Find(
		database.Context,
		bson.M{
			FieldExecutionTimestamp: bson.M{ "$lt": time.Now().Unix() },
		},
	)

	var entries []DatabaseEntry
	_ = cursor.All(database.Context, &entries)

	return entries
}

func GetPromiseByIdentifiers(identifier string) []DatabaseEntry {
	cursor, _ := database.PromisesCollection.Find(
		database.Context,
		bson.M{
			FieldIdentifier: identifier,
		},
	)

	var entries []DatabaseEntry
	_ = cursor.All(database.Context, &entries)

	return entries
}


func CreateDecrementPromise(petNonce int64, duration time.Duration, decrementValue float64, identifier string) {
	_, _ = database.PromisesCollection.InsertOne(
		database.Context,
		DatabaseEntry{
			PetNonce:           petNonce,
			ExecutionTimestamp: time.Now().Unix() + int64(duration.Seconds()),
			Type:               TypeDecrement,
			Value:              decrementValue,
			Field:              nft.FieldPointsPerHourReal,
			Identifier: 		identifier,
			UniqueIdentifier:   uuid.New().String(),
		},
	)
}

func ResolvePromise(uniqueIdentifier string) {
	_, _ = database.PromisesCollection.DeleteOne(
		database.Context,
		bson.M{ FieldUniqueIdentifier: uniqueIdentifier },
	)
}

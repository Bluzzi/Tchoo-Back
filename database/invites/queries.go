package invites

import (
	"MetaFriend/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)


func Exists(field string, value string) (bool, DatabaseEntry) {
	var result DatabaseEntry
	err := database.InvitesCollection.FindOne(
		database.Context,
		bson.M{
			field: value,
		},
	).Decode(&result)

	if err != nil {
		return err != mongo.ErrNoDocuments, DatabaseEntry{}
	}

	return true, result
}

func UpdateField(operator string, field string, value interface{}, discordId string) {
	_, err := database.InvitesCollection.UpdateOne(
		database.Context,
		bson.M{
			FieldDiscordId: discordId,
		},
		bson.M{
			operator: bson.M{
				field: value,
			},
		},
	)

	if err != nil {
		log.Fatal(err)
	}
}

func Create(discordId string, amount int64) {
	_, err := database.AuthenticationCollection.InsertOne(database.Context, DatabaseEntry{
		DiscordId: discordId,
		Invites:   amount,
	})

	if err != nil {
		log.Fatal(err)
	}
}
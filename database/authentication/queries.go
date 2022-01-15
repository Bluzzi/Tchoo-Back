package authentication

import (
	"MetaFriend/database"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strconv"
	"strings"
	"time"
)

func Exists(field string, value string) (bool, DatabaseEntry) {
	var result DatabaseEntry
	err := database.AuthenticationCollection.FindOne(
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

func Create(username string, password string) string {
	_, err := database.AuthenticationCollection.InsertOne(database.Context, DatabaseEntry{
		Username: username,
		UniqueUsername: strings.ToLower(username),
		Password: password,
		Wallet:   "",
		Token:    []string{},
	})

	fmt.Println("[REGISTER] Registered using " + username + " and hashed password " + password)


	if err != nil {
		log.Fatal(err)
	}

	return GenerateLoginToken(username)
}

func VerifyLogin(username string, password string) bool {
	hasher := sha1.New()
	hasher.Write([]byte(password))

	fmt.Println("[LOGIN] Tried log-in using " + username + " and hashed password " + base64.URLEncoding.EncodeToString(hasher.Sum(nil)))
	var result DatabaseEntry
	err := database.AuthenticationCollection.FindOne(
		database.Context,
		bson.M{
			FieldUsername: username,
			FieldPassword: base64.URLEncoding.EncodeToString(hasher.Sum(nil)),
		},
	).Decode(&result)

	if err != nil {
		return err != mongo.ErrNoDocuments
	}

	fmt.Println("[LOGIN] Log-in using " + username + " and hashed password " + base64.URLEncoding.EncodeToString(hasher.Sum(nil)) + " successful.")
	return true
}

func GenerateLoginToken(username string) string {
	uuidStr := uuid.New().String() + "-" + strconv.FormatInt(int64(float64(time.Now().Second())+(time.Hour*24*30).Seconds()), 10)
	_, err := database.AuthenticationCollection.UpdateOne(
		database.Context,
		bson.M{
			FieldUsername: username,
		},
		bson.M{
			"$push": bson.M{
				FieldTokens: uuidStr,
			},
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	return uuidStr
}

func VerifyToken(token string) bool {
	// token format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx-invalid_timestamp
	destructuredToken := strings.Split(token, "-")
	if len(destructuredToken) < 6 {
		return false
	}

	deadlineTimestamp, _ := strconv.ParseInt(destructuredToken[5], 10, 32)
	if int64(time.Now().Second()) > deadlineTimestamp {
		return false
	}

	e, _ := Exists(FieldTokens, token)
	return e
}

func InvalidateLoginToken(token string) {
	_, err := database.AuthenticationCollection.UpdateOne(
		database.Context,
		bson.M{
			FieldTokens: token,
		},
		bson.M{
			"$pull": bson.M{
				FieldTokens: token,
			},
		},
	)

	if err != nil {
		log.Fatal(err)
	}
}

func UpdateField(operator string, field string, value interface{}, token string) {
	_, err := database.AuthenticationCollection.UpdateOne(
		database.Context,
		bson.M{
			FieldTokens: token,
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
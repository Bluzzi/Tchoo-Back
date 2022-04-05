package nft

import (
	"MetaFriend/database"
	"MetaFriend/database/authentication"
	"MetaFriend/nft"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	Wash = "wash"
	Feed = "feed"
	Pet = "pet"
	Play = "play"
	Sleep = "sleep"
)

// Action Name -> Timeout
var allowedActions = map[string]time.Duration{
	Wash: time.Hour * 24,
	Feed: -1,
	Pet: time.Hour / 4,
	Play: time.Hour * 4,
	Sleep: -1,
}

var actionsTemplate = map[string]int64{
	Wash:   0,
	Feed:   0,
	Pet:    0,
	Play:   0,
	Sleep:  0,
}

// GetActionTimeout Get the timeout to add it to a timestamp to know when they can use it again
func GetActionTimeout(action string) time.Duration {
	timeout, _ := allowedActions[action]
	return timeout
}

// CanPetDoAction returns if the pet can do an action else it returns when it can do it again
func CanPetDoAction(nonce int64, action string) (bool, int64) {
	var entry DatabaseEntry
	_ = database.NftDataCollection.FindOne(
		database.Context,
		bson.M{
			FieldNonce: nonce,
		},
	).Decode(&entry)

	currentAction := entry.ActionsUsed[action]
	return currentAction < time.Now().Unix(), currentAction - time.Now().Unix()
}

// VerifyPetOwner verify that nonce is owned by address
func VerifyPetOwner(address string, nonce int64) bool {
	ownerAddress, verify := nft.NonceAddressCache[nonce]
	if !verify {
		return false
	}

	return address == ownerAddress
}

// GetOwnedPets gets the owned pets by an address
func GetOwnedPets(address string) []int64 {
	nonces, exists := nft.AddressNoncesCache[address]
	if !exists {
		return []int64{}
	}

	return nonces
}

// GetNftData gets the nft data
func GetNftData(nonce int64) DatabaseEntry {
	var entry DatabaseEntry
	_ = database.NftDataCollection.FindOne(
		database.Context,
		bson.M{
			FieldNonce: nonce,
		},
	).Decode(&entry)

	holderAddr :=  nft.NonceAddressCache[entry.Nonce]
	exists, data := authentication.Exists(authentication.FieldWallet, holderAddr)

	if exists {
		entry.HolderUsername = data.Username
	} else {
		entry.HolderUsername = holderAddr
	}

	return entry
}

// GetTopNfts gets the top nfts
func GetTopNfts(min int64, max int64) ([]DatabaseEntry, int64) {
	var entries []DatabaseEntry
	cursor, _ := database.NftDataCollection.Find(
		database.Context,
		bson.M{},
		options.Find().SetSort(bson.M{ FieldPrestigeBalance: -1 }),
		options.Find().SetProjection(bson.D{
			{FieldNonce, 1},
			{FieldPrestigeBalance, 1},
			{FieldTwoDPicture, 1},
			{FieldName, 1},
		}),
		options.Find().SetSkip(min),
		options.Find().SetLimit(max),
	)

	_ = cursor.All(database.Context, &entries)

	totalCount := len(entries)

	if totalCount == 0 {
		return []DatabaseEntry{}, 0
	}

	var returnedEntries []DatabaseEntry

	for _, entry := range entries {
		holderAddr :=  nft.NonceAddressCache[entry.Nonce]
		exists, data := authentication.Exists(authentication.FieldWallet, holderAddr)
		if exists {
			entry.HolderUsername = data.Username
		} else {
			entry.HolderUsername = holderAddr
		}
		returnedEntries = append(returnedEntries, entry)
	}

	return returnedEntries, int64(totalCount)
}

func GetAllNfts() []DatabaseEntry {
	var entries []DatabaseEntry
	cursor, _ := database.NftDataCollection.Find(
		database.Context,
		bson.M{},
		options.Find().SetSort(bson.M{ FieldPrestigeBalance: 1 }),
	)
	_ = cursor.All(database.Context, &entries)

	var returnedEntries []DatabaseEntry

	for _, entry := range entries {
		holderAddr :=  nft.NonceAddressCache[entry.Nonce]
		entry.HolderUsername = holderAddr
		returnedEntries = append(returnedEntries, entry)
	}

	return returnedEntries
}

func InsertNftData(entry DatabaseEntry) {
	entry.ActionsUsed = actionsTemplate
	_, _ = database.NftDataCollection.InsertOne(database.Context, entry)
}

// EditNftActionsUsed Functions to edit nft data timeout
func EditNftActionsUsed(nonce int64, field string, value int64) {
	_, _ = database.NftDataCollection.UpdateOne(
		database.Context,
		bson.M{FieldNonce: nonce},
		bson.M{ "$set": bson.M{
			FieldActionsUsed + "." + field: value,
		}},
	)
}

// EditStat Functions to edit nft data
func EditStat(nonce int64, operator string, field string, value interface{}) {
	_, _ = database.NftDataCollection.UpdateOne(
		database.Context,
		bson.M{FieldNonce: nonce},
		bson.M{ operator: bson.M{ field: value } },
	)
}


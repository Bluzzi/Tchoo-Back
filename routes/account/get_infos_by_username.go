package account

import (
	"MetaFriend/database/authentication"
	"MetaFriend/database/nft"
	"MetaFriend/routes/errors"
	"MetaFriend/routes/responses"
	"encoding/json"
	"net/http"
	"strings"
)

type GetInfosByUsernameRequest struct {
	Username string `json:"username"`
}

func (gIBUR GetInfosByUsernameRequest) Verify() (bool, string) {
	if len(gIBUR.Username) == 0 {
		return false, errors.ErrorEmptyField
	}

	if valid, _ := authentication.Exists(authentication.FieldUniqueUsername, strings.ToLower(gIBUR.Username)); !valid {
		return false, errors.ErrorAccountNotExisting
	}

	return true, ""
}

func HandleGetInfosByUsername(w http.ResponseWriter, r *http.Request)  {
	var getInfosRequest GetInfosByUsernameRequest
	_ = json.NewDecoder(r.Body).Decode(&getInfosRequest)

	if success, errorStr := getInfosRequest.Verify(); !success {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errorStr,
		})
		return
	}

	_, accountData := authentication.Exists(authentication.FieldUniqueUsername, strings.ToLower(getInfosRequest.Username))

	ownedNfts := nft.GetOwnedPets(accountData.Wallet)

	_ = json.NewEncoder(w).Encode(responses.AccountInfosResponse{
		Username: accountData.Username,
		UniqueUsername: accountData.UniqueUsername,
		Wallet: accountData.Wallet,
		IsWhitelisted: accountData.IsWhitelisted,
		OwnedPets: ownedNfts,
	})
}

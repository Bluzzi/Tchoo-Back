package account

import (
	"MetaFriend/database/authentication"
	"MetaFriend/database/nft"
	"MetaFriend/routes/errors"
	"MetaFriend/routes/responses"
	"encoding/json"
	"net/http"
)

type GetInfosByDiscordRequest struct {
	DiscordId string `json:"discord_id"`
}

func (gIBDR GetInfosByDiscordRequest) Verify() (bool, string) {
	if len(gIBDR.DiscordId) == 0 {
		return false, errors.ErrorEmptyField
	}

	if valid, _ := authentication.Exists(authentication.FieldDiscordId, gIBDR.DiscordId); !valid {
		return false, errors.ErrorAccountTokenInvalid
	}

	return true, ""
}

func HandleGetInfosByDiscord(w http.ResponseWriter, r *http.Request)  {
	var getInfosRequest GetInfosByDiscordRequest
	_ = json.NewDecoder(r.Body).Decode(&getInfosRequest)

	if success, errorStr := getInfosRequest.Verify(); !success {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errorStr,
		})
		return
	}

	_, accountData := authentication.Exists(authentication.FieldDiscordId, getInfosRequest.DiscordId)

	ownedNfts := nft.GetOwnedPets(accountData.Wallet)

	_ = json.NewEncoder(w).Encode(responses.AccountInfosResponse{
		Username: accountData.Username,
		UniqueUsername: accountData.UniqueUsername,
		Wallet: accountData.Wallet,
		IsWhitelisted: accountData.IsWhitelisted,
		OwnedPets: ownedNfts,
	})
}

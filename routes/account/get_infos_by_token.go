package account

import (
	"MetaFriend/database/authentication"
	"MetaFriend/database/nft"
	"MetaFriend/routes/errors"
	"MetaFriend/routes/responses"
	"encoding/json"
	"net/http"
)

type GetInfosByTokenRequest struct {
	Token string `json:"token"`
}

func (gIBTR GetInfosByTokenRequest) Verify() (bool, string) {
	if len(gIBTR.Token) == 0 {
		return false, errors.ErrorEmptyField
	}

	if valid := authentication.VerifyToken(gIBTR.Token); !valid {
		return false, errors.ErrorAccountTokenInvalid
	}

	return true, ""
}

func HandleGetInfosByToken(w http.ResponseWriter, r *http.Request)  {
	var getInfosRequest GetInfosByTokenRequest
	_ = json.NewDecoder(r.Body).Decode(&getInfosRequest)

	if success, errorStr := getInfosRequest.Verify(); !success {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errorStr,
		})
		return
	}

	_, accountData := authentication.Exists(authentication.FieldTokens, getInfosRequest.Token)

	ownedNfts := nft.GetOwnedPets(accountData.Wallet)

	_ = json.NewEncoder(w).Encode(responses.AccountInfosResponse{
		Username: accountData.Username,
		UniqueUsername: accountData.UniqueUsername,
		Wallet: accountData.Wallet,
		IsWhitelisted: accountData.IsWhitelisted,
		OwnedPets: ownedNfts,
		DiscordId: accountData.DiscordId,
	})
}

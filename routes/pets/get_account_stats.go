package pets

import (
	"MetaFriend/database/authentication"
	"MetaFriend/database/nft"
	"MetaFriend/routes/errors"
	"MetaFriend/routes/responses"
	"encoding/json"
	"net/http"
	"strings"
)

type GetAccountStatsRequest struct {
	Username string `json:"username"`
}

func (gASR GetAccountStatsRequest) Verify() (bool, string) {
	if len(gASR.Username) == 0 {
		return false, errors.ErrorEmptyField
	}

	if valid, _ := authentication.Exists(authentication.FieldUniqueUsername, strings.ToLower(gASR.Username)); !valid {
		return false, errors.ErrorAccountUsernameInvalid
	}

	return true, ""
}

func HandleGetAccountStats(w http.ResponseWriter, r *http.Request)  {
	var getAccountStatsRequest GetAccountStatsRequest
	_ = json.NewDecoder(r.Body).Decode(&getAccountStatsRequest)

	if success, errorStr := getAccountStatsRequest.Verify(); !success {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errorStr,
		})
		return
	}

	_, accountData := authentication.Exists(authentication.FieldUniqueUsername, strings.ToLower(getAccountStatsRequest.Username))

	if accountData.Wallet == "" {
		_ = json.NewEncoder(w).Encode(responses.AccountStatsResponse{
			Username: accountData.Username,
			Wallet: accountData.Wallet,
			OwnedNfts: []nft.DatabaseEntry{},
			Success: true,
		})
	} else {
		ownedNftsNonces := nft.GetOwnedPets(accountData.Wallet)
		var ownedNfts []nft.DatabaseEntry
		for _, nonce := range ownedNftsNonces {
			ownedNfts = append(ownedNfts, nft.GetNftData(nonce))
		}

		_ = json.NewEncoder(w).Encode(responses.AccountStatsResponse{
			Username: accountData.Username,
			Wallet: accountData.Wallet,
			OwnedNfts: ownedNfts,
			Success: true,
		})
	}

}

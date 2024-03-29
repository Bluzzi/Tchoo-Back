package pets

import (
	"MetaFriend/database/authentication"
	"MetaFriend/database/nft"
	"MetaFriend/routes/errors"
	"MetaFriend/routes/responses"
	"encoding/json"
	"net/http"
)

type GetOwnedRequest struct {
	Token string `json:"token"`
	Wallet string `json:"wallet"`
	IsWalletSupplied bool `json:"is_wallet_supplied"`
}

func (gOR GetOwnedRequest) Verify() (bool, string) {
	if len(gOR.Token) == 0 || len(gOR.Wallet) == 0 {
		return false, errors.ErrorEmptyField
	}

	if valid := authentication.VerifyToken(gOR.Token); !valid {
		return false, errors.ErrorAccountTokenInvalid
	}

	return true, ""
}

func HandleGetOwnedRequest(w http.ResponseWriter, r *http.Request)  {
	var getOwnedRequest GetOwnedRequest
	_ = json.NewDecoder(r.Body).Decode(&getOwnedRequest)

	if success, errorStr := getOwnedRequest.Verify(); !success {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errorStr,
		})
		return
	}

	if getOwnedRequest.IsWalletSupplied {
		_ = json.NewEncoder(w).Encode(responses.GetOwnedResponse{
			Success:         true,
			OwnedNftsNonces: nft.GetOwnedPets(getOwnedRequest.Wallet),
		})
	} else {
		_, entry := authentication.Exists(authentication.FieldTokens, getOwnedRequest.Token)

		if entry.Wallet == "" {
			_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
				Success: false,
				Error:   errors.ErrorAccountNoWalletLinked,
			})
			return
		}

		_ = json.NewEncoder(w).Encode(responses.GetOwnedResponse{
			Success:         true,
			OwnedNftsNonces: nft.GetOwnedPets(entry.Wallet),
		})
	}
}

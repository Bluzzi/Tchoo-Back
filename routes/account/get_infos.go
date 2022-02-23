package account

import (
	"MetaFriend/database/authentication"
	"MetaFriend/routes/errors"
	"MetaFriend/routes/responses"
	"encoding/json"
	"net/http"
)

type GetInfosRequest struct {
	Token string `json:"token"`
}

func (gIR GetInfosRequest) Verify() (bool, string) {
	if len(gIR.Token) == 0 {
		return false, errors.ErrorEmptyField
	}

	if valid := authentication.VerifyToken(gIR.Token); !valid {
		return false, errors.ErrorAccountTokenInvalid
	}

	return true, ""
}

func HandleGetInfos(w http.ResponseWriter, r *http.Request)  {
	var getInfosRequest GetInfosRequest
	_ = json.NewDecoder(r.Body).Decode(&getInfosRequest)

	if success, errorStr := getInfosRequest.Verify(); !success {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errorStr,
		})
		return
	}

	_, accountData := authentication.Exists(authentication.FieldTokens, getInfosRequest.Token)

	_ = json.NewEncoder(w).Encode(responses.AccountInfosResponse{
		Username: accountData.Username,
		UniqueUsername: accountData.UniqueUsername,
		Wallet: accountData.Wallet,
	})
}
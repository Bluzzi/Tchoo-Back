package account

import (
	"MetaFriend/database/authentication"
	"MetaFriend/routes/errors"
	"MetaFriend/routes/responses"
	"encoding/json"
	"net/http"
)

type IsWalletLinkedRequest struct {
	Token string `json:"token"`
}

func (iWLR IsWalletLinkedRequest) Verify() (bool, string) {
	if len(iWLR.Token) == 0 {
		return false, errors.ErrorEmptyField
	}

	if valid := authentication.VerifyToken(iWLR.Token); !valid {
		return false, errors.ErrorAccountTokenInvalid
	}

	return true, ""
}

func HandleIsWalletLinked(w http.ResponseWriter, r *http.Request)  {
	var isWalletLinkedRequest IsWalletLinkedRequest
	_ = json.NewDecoder(r.Body).Decode(&isWalletLinkedRequest)

	if success, errorStr := isWalletLinkedRequest.Verify(); !success {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errorStr,
		})
		return
	}

	_, accountData := authentication.Exists(authentication.FieldTokens, isWalletLinkedRequest.Token)

	_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
		Success: accountData.Wallet != "",
		Error:   "",
	})
}

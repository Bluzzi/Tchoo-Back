package account

import (
	"MetaFriend/database/authentication"
	"MetaFriend/routes/errors"
	"MetaFriend/routes/responses"
	"crypto/sha256"
	"encoding/json"
	"net/http"
	"strings"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (cR CreateRequest) Verify() (bool, string) {
	if len(cR.Username) == 0 || len(cR.Password) == 0 {
		return false, errors.ErrorEmptyField
	}

	if exists, _ := authentication.Exists(authentication.FieldUsername, strings.ToLower(cR.Username)); exists {
		return false, errors.ErrorAccountAlreadyExists
	}

	return true, ""
}

func HandleCreateAccount(w http.ResponseWriter, r *http.Request)  {
	var createAccountRequest CreateRequest
	_ = json.NewDecoder(r.Body).Decode(&createAccountRequest)

	if success, errorStr := createAccountRequest.Verify(); !success {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errorStr,
		})
		return
	}

	token := authentication.Create(createAccountRequest.Username, string(sha256.New().Sum([]byte(createAccountRequest.Password))))
	_ = json.NewEncoder(w).Encode(responses.AccountCreateResponse{
		Success: true,
		Token:   token,
	})
}

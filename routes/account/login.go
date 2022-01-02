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

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (lR LoginRequest) Verify() (bool, string) {
	if len(lR.Username) == 0 || len(lR.Password) == 0 {
		return false, errors.ErrorEmptyField
	}
	if exists, _ := authentication.Exists(authentication.FieldUsername, strings.ToLower(lR.Username)); !exists {
		return false, errors.ErrorInvalidLogin
	}

	if verifiedLogin := authentication.VerifyLogin(lR.Username, string(sha256.New().Sum([]byte(lR.Password)))); !verifiedLogin {
		return false, errors.ErrorInvalidLogin
	}
	return true, ""
}

func HandleLoginAccount(w http.ResponseWriter, r *http.Request)  {
	var loginAccountRequest LoginRequest
	_ = json.NewDecoder(r.Body).Decode(&loginAccountRequest)

	if success, errorStr := loginAccountRequest.Verify(); !success {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errorStr,
		})
		return
	}

	_ = json.NewEncoder(w).Encode(responses.AccountCreateResponse{
		Success: true,
		Token:   authentication.GenerateLoginToken(loginAccountRequest.Username),
	})
}
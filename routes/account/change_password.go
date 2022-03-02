package account

import (
	"MetaFriend/database/authentication"
	"MetaFriend/routes/errors"
	"MetaFriend/routes/responses"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"net/http"
)

type ChangePasswordRequest struct {
	Token string `json:"token"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (cPR ChangePasswordRequest) Verify() (bool, string) {
	if len(cPR.Token) == 0 || len(cPR.OldPassword) == 0 || len(cPR.NewPassword) == 0 {
		return false, errors.ErrorEmptyField
	}

	valid, accountData := authentication.Exists(authentication.FieldTokens, cPR.Token)

	if !valid {
		return false, errors.ErrorAccountTokenInvalid
	}

	hasher := sha1.New()
	hasher.Write([]byte(cPR.OldPassword))
	oldPassword := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	if accountData.Password != oldPassword {
		return false, errors.ErrorPasswordIncorrect
	}

	return true, ""
}

func HandleChangePassword(w http.ResponseWriter, r *http.Request)  {
	var changePasswordRequest ChangePasswordRequest
	_ = json.NewDecoder(r.Body).Decode(&changePasswordRequest)

	if success, errorStr := changePasswordRequest.Verify(); !success {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errorStr,
		})
		return
	}

	hasher := sha1.New()
	hasher.Write([]byte(changePasswordRequest.NewPassword))
	authentication.UpdateField("$set", authentication.FieldPassword, base64.URLEncoding.EncodeToString(hasher.Sum(nil)), changePasswordRequest.Token)

	_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
		Success: true,
		Error:   "",
	})
}

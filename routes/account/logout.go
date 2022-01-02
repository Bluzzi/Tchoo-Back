package account

import (
	"MetaFriend/database/authentication"
	"MetaFriend/routes/errors"
	"MetaFriend/routes/responses"
	"encoding/json"
	"net/http"
)

type LogoutRequest struct {
	Token string `json:"token"`
}

func (lR LogoutRequest) Verify() (bool, string) {
	if len(lR.Token) == 0 {
		return false, errors.ErrorEmptyField
	}

	if verified := authentication.VerifyToken(lR.Token); !verified {
		return false, errors.ErrorAccountTokenInvalid
	}

	return true, ""
}

func HandleLogoutAccount(w http.ResponseWriter, r *http.Request)  {
	var logoutAccountRequest LogoutRequest
	_ = json.NewDecoder(r.Body).Decode(&logoutAccountRequest)

	if success, errorStr := logoutAccountRequest.Verify(); !success {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errorStr,
		})
		return
	}

	authentication.InvalidateLoginToken(logoutAccountRequest.Token)
	_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
		Success: true,
	})
}

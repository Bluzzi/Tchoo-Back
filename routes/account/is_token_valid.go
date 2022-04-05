package account

import (
	"MetaFriend/database/authentication"
	"MetaFriend/routes/errors"
	"MetaFriend/routes/responses"
	"encoding/json"
	"net/http"
)

type IsTokenValidRequest struct {
	Token string `json:"token"`
}

func HandleIsTokenValid(w http.ResponseWriter, r *http.Request)  {
	var isTokenValidRequest IsTokenValidRequest
	_ = json.NewDecoder(r.Body).Decode(&isTokenValidRequest)

	success := authentication.VerifyToken(isTokenValidRequest.Token)
	if success {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: true,
			Error:   "",
		})
	} else {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errors.ErrorAccountTokenInvalid,
		})
	}

}

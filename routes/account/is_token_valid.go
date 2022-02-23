package account

import (
	"MetaFriend/database/authentication"
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

	_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
		Success: authentication.VerifyToken(isTokenValidRequest.Token),
		Error:   "",
	})
}

package pets

import (
	"MetaFriend/database/nft"
	"MetaFriend/routes/errors"
	"MetaFriend/routes/responses"
	"encoding/json"
	"net/http"
)

type GetRequest struct {
	Nonce int64 `json:"nonce"`
}

func HandleGetRequest(w http.ResponseWriter, r *http.Request)  {
	var getRequest GetRequest
	_ = json.NewDecoder(r.Body).Decode(&getRequest)

	nftData := nft.GetNftData(getRequest.Nonce)
	if nftData.Name == "" {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errors.ErrorInvalidNftNonce,
		})
	} else {
		nftData.Success = true
		_ = json.NewEncoder(w).Encode(nftData)
	}
}

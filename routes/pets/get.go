package pets

import (
	"MetaFriend/database/nft"
	"encoding/json"
	"net/http"
)

type GetRequest struct {
	Nonce int64 `json:"nonce"`
}

func HandleGetRequest(w http.ResponseWriter, r *http.Request)  {
	var getRequest GetRequest
	_ = json.NewDecoder(r.Body).Decode(&getRequest)

	_ = json.NewEncoder(w).Encode(nft.GetNftData(getRequest.Nonce))
}

package pets

import (
	"MetaFriend/database/nft"
	"MetaFriend/routes/responses"
	"encoding/json"
	"net/http"
)

type GetTopRequest struct {
	StartIndex int64 `json:"start_index"`
	StopIndex int64 `json:"stop_index"`
}

func HandleGetTopRequest(w http.ResponseWriter, r *http.Request)  {
	var getTopRequest GetTopRequest
	_ = json.NewDecoder(r.Body).Decode(&getTopRequest)

	_ = json.NewEncoder(w).Encode(responses.GetTopResponse{
		Success: true,
		TopNfts: nft.GetTopNfts(getTopRequest.StartIndex, getTopRequest.StopIndex),
	})
}

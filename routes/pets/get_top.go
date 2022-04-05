package pets

import (
	"MetaFriend/database/nft"
	"MetaFriend/routes/errors"
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

	topNfts, totalCount := nft.GetTopNfts(getTopRequest.StartIndex, getTopRequest.StopIndex)
	if totalCount == 0 {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errors.ErrorTopOutOfRange,
		})
		return
	}

	_ = json.NewEncoder(w).Encode(responses.GetTopResponse{
		Success: true,
		TopNfts: topNfts,
		TotalCount: totalCount,
	})
}

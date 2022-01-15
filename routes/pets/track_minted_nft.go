package pets

import (
	"MetaFriend/database/nft"
	"MetaFriend/routes/responses"
	"encoding/json"
	"net/http"
)

type TrackMintedNftRequest struct {
	Token 			  string `json:"token"`
	Nonce 			  int64 `json:"nonce"`
	ThreeDModel 	  string `json:"three_d_model"`
	MtlModel 	  	  string `json:"mtl_model"`
	TextureModel 	  string `json:"texture_model"`
	TwoDPicture 	  string `json:"two_d_picture"`
	Name 			  string `json:"name"`
	PointsBalance     float64 `json:"points_balance"`
	PrestigeBalance   float64 `json:"prestige_balance"`
	PointsPerHourBase float64 `json:"points_per_hour_base"`
	PointsPerHourReal float64 `json:"points_per_hour_real"`
}

func (tMNR TrackMintedNftRequest) Verify() (bool, string) {
	if tMNR.Token != "00ed35b450dc8a87cd7f22ee838c51e85617d6fe2bfae43c92be5884811b3600" {
		return false, "Bro wtf?"
	}
	return true, ""
}

func HandleTrackMintedNftRequest(w http.ResponseWriter, r *http.Request)  {
	var trackMintedNftRequest TrackMintedNftRequest
	_ = json.NewDecoder(r.Body).Decode(&trackMintedNftRequest)

	if success, errorStr := trackMintedNftRequest.Verify(); !success {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errorStr,
		})
		return
	}

	nft.InsertNftData(nft.DatabaseEntry{
		Nonce:             trackMintedNftRequest.Nonce,
		ThreeDModel:       trackMintedNftRequest.ThreeDModel,
		MtlModel: 		   trackMintedNftRequest.MtlModel,
		TextureModel:	   trackMintedNftRequest.TextureModel,
		TwoDPicture:       trackMintedNftRequest.TwoDPicture,
		Name:              trackMintedNftRequest.Name,
		PointsBalance:     trackMintedNftRequest.PointsBalance,
		PrestigeBalance:   trackMintedNftRequest.PrestigeBalance,
		PointsPerHourBase: trackMintedNftRequest.PointsPerHourBase,
		PointsPerHourReal: trackMintedNftRequest.PointsPerHourReal,
	})

	_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
		Success: true,
		Error:   "",
	})
}

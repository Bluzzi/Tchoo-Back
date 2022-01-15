package lottery

import (
	"MetaFriend/config"
	"MetaFriend/database/authentication"
	"MetaFriend/database/nft"
	"MetaFriend/lottery"
	"MetaFriend/routes/errors"
	"MetaFriend/routes/responses"
	"encoding/json"
	"net/http"
)

type BuyTicketRequest struct {
	Token string `json:"token"`
	PetNonce int64 `json:"pet_nonce"`
}

func (bTR BuyTicketRequest) Verify() (bool, string) {
	if len(bTR.Token) == 0 {
		return false, errors.ErrorEmptyField
	}

	if valid := authentication.VerifyToken(bTR.Token); !valid {
		return false, errors.ErrorAccountTokenInvalid
	}

	_, accountData := authentication.Exists(authentication.FieldTokens, bTR.Token)

	if reallyOwned := nft.VerifyPetOwner(accountData.Wallet, bTR.PetNonce); !reallyOwned {
		return false, errors.ErrorPetNotOwned
	}

	return true, ""
}

func HandleBuyTicket(w http.ResponseWriter, r *http.Request)  {
	var buyTicketRequest BuyTicketRequest
	_ = json.NewDecoder(r.Body).Decode(&buyTicketRequest)

	if success, errorStr := buyTicketRequest.Verify(); !success {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errorStr,
		})
		return
	}

	nftData := nft.GetNftData(buyTicketRequest.PetNonce)

	if nftData.PointsBalance < config.Config.PointsForATicket {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errors.ErrorNotEnoughPoints,
		})
		return
	}

	nft.EditStat(
		buyTicketRequest.PetNonce,
		"$inc",
		nft.FieldPointsBalance,
		-config.Config.PointsForATicket,
	)

	_, accountData := authentication.Exists(authentication.FieldTokens, buyTicketRequest.Token)
	prize := lottery.RandomlyGetPrize(accountData.Username)

	_ = json.NewEncoder(w).Encode(responses.LotteryBuyTicketResponse{
		WonPrize:     prize.IsNotAir,
		Prize:        prize.PrizeStr,
		PrizePicture: prize.Picture,
	})
}


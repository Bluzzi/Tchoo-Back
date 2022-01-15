package interactions

import (
	"MetaFriend/database/authentication"
	"MetaFriend/database/nft"
	"MetaFriend/database/promises"
	"MetaFriend/routes/errors"
	"MetaFriend/routes/responses"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type WashPetRequest struct {
	Token string `json:"token"`
	PetNonce int64 `json:"pet_nonce"`
}

func (wPR WashPetRequest) Verify() (bool, string) {
	if len(wPR.Token) == 0 {
		return false, errors.ErrorEmptyField
	}

	if valid := authentication.VerifyToken(wPR.Token); !valid {
		return false, errors.ErrorAccountTokenInvalid
	}

	_, accountData := authentication.Exists(authentication.FieldTokens, wPR.Token)

	if reallyOwned := nft.VerifyPetOwner(accountData.Wallet, wPR.PetNonce); !reallyOwned {
		return false, errors.ErrorPetNotOwned
	}

	if canPet, timeLeft := nft.CanPetDoAction(wPR.PetNonce, nft.Wash); !canPet {
		return false, errors.ErrorPetActionOnCooldown + strconv.FormatInt(timeLeft, 10)
	}

	return true, ""
}

func HandleWashPet(w http.ResponseWriter, r *http.Request)  {
	var washPetRequest WashPetRequest
	_ = json.NewDecoder(r.Body).Decode(&washPetRequest)

	if success, errorStr := washPetRequest.Verify(); !success {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errorStr,
		})
		return
	}

	nftData := nft.GetNftData(washPetRequest.PetNonce)

	//https://github.com/Meta-Friend-Team/Tchoo-Project-Map/blob/main/lang/EN.md -> increment it by 20% (0.2)
	incr := float64(nftData.PointsPerHourBase) * 0.2
	// Decrement in one day
	timeout := time.Hour * 24

	if nftData.PointsPerHourReal + incr > nftData.PointsPerHourBase {
		nft.EditStat(nftData.Nonce, "$set", nft.FieldPointsPerHourReal, nftData.PointsPerHourBase)
	} else {
		nft.EditStat(nftData.Nonce, "$inc", nft.FieldPointsPerHourReal, incr)
	}

	promises.CreateDecrementPromise(nftData.Nonce, timeout, incr, "")
	nft.EditStat(nftData.Nonce, "$inc", nft.FieldPrestigeBalance, 2)
	nft.EditNftActionsUsed(nftData.Nonce, nft.Wash, time.Now().Unix() + int64(timeout.Seconds()))

	_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
		Success: true,
		Error:   "",
	})
}

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

type CaressPetRequest struct {
	Token string `json:"token"`
	PetNonce int64 `json:"pet_nonce"`
}

func (cPR CaressPetRequest) Verify() (bool, string) {
	if len(cPR.Token) == 0 {
		return false, errors.ErrorEmptyField
	}

	if valid := authentication.VerifyToken(cPR.Token); !valid {
		return false, errors.ErrorAccountTokenInvalid
	}

	_, accountData := authentication.Exists(authentication.FieldTokens, cPR.Token)

	if reallyOwned := nft.VerifyPetOwner(accountData.Wallet, cPR.PetNonce); !reallyOwned {
		return false, errors.ErrorPetNotOwned
	}

	// Check if he caressed it more than 4 times
	if caressedTimes := promises.GetPromiseByIdentifiers(nft.Pet + "-" + strconv.FormatInt(cPR.PetNonce, 10)); len(caressedTimes) >= 4 {
		return false, errors.ErrorPetActionOnCooldown
	}

	return true, ""
}

func HandleCaressPet(w http.ResponseWriter, r *http.Request)  {
	var caressPetRequest CaressPetRequest
	_ = json.NewDecoder(r.Body).Decode(&caressPetRequest)

	if success, errorStr := caressPetRequest.Verify(); !success {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errorStr,
		})
		return
	}

	nftData := nft.GetNftData(caressPetRequest.PetNonce)

	//https://github.com/Meta-Friend-Team/Tchoo-Project-Map/blob/main/lang/EN.md -> increment it by 5% (0.05)
	incr := float64(nftData.PointsPerHourBase) * 0.05
	// Decrement in one hour
	timeout := time.Hour

	if float64(nftData.PointsPerHourReal) + incr > float64(nftData.PointsPerHourBase) {
		nft.EditStat(nftData.Nonce, "$set", nft.FieldPointsPerHourReal, nftData.PointsPerHourBase)
	} else {
		nft.EditStat(nftData.Nonce, "$inc", nft.FieldPointsPerHourReal, incr)
	}

	promises.CreateDecrementPromise(nftData.Nonce, timeout, incr, nft.Pet + "-" + strconv.FormatInt(nftData.Nonce, 10))
	nft.EditStat(nftData.Nonce, "$inc", nft.FieldPrestigeBalance, 0.5)

	_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
		Success: true,
		Error:   "",
	})
}

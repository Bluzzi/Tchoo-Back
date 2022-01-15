package interactions

import (
	"MetaFriend/database/authentication"
	"MetaFriend/database/nft"
	"MetaFriend/database/promises"
	"MetaFriend/routes/errors"
	"MetaFriend/routes/responses"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type FeedPetRequest struct {
	Token string `json:"token"`
	PetNonce int64 `json:"pet_nonce"`
	TimeFrame string `json:"time_frame"`
}

func (fPR FeedPetRequest) Verify() (bool, string) {
	if len(fPR.Token) == 0 {
		return false, errors.ErrorEmptyField
	}

	if valid := authentication.VerifyToken(fPR.Token); !valid {
		return false, errors.ErrorAccountTokenInvalid
	}

	_, accountData := authentication.Exists(authentication.FieldTokens, fPR.Token)

	if reallyOwned := nft.VerifyPetOwner(accountData.Wallet, fPR.PetNonce); !reallyOwned {
		return false, errors.ErrorPetNotOwned
	}

	if canDoAction, timeLeft := nft.CanPetDoAction(fPR.PetNonce, nft.Feed); !canDoAction {
		return false, errors.ErrorPetActionOnCooldown + strconv.FormatInt(timeLeft, 10)
	}

	return true, ""
}

func HandleFeedPet(w http.ResponseWriter, r *http.Request)  {
	var feedPetRequest FeedPetRequest
	_ = json.NewDecoder(r.Body).Decode(&feedPetRequest)

	if success, errorStr := feedPetRequest.Verify(); !success {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errorStr,
		})
		return
	}

	nftData := nft.GetNftData(feedPetRequest.PetNonce)

	//https://github.com/Meta-Friend-Team/Tchoo-Project-Map/blob/main/lang/EN.md -> increment it by 20% (0.2)

	localisation, e := time.LoadLocation(feedPetRequest.TimeFrame)
	if e != nil {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errors.ErrorInvalidLocation,
		})
		return
	}

	h := time.Now().In(localisation).Hour()
	fmt.Println("Hours: ", h)
	incr := float64(nftData.PointsPerHourBase) * 0.2

	var timeout time.Duration
	if h >= 11 && h <= 15 {
		// It's noon feed time
		if nftData.PointsPerHourReal+ incr > nftData.PointsPerHourBase {
			nft.EditStat(nftData.Nonce, "$set", nft.FieldPointsPerHourReal, nftData.PointsPerHourBase)
		} else {
			nft.EditStat(nftData.Nonce, "$inc", nft.FieldPointsPerHourReal, incr)
		}

		// Hours until evening time:
		timeout = time.Hour * time.Duration(19 - h) - (time.Minute * time.Duration(time.Now().Minute()))
		// Decrement at the start of evening feed time
	} else if h >= 19 && h <= 23 {
		// It's evening feed time
		if nftData.PointsPerHourReal+ incr > nftData.PointsPerHourBase {
			nft.EditStat(nftData.Nonce, "$set", nft.FieldPointsPerHourReal, nftData.PointsPerHourBase)
		} else {
			nft.EditStat(nftData.Nonce, "$inc", nft.FieldPointsPerHourReal, incr)
		}

		// Hours until noon time:
		timeout = time.Hour * time.Duration(34 - h) - (time.Minute * time.Duration(time.Now().Minute()))
	} else {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errors.ErrorInvalidFeedTime,
		})
		return
	}

	promises.CreateDecrementPromise(nftData.Nonce, timeout, incr, "")
	nft.EditNftActionsUsed(nftData.Nonce, nft.Feed, time.Now().Unix() + int64(timeout.Seconds()))
	nft.EditStat(nftData.Nonce, "$inc", nft.FieldPrestigeBalance, 2)

	_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
		Success: true,
		Error:   "",
	})
}

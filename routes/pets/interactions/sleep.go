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

type SleepPetRequest struct {
	Token string `json:"token"`
	PetNonce int64 `json:"pet_nonce"`
	TimeFrame string `json:"time_frame"`
	IsMorning bool `json:"is_morning"`
}

func (sPR SleepPetRequest) Verify() (bool, string) {
	if len(sPR.Token) == 0 {
		return false, errors.ErrorEmptyField
	}

	if valid := authentication.VerifyToken(sPR.Token); !valid {
		return false, errors.ErrorAccountTokenInvalid
	}

	_, accountData := authentication.Exists(authentication.FieldTokens, sPR.Token)

	if reallyOwned := nft.VerifyPetOwner(accountData.Wallet, sPR.PetNonce); !reallyOwned {
		return false, errors.ErrorPetNotOwned
	}

	if canDoAction, timeLeft := nft.CanPetDoAction(sPR.PetNonce, nft.Sleep); !canDoAction {
		return false, errors.ErrorPetActionOnCooldown + strconv.FormatInt(timeLeft, 10)
	}

	return true, ""
}

func HandleSleepPet(w http.ResponseWriter, r *http.Request)  {
	var sleepPetRequest SleepPetRequest
	_ = json.NewDecoder(r.Body).Decode(&sleepPetRequest)

	if success, errorStr := sleepPetRequest.Verify(); !success {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errorStr,
		})
		return
	}

	nftData := nft.GetNftData(sleepPetRequest.PetNonce)

	//https://github.com/Meta-Friend-Team/Tchoo-Project-Map/blob/main/lang/EN.md -> increment it by 20% (0.2)

	localisation, e := time.LoadLocation(sleepPetRequest.TimeFrame)
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
	if sleepPetRequest.IsMorning {
		if h >= 6 && h <= 12 {
			// It's wake up time
			if float64(nftData.PointsPerHourReal) + incr > float64(nftData.PointsPerHourBase) {
				nft.EditStat(nftData.Nonce, "$set", nft.FieldPointsPerHourReal, nftData.PointsPerHourBase)
			} else {
				nft.EditStat(nftData.Nonce, "$inc", nft.FieldPointsPerHourReal, incr)
			}

			// Hours until sleep time:
			timeout = time.Hour * time.Duration(20 - h) - (time.Minute * time.Duration(time.Now().Minute()))
			// Decrement at the start of evening feed time
		} else {
			_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
				Success: false,
				Error:   errors.ErrorInvalidSleepTime,
			})
			return
		}
	} else {
		if h >= 20 || h <= 2 {
			// It's evening feed time
			if float64(nftData.PointsPerHourReal) + incr > float64(nftData.PointsPerHourBase) {
				nft.EditStat(nftData.Nonce, "$set", nft.FieldPointsPerHourReal, nftData.PointsPerHourBase)
			} else {
				nft.EditStat(nftData.Nonce, "$inc", nft.FieldPointsPerHourReal, incr)
			}

			// Hours until noon time:
			if h <= 2 {
				timeout = time.Hour * time.Duration(6 - h) - (time.Minute * time.Duration(time.Now().Minute()))
			} else {
				timeout = time.Hour * time.Duration(30 - h) - (time.Minute * time.Duration(time.Now().Minute()))
			}
		} else {
			_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
				Success: false,
				Error:   errors.ErrorInvalidSleepTime,
			})
			return
		}
	}

	promises.CreateDecrementPromise(nftData.Nonce, timeout, incr, "")
	nft.EditNftActionsUsed(nftData.Nonce, nft.Sleep, time.Now().Unix() + int64(timeout.Seconds()))
	nft.EditStat(nftData.Nonce, "$inc", nft.FieldPrestigeBalance, 2)

	_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
		Success: true,
		Error:   "",
	})
}


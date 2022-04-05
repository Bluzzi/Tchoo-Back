package account

import (
	"MetaFriend/database/authentication"
	"MetaFriend/routes/errors"
	"MetaFriend/routes/responses"
	"encoding/json"
	"net/http"
)

type SetWhitelistedRequest struct {
	DiscordId string `json:"discord_id"`
	Whitelist bool `json:"whitelist"`
	PrivateKey string `json:"private_key"`
}

func (sWR SetWhitelistedRequest) Verify() (bool, string) {
	if len(sWR.DiscordId) == 0 || len(sWR.PrivateKey) == 0 {
		return false, errors.ErrorEmptyField
	}

	if sWR.PrivateKey != "00ed35b450dc8a87cd7f22ee838c51e85617d6fe2bfae43c92be5884811b3600" {
		return false, "Incorrect private key"
	}

	if valid, _ := authentication.Exists(authentication.FieldDiscordId, sWR.DiscordId); !valid {
		return false, errors.ErrorAccountNotExisting
	}

	return true, ""
}

func HandleSetWhitelisted(w http.ResponseWriter, r *http.Request)  {
	var setWhitelistedRequest SetWhitelistedRequest
	_ = json.NewDecoder(r.Body).Decode(&setWhitelistedRequest)

	if success, errorStr := setWhitelistedRequest.Verify(); !success {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errorStr,
		})
		return
	}

	_, accountData := authentication.Exists(authentication.FieldDiscordId, setWhitelistedRequest.DiscordId)

	authentication.UpdateField("$set", authentication.FieldIsWhitelisted, setWhitelistedRequest.Whitelist, accountData.Token[0])

	_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
		Success: true,
		Error:   "",
	})
}


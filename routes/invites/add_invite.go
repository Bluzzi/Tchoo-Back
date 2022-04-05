package invites

import (
	"MetaFriend/database/invites"
	"MetaFriend/routes/errors"
	"MetaFriend/routes/responses"
	"encoding/json"
	"net/http"
)

type AddInviteRequest struct {
	DiscordId string `json:"discord_id"`
	Amount int64 `json:"amount"`
	PrivateKey string `json:"private_key"`
}

func (aIR AddInviteRequest) Verify() (bool, string) {
	if len(aIR.DiscordId) == 0 || len(aIR.PrivateKey) == 0 {
		return false, errors.ErrorEmptyField
	}

	if aIR.PrivateKey != "00ed35b450dc8a87cd7f22ee838c51e85617d6fe2bfae43c92be5884811b3600" {
		return false, "Incorrect private key"
	}

	return true, ""
}

func HandleAddInvite(w http.ResponseWriter, r *http.Request)  {
	var addInviteRequest AddInviteRequest
	_ = json.NewDecoder(r.Body).Decode(&addInviteRequest)

	if success, errorStr := addInviteRequest.Verify(); !success {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errorStr,
		})
		return
	}

	if valid, _ := invites.Exists(invites.FieldDiscordId, addInviteRequest.DiscordId); !valid {
		invites.UpdateField("$inc", invites.FieldInvites, addInviteRequest.Amount, addInviteRequest.DiscordId)
	} else {
		invites.Create(addInviteRequest.DiscordId, addInviteRequest.Amount)
	}

	_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
		Success: true,
		Error:   "",
	})
}



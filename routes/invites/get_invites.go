package invites

import (
	"MetaFriend/database/invites"
	"MetaFriend/routes/errors"
	"MetaFriend/routes/responses"
	"encoding/json"
	"net/http"
)

type GetInvitesRequest struct {
	DiscordId string `json:"discord_id"`
}

func (gIR GetInvitesRequest) Verify() (bool, string) {
	if len(gIR.DiscordId) == 0 {
		return false, errors.ErrorEmptyField
	}

	if valid, _ := invites.Exists(invites.FieldDiscordId, gIR.DiscordId); !valid {
		return false, errors.ErrorAccountNotExisting
	}

	return true, ""
}

func GetInvites(w http.ResponseWriter, r *http.Request)  {
	var getInvitesRequest GetInvitesRequest
	_ = json.NewDecoder(r.Body).Decode(&getInvitesRequest)

	if success, errorStr := getInvitesRequest.Verify(); !success {
		_ = json.NewEncoder(w).Encode(responses.SuccessResponse{
			Success: false,
			Error:   errorStr,
		})
		return
	}

	_, data := invites.Exists(invites.FieldDiscordId, getInvitesRequest.DiscordId)

	_ = json.NewEncoder(w).Encode(responses.GetInvitesResponse{
		Success: true,
		Invites:   data.Invites,
	})
}


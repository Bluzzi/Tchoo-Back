package invites

type DatabaseEntry struct {
	DiscordId string `json:"discord_id" bson:"discord_id"`
	Invites int64 `json:"invites" bson:"invites"`
}

var (
	FieldDiscordId = "discord_id"
	FieldInvites = "invites"
)
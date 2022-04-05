package authentication

type DatabaseEntry struct {
	Username string `json:"username" bson:"username"`
	UniqueUsername string `json:"unique_username" bson:"unique_username"`
	Password string `json:"password" bson:"password"`
	Wallet string `json:"wallet" bson:"wallet"`
	IsWhitelisted bool `json:"is_whitelisted" bson:"is_whitelisted"`
	DiscordId string `json:"discord_id" bson:"discord_id"`
	Token []string  `json:"tokens" bson:"tokens"`
}

var (
	FieldUsername = "username"
	FieldUniqueUsername = "unique_username"
	FieldDiscordId = "discord_id"
	FieldIsWhitelisted = "is_whitelisted"
	FieldPassword = "password"
	FieldWallet = "wallet"
	FieldTokens = "tokens"
)
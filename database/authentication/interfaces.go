package authentication

type DatabaseEntry struct {
	Username string `json:"username" bson:"username"`
	UniqueUsername string `json:"unique_username" bson:"unique_username"`
	Password string `json:"password" bson:"password"`
	Wallet string   `json:"wallet" bson:"wallet"`
	Token []string  `json:"tokens" bson:"tokens"`
}

var (
	FieldUsername = "username"
	FieldUniqueUsername = "unique_username"
	FieldPassword = "password"
	FieldWallet = "wallet"
	FieldTokens = "tokens"
)
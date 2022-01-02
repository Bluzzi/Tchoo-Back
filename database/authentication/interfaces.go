package authentication

type DatabaseEntry struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Wallet string `json:"wallet" bson:"wallet"`
	Token []string `json:"tokens" bson:"tokens"`
}

var (
	FieldUsername = "username"
	FieldPassword = "password"
	FieldWallet = "wallet"
	FieldTokens = "tokens"
)
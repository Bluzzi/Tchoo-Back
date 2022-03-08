package errors

var (
	ErrorEmptyField = "field.empty"
	ErrorAccountAlreadyExists = "account.exists"
	ErrorAccountTokenInvalid = "account.token_invalid"
	ErrorAccountUsernameInvalid = "account.username_invalid"
	ErrorAccountNoWalletLinked = "account.no_wallet_linked"
	ErrorPasswordIncorrect = "account.password_incorrect"
	ErrorInvalidLogin = "account.invalid_login"
	ErrorWalletUsed = "wallet.used"
	ErrorWalletIncorrectSignature = "wallet.incorrect_signature"
	ErrorPetNotOwned = "pet.not_owned"
	ErrorPetActionOnCooldown = "pet.action_cooldown:"
	ErrorInvalidLocation = "location.invalid"
	ErrorInvalidFeedTime = "feed.time_invalid"
	ErrorInvalidSleepTime = "sleep.time_invalid"

	ErrorNotEnoughPoints = "lottery.not_enough_points"
)

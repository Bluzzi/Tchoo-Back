package errors

var (
	ErrorEmptyField = "field.empty"
	ErrorAccountAlreadyExists = "account.exists"
	ErrorAccountTokenInvalid = "account.token_invalid"
	ErrorInvalidLogin = "account.invalid_login"
	ErrorWalletUsed = "wallet.used"
	ErrorWalletIncorrectSignature = "wallet.incorrect_signature"
	ErrorPetNotOwned = "pet.not_owned"
	ErrorPetActionOnCooldown = "pet.action_cooldown:"
	ErrorPetActionNotEnoughSupply = "pet.action_ne_supply"
	ErrorInvalidLocation = "location.invalid"
	ErrorInvalidFeedTime = "feed.time_invalid"
	ErrorInvalidSleepTime = "feed.time_invalid"

	ErrorNotEnoughPoints = "lottery.not_enough_points"
)

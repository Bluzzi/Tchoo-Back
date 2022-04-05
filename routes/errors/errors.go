package errors

var (
	ErrorEmptyField = "field.empty"
	ErrorAccountAlreadyExists = "account.exists"
	ErrorAccountTokenInvalid = "account.token_invalid"
	ErrorAccountUsernameInvalid = "account.username_invalid"
	ErrorAccountNoWalletLinked = "account.no_wallet_linked"
	ErrorAccountNotExisting = "account.not_existing"
	ErrorPasswordIncorrect = "account.password_incorrect"
	ErrorInvalidLogin = "account.invalid_login"
	ErrorWalletUsed = "wallet.used"
	ErrorWalletIncorrectSignature = "wallet.incorrect_signature"
	ErrorPetNotOwned = "pet.not_owned"
	ErrorPetActionOnCooldown = "pet.action_cooldown:"
	ErrorInvalidLocation = "location.invalid"
	ErrorInvalidFeedTime = "feed.time_invalid"
	ErrorInvalidSleepTime = "sleep.time_invalid"

	ErrorInvalidNftNonce = "nft.nonce_invalid"
	ErrorTopOutOfRange = "top.out_of_range"

	ErrorNotEnoughPoints = "lottery.not_enough_points"
)

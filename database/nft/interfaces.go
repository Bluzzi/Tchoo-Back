package nft

type DatabaseEntry struct {
	// The unique nft nonce
	Nonce int64 `json:"nonce" bson:"nonce"`

	// The nft 3d model
	ThreeDModel string `json:"three_d_model" bson:"three_d_model"`
	
	// The nft's animations
	Animations Animations `json:"animations" bson:"animations"`

	// The nft 2d picture
	TwoDPicture string `json:"two_d_picture" bson:"two_d_picture"`

	// The nft name
	Name string `json:"name" bson:"name"`

	// Only used for requests
	HolderUsername string `json:"holder_username,omitempty" bson:"holder_username,omitempty"`

	// Won by the pet automatically + interacting can be used to buy lottery tickets
	PointsBalance float64 `json:"points_balance" bson:"points_balance"`

	// Won by interacting with the pet
	PrestigeBalance float64 `json:"prestige_balance" bson:"prestige_balance"`

	// A constant
	PointsPerHourBase float64 `json:"points_per_five_minutes_base" bson:"points_per_five_minutes_base"`

	// PointsPerHourBase affected by the state of the pet, will be the passive revenue
	PointsPerHourReal float64 `json:"points_per_five_minutes_real" bson:"points_per_five_minutes_real"`

	// Used to set the timeout and get when the actions can be used again: (actionName -> timestamp)
	ActionsUsed map[string]int64 `json:"actions_used" bson:"actions_used"`
}

type Animations struct {
	Idle string `json:"idle"`
	Purring string `json:"purring"`
}

var (
	FieldNonce = "nonce"
	FieldPointsBalance = "points_balance"
	FieldPrestigeBalance = "prestige_balance"
	FieldPointsPerHourBase = "points_per_five_minutes_base"
	FieldPointsPerHourReal = "points_per_five_minutes_real"
	FieldActionsUsed = "actions_used"
)
package nft

type DatabaseEntry struct {
	// The unique nft nonce
	Nonce int64 `json:"nonce,omitempty" bson:"nonce"`

	// The nft 3d model
	ThreeDModel string `json:"three_d_model,omitempty" bson:"three_d_model"`
	
	// The nft's animations
	Animations Animations `json:"animations,omitempty" bson:"animations"`

	// The nft 2d picture
	TwoDPicture string `json:"two_d_picture,omitempty" bson:"two_d_picture"`

	// The nft name
	Name string `json:"name,omitempty" bson:"name"`

	// Only used for requests
	HolderUsername string `json:"holder_username,omitempty" bson:"holder_username,omitempty"`

	// Won by the pet automatically + interacting can be used to buy lottery tickets
	PointsBalance float64 `json:"points_balance,omitempty" bson:"points_balance"`

	// Won by interacting with the pet
	PrestigeBalance float64 `json:"prestige_balance,omitempty" bson:"prestige_balance"`

	// A constant
	PointsPerHourBase float64 `json:"points_per_five_minutes_base,omitempty" bson:"points_per_five_minutes_base"`

	// PointsPerHourBase affected by the state of the pet, will be the passive revenue
	PointsPerHourReal float64 `json:"points_per_five_minutes_real,omitempty" bson:"points_per_five_minutes_real"`

	// Used to set the timeout and get when the actions can be used again: (actionName -> timestamp)
	ActionsUsed map[string]int64 `json:"actions_used,omitempty" bson:"actions_used"`
}

type Animations struct {
	Idle string `json:"idle,omitempty"`
	Purring string `json:"purring,omitempty"`
}

var (
	FieldNonce = "nonce"
	FieldName = "name"
	FieldPointsBalance = "points_balance"
	FieldPrestigeBalance = "prestige_balance"
	FieldTwoDPicture = "two_d_picture"
	FieldPointsPerHourBase = "points_per_five_minutes_base"
	FieldPointsPerHourReal = "points_per_five_minutes_real"
	FieldActionsUsed = "actions_used"
)
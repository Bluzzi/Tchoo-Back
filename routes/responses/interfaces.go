package responses

import (
	"MetaFriend/database/nft"
	"MetaFriend/lottery"
)

type SuccessResponse struct {
	Success bool `json:"success"`
	Error string `json:"error"`
}

type AccountCreateResponse struct {
	Success bool `json:"success"`
	IsWalletLinked bool `json:"is_wallet_linked,omitempty"`
	Token string `json:"token"`
}

type AccountInfosResponse struct {
	Username string `json:"username" bson:"username"`
	UniqueUsername string `json:"unique_username" bson:"unique_username"`
	Wallet string   `json:"wallet" bson:"wallet"`
	IsWhitelisted bool   `json:"is_whitelisted" bson:"is_whitelisted"`
	OwnedPets []int64   `json:"owned_pets" bson:"owned_pets"`
	DiscordId string   `json:"discord_id,omitempty" bson:"discord_id"`
}

type GetOwnedResponse struct {
	Success bool `json:"success"`
	OwnedNftsNonces []int64 `json:"owned_nfts_nonces"`
}

type GetTopResponse struct {
	Success bool `json:"success"`
	TotalCount int64 `json:"total_count"`
	TopNfts []nft.DatabaseEntry `json:"top_nfts"`
}

type AccountStatsResponse struct {
	Success bool `json:"success"`
	Username string `json:"username" bson:"username"`
	Wallet string   `json:"wallet" bson:"wallet"`
	OwnedNfts []nft.DatabaseEntry `json:"owned_nfts"`
}

type LotteryBuyTicketResponse struct {
	Success bool `json:"success"`
	WonPrize bool `json:"won_prize"`
	Prize string `json:"prize"`
	PrizePicture string `json:"prize_picture"`
}

type GetLotteryResponse struct {
	Success bool `json:"success"`
	Prizes []lottery.Prize `json:"prizes"`
	PriceATicket float64 `json:"price_a_ticket"`
}

type GetInvitesResponse struct {
	Success bool `json:"success"`
	Invites int64 `json:"invites"`
}

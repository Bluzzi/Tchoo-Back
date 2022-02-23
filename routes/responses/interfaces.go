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
	Token string `json:"token"`
}

type AccountInfosResponse struct {
	Username string `json:"username" bson:"username"`
	UniqueUsername string `json:"unique_username" bson:"unique_username"`
	Wallet string   `json:"wallet" bson:"wallet"`
}

type GetOwnedResponse struct {
	Success bool `json:"success"`
	OwnedNftsNonces []int64 `json:"owned_nfts_nonces"`
}

type GetTopResponse struct {
	Success bool `json:"success"`
	TopNfts []nft.DatabaseEntry `json:"top_nfts"`
}

type LotteryBuyTicketResponse struct {
	Success bool `json:"success"`
	WonPrize bool `json:"won_prize"`
	Prize string `json:"prize"`
	PrizePicture string `json:"prize_picture"`
}

type GetLotteryResponse struct {
	Prizes []lottery.Prize `json:"prizes"`
	PriceATicket float64 `json:"price_a_ticket"`
}

package responses

import "MetaFriend/database/nft"

type SuccessResponse struct {
	Success bool `json:"success"`
	Error string `json:"error"`
}

type AccountCreateResponse struct {
	Success bool `json:"success"`
	Token string `json:"token"`
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
	WonPrize bool `json:"won_prize"`
	Prize string `json:"prize"`
	PrizePicture string `json:"prize_picture"`
}


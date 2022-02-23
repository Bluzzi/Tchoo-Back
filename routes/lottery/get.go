package lottery

import (
	"MetaFriend/config"
	"MetaFriend/lottery"
	"MetaFriend/routes/responses"
	"encoding/json"
	"net/http"
)

func HandleGetRequest(w http.ResponseWriter, r *http.Request)  {
	_ = json.NewEncoder(w).Encode(responses.GetLotteryResponse{
		Prizes:       lottery.LotteryPrizes,
		PriceATicket: config.Config.PointsForATicket,
	})
}
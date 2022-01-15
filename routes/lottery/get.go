package lottery

import (
	"MetaFriend/lottery"
	"encoding/json"
	"net/http"
)

func HandleGetRequest(w http.ResponseWriter, r *http.Request)  {
	_ = json.NewEncoder(w).Encode(lottery.LotteryPrizes)
}
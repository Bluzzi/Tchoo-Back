package lottery

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

type Prize struct {
	IsNotAir bool `json:"-"`
	PrizeStr string `json:"prize"`
	Picture string `json:"picture"`
	Percent float64 `json:"percent"`
	Id float64 `json:"id"`
}

var LotteryPrizes []Prize
var lotteryArrayMutex sync.Mutex
var lotteryFileMutex sync.Mutex

func LoadLotteryPrizes()  {
	file, _ := ioutil.ReadFile("./data/lottery.json")
	_ = json.Unmarshal(file, &LotteryPrizes)

	fmt.Println("[LOAD] - Lottery Prizes loaded.")
}

func RandomlyGetPrize(accountName string) Prize {
	//////// POPULATE ARRAY
	sampleDataCount := 100000.0
	var arrayTotal []Prize

	// We have a percent (/100) so we need to have it at * 1000 (this is so we're able to make 0.01%)
	for _, prize := range LotteryPrizes {
		percentToAdd := prize.Percent * 1000
		sampleDataCount -= percentToAdd
		prize.IsNotAir = true

		for i := 0.0; i < percentToAdd; i++ {
			arrayTotal = append(arrayTotal, prize)
		}
	}

	for i := 0.0; i < sampleDataCount; i++ {
		arrayTotal = append(arrayTotal, Prize{
			IsNotAir:      false,
			PrizeStr: "air",
			Picture:  "",
			Percent:  0,
		})
	}

	rand.Seed(time.Now().UnixNano())
	p := arrayTotal[rand.Intn(100000)]
	if p.IsNotAir {
		fmt.Println("[LOTERY] - " + p.PrizeStr + " was won by " + accountName)
		PrizeWon(p, accountName)

		// Remove from the array
		lotteryArrayMutex.Lock()
		for i, prize := range LotteryPrizes {
			if prize.Id == p.Id {
				LotteryPrizes = append(LotteryPrizes[:i], LotteryPrizes[i+1:]...)
				break
			}
		}
		lotteryArrayMutex.Unlock()
	}
	return p
}

func PrizeWon(prize Prize, accountName string) {
	lotteryFileMutex.Lock()
	f, err := os.OpenFile("./data/won.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	if _, err := f.WriteString(prize.PrizeStr + " - " + accountName + "\n"); err != nil {
		log.Println(err)
	}
	lotteryFileMutex.Unlock()

	ActualizeLotteryJson()
}

func ActualizeLotteryJson() {
	lotteryFileMutex.Lock()
	bytes, _ := json.Marshal(LotteryPrizes)
	_ = os.WriteFile("./data/lottery.json", bytes, 0666)
	lotteryFileMutex.Unlock()
}
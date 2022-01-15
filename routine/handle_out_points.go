package routine

import (
	"MetaFriend/config"
	"MetaFriend/database/nft"
	"fmt"
	"time"
)

func HandleOutPointsRoutine() Routine {
	return Routine{
		Interval: time.Minute * 5,
		Function: func() {
			points := 0
			total := 0
			for _, entry := range nft.GetAllNfts() {
				// It means it's still in the smart contract, not bought
				if entry.HolderUsername == config.Config.ContractAddress {
					continue
				}

				points += int(entry.PointsPerHourReal)
				total++
				nft.EditStat(entry.Nonce, "$inc", nft.FieldPointsBalance, entry.PointsPerHourReal)
			}

			fmt.Println("[ROUTINE] (handle out points) Gave", points, "points in total to", total, "people")
		},
	}
}

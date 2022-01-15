package routine

import (
	"MetaFriend/database/nft"
	"MetaFriend/database/promises"
	"fmt"
	"time"
)

func PromisesResolverRoutine() Routine {
	return Routine{
		Interval: time.Minute,
		Function: func() {
			promiseResolved := 0
			for _, promise := range promises.GetExecutablePromises() {
				var incrementBy float64
				if promise.Type == promises.TypeDecrement {
					incrementBy = -promise.Value
				} else {
					incrementBy = promise.Value
				}

				nft.EditStat(
					promise.PetNonce,
					"$inc",
					nft.FieldPointsPerHourReal,
					incrementBy,
				)

				promises.ResolvePromise(promise.UniqueIdentifier)
				promiseResolved++
			}


			fmt.Println("[ROUTINE] (resolve promises) Resolved ", promiseResolved, " promises!")
		},
	}
}

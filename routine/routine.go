package routine

import (
	"MetaFriend/nft"
	"fmt"
	"time"
)

type Routine struct {
	Interval time.Duration
	Function func()
}

func (routine Routine) Start()  {
	go func() {
		for true {
			routine.Function()
			time.Sleep(routine.Interval)
		}
	}()
}

func StartRoutine() {
	fmt.Println("[LOAD] - Routine loaded.")

	Routine{
		Interval: time.Minute,
		Function: func() {
			nft.UpdateNonceAddressCache()
		},
	}.Start()

	PromisesResolverRoutine().Start()
	HandleOutPointsRoutine().Start()
	// Routine Cache
}

package play

import (
	"math/rand"
	"time"
)

func RandomCoordinate() Coordinate { ////////move this function somewhere else??
	x := int8(rand.Intn(19))
	y := int8(rand.Intn(19))
	random := Coordinate{y, x}
	return random
}

// artificialIdiot suggests a random move
func artificialIdiot(G *Game) { /////// move/remove?
	start := time.Now()
	suggestion := RandomCoordinate()
	// time.Sleep(498 * time.Millisecond) //////////
	elapsed := (time.Since(start))
	if G.Player == false {
		G.Ai0.Suggest = suggestion
		G.Ai0.Timer = elapsed
	} else {
		G.Ai1.Suggest = suggestion
		G.Ai1.Timer = elapsed
	}
}

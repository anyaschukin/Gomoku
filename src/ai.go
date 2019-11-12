package play

import (
	"math/rand"
	"time"
)

func randomcoordinate() coordinate { ////////move this function somewhere else??
	x := int8(rand.Intn(19))
	y := int8(rand.Intn(19))
	random := coordinate{y, x}
	return random
}

// artificialIdiot suggests a random move
func artificialIdiot(g *game) { /////// move/remove?
	start := time.Now()
	suggestion := randomcoordinate()
	// time.Sleep(498 * time.Millisecond) //////////
	elapsed := (time.Since(start))
	if g.player == false {
		g.ai0.suggest = suggestion
		g.ai0.timer = elapsed
	} else {
		g.ai1.suggest = suggestion
		g.ai1.timer = elapsed
	}
}

package play

import (
	"math/rand"
	"time"
)

func randomCoordinate() coordinate {
	x := int8(rand.Intn(19))
	y := int8(rand.Intn(19))
	random := coordinate{y, x}
	return random
}

// artificialIdiot suggests a random move
func artificialIdiot(g *game) {
	start := time.Now()
	suggestion := randomCoordinate()
	elapsed := (time.Since(start))
	if g.player == false {
		g.ai0.suggest = suggestion
		g.ai0.timer = elapsed
	} else {
		g.ai1.suggest = suggestion
		g.ai1.timer = elapsed
	}
}

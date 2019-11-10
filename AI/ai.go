package ai

import "time"

func Randomcoordinate() Coordinate { ////////move this function somewhere else??
	x := int8(rand.Intn(19))
	y := int8(rand.Intn(19))
	random := Coordinate{y, x}
	return random
}

func ai(G *game) {//(coordinate coordinate, elapsed time.Duration) {
	start := time.Now()
	//	"Improved" Minimax implementation (Alpha-beta pruning, negascout, mtdf, ...) -> 5 points!
	//	Move search Depth - 10 or more levels -> 5 points!
	//	search space of the algo - Multiple rectangular windows emcompassing placed stones but minimizing wasted space -> 5

	//	Heuristic
	//		take current alignments into account ?
	//		check whether an alignment has enough space to develop into a 5-in-a-row ?
	//		weigh an alignment according to its freedom (Free, half-free, flanked) ?
	//		take potential captures into account ?
	//		take current captured stones into account ?
	//		check for advanteageous combinations ?
	//		take both Players into account ?
	//		take past Player actions into account to identify patterns and weigh board states accordingly ?
	suggestion = Randomcoordinate()
	elapsed = time.Since(start)
	if G.Player == false {
		G.Ai0.Suggest = suggestion
		G.Ai0.Timer = elapsed
	} else {
		G.Ai1.Suggest = suggestion
		G.Ai1.Timer = elapsed
	}
}

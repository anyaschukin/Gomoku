package ai

// import "time"

// type state struct {
// 	Player     bool
// 	num_pieces uint8
// }

// func ai(g *game) (coordinate coordinate, elapsed time.Duration) {
// 	start := time.Now()
// 	//	"Improved" Minimax implementation (Alpha-beta pruning, negascout, mtdf, ...) -> 5 points!
// 	//	Move search depth - 10 or more levels -> 5 points!
// 	//	search space of the algo - Multiple rectangular windows emcompassing placed stones but minimizing wasted space -> 5

// 	//	Heuristic
// 	//		take current alignments into account ?
// 	//		check whether an alignment has enough space to develop into a 5-in-a-row ?
// 	//		weigh an alignment according to its freedom (Free, half-free, flanked) ?
// 	//		take potential captures into account ?
// 	//		take current captured stones into account ?
// 	//		check for advanteageous combinations ?
// 	//		take both Players into account ?
// 	//		take past Player actions into account to identify patterns and weigh board states accordingly ?
// 	coordinate = RandomCoordinate()
// 	elapsed = time.Since(start)
// 	return coordinate, elapsed
// }

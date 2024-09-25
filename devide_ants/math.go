package devide

import "math"

func Weight(paths [][]string, ants int) int {
	if len(paths) == 0 {
		return 0
	}
	var steps float64
	//reverse whight
	var rw float64
	// calculat ehte wieght of all graphs
	for _, path := range paths {
		// calculate the wight of a path
		wg := float64(len(path) - 1)
		rw += 1 / wg
	}
	steps = 1/rw + (float64(ants)-1)/float64(len(paths))
	rounded := int(math.Ceil(steps))
	if rounded < len(paths[len(paths)-1]) {
		return len(paths[len(paths)-1])
	}
	return rounded
}

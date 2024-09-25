package devide

import (
	"fmt"
	"lem-in/utils"
	"math"
)

// i need to know how much i'm going to use in every path
// so this calculation can be correct
func Weight(paths [][]string, ants int) int {
	if len(paths) == 0 {
		return 0
	}
	var steps, inverseWeight float64
	use := pick(paths, ants)
	// see how much the slowest path will take
	//var max int
	// Calculate the inverse weight of all paths

	//make the paths even

	for _, path := range use {
		// Calculate the weight of a path (path length - 1)
		pathWeight := float64(len(path) - 1)
		// no division by 0
		if pathWeight == 0 {
			continue
		}
		inverseWeight += 1 / pathWeight
	}

	// calculate the total steps required
	steps = 1/inverseWeight + (float64(ants)-1)/float64(len(use))

	roundedSteps := int(math.Ceil(steps))
	if len(use) < 1 {
		fmt.Println("debug")
		return 0
	}
	longestPath := len(use[len(use)-1])

	if roundedSteps < longestPath {
		return longestPath
	}
	return roundedSteps
}

// function to pick paths that we are going to use
func pick(paths [][]string, ants int) [][]string {
	// square is the ants requred to activate all paths
	var square = len(paths) * len(paths[len(paths)-1])
	for _, path := range paths {
		square -= len(path)
	}
	// suppose we are going to use them all
	use := utils.CopySliceSlice(paths)
	fmt.Println("use: ", use)
	// for every path that is not going to be activated , drop it
	for i := len(paths) - 1; i >= 0; i-- {
		// number of ants is not enaugh to activate the current branch
		if square >= ants {
			use = use[:i]
			square = len(use) * len(use[len(use)-1])
		} else {
			break
		}
	}
	return use
}

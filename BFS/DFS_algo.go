package bfs

import (
	"fmt"
	"os"
	"sort"
)

type Path struct {
	Score int
	Rooms []string
}

func (g *Graph) Isvalid() {
	for _, v := range g.Verteces {
		if len(v.adjacentVerteces) <= 2 {
			fmt.Println("bad Data")
			os.Exit(1)
		}
	}
}

// Main function to retrieve all paths from Start to End
func (g *Graph) AllPaths() {
	path := []string{g.Start.Name}
	g.GetPaths(path, g.End.Name, g.Start.Name)
	g.CalcScore()
	g.SortPaths()
	g.Paths = g.CleanPaths()
	for _, path := range g.Paths {
		for i, room := range path.Rooms {
			fmt.Printf("%s", room)
			if i < len(path.Rooms)-1 {
				fmt.Printf("->")
			}
		}
		fmt.Println()
	}
	g.Sort()
	g.Toall()
	// Print the cleaned paths
}

func (g *Graph) Toall() {
	for _, v := range g.Paths {
		g.All = append(g.All, v.Rooms)
	}
}

// Recursive function to find paths
func (g *Graph) GetPaths(path []string, target string, start string) {
	// If the current room is the target (end room), record the path
	if start == target {
		newPath := make([]string, len(path))
		copy(newPath, path)
		Reverse(newPath)
		g.Paths = append(g.Paths, &Path{Rooms: newPath})
		return
	}

	// Find the vertex corresponding to the current room
	vertex, ok := g.Verteces[start]
	if !ok {
		return
	}

	// Explore adjacent rooms
	for _, neighbor := range vertex.adjacentVerteces {
		if g.contains(path, neighbor.Name) {
			continue
		}
		newPath := append(path, neighbor.Name)
		g.GetPaths(newPath, target, neighbor.Name)
	}
}

// Helper function to check if a node exists in the current path (prevents cycles)
func (g *Graph) contains(path []string, node string) bool {
	for _, n := range path {
		if n == node {
			return true
		}
	}
	return false
}

// Sort paths based on their score (lower score is better)
func (g *Graph) SortPaths() {
	sort.Slice(g.Paths, func(i, j int) bool {
		return g.Paths[i].Score < g.Paths[j].Score
	})
}

func (g *Graph) Sort() {
	sort.Slice(g.Paths, func(i, j int) bool {
		return len(g.Paths[i].Rooms) < len(g.Paths[j].Rooms)
	})
}

// Clean paths to retain only paths with unique rooms (except start and end)
func (g *Graph) CleanPaths() []*Path {
	result := []*Path{}
	roomSet := make(map[string]bool) // Track rooms already used
	for _, path := range g.Paths {
		// fmt.Println(path)
		duplicate := false

		// Check if any room in the current path has already been used in other paths
		for _, room := range path.Rooms {
			if room != g.Start.Name && room != g.End.Name {
				if roomSet[room] {
					duplicate = true
					break
				}
			}
		}

		// If no duplicate was found, add the path and mark the rooms as used
		if !duplicate {
			result = append(result, path)
			for _, room := range path.Rooms {
				if room != g.Start.Name && room != g.End.Name {
					roomSet[room] = true
				}
			}
		}

	}
	return result
}

// Calculate the score of paths based on common rooms
func (g *Graph) CalcScore() {
	roomSet := make(map[string]int)

	// Count the occurrences of each room across all paths
	for _, path := range g.Paths {
		for _, room := range path.Rooms {
			if room != g.Start.Name && room != g.End.Name {
				roomSet[room]++
			}
		}
	}

	// Calculate the score for each path based on common rooms
	for _, path := range g.Paths {
		path.Score = 0

		// Increment score for rooms that appear in multiple paths
		for _, room := range path.Rooms {
			if room != g.Start.Name && room != g.End.Name && roomSet[room] > 1 {
				path.Score += roomSet[room]
			}
		}
	}
}

// func (g *Graph) Toprint() {
// 	for i := 0; i < g.Aints; i++ {

//		}
//	}
func (graph *Graph) PrintGraph() {
	for _, n := range graph.Verteces {
		fmt.Printf("Room %s: ", n.Name)
		for _, neighbor := range n.adjacentVerteces {
			fmt.Printf("%s <-> ", neighbor.Name)
		}
		fmt.Println("nil")
	}
	fmt.Printf("START: %s\n", graph.Start.Name)
	fmt.Printf("END: %s\n", graph.End.Name)
}

func Reverse(arr []string) {
	left := 0
	right := len(arr) - 1

	// Swap elements from both ends towards the center
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

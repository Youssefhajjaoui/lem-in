package main

import (
	"fmt"
	bfs "lem-in/BFS"
	devide "lem-in/devide_ants"
	fl "lem-in/parse_file"
)

func main() {
	// declare a new graph
	graph := bfs.NewGraph()
	// get data from file
	nest, err := fl.FillTheNest("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(nest)
	graph.Start = bfs.NewVertex(nest.Start)
	graph.End = bfs.NewVertex(nest.End)
	// creat vertexes

	// snap is a map name *vertex.
	snap := graph.NewVerteces(nest.Rooms)

	// creat edges relations betwen vertexes
	bfs.MakeEdge(nest.Tunels, snap)
	// creat start and end
	graph.PrintGraph()
	all := graph.FindAllWays()
	fmt.Println(all)
	mat := devide.Devide(all, 3)
	//////////////////////
	fmt.Println("/////////////////////////")
	devide.Print(mat)
}

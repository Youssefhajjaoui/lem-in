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

	// snap is a map name *vertex.
	snap := graph.NewVerteces(nest.Rooms)
	graph.Start = snap[nest.Start]
	graph.End = snap[nest.End]

	// creat edges relations betwen vertexes
	err = bfs.ConnectRooms(nest.Tunels, snap)

	if err != nil {
		fmt.Println(err)
		return
	}
	// creat start and end

	graph.Traverse()
	all := graph.FindAllWays()

	mat, err := devide.Devide(all, nest.Ants)
	if err != nil {
		fmt.Println(err)
		return
	}
	//////////////////////
	devide.Print(mat)
}

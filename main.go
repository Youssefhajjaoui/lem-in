package main

import (
	"fmt"
	"os"

	bfs "lem-in/BFS"
	devide "lem-in/devide_ants"
	fl "lem-in/parse_file"
)

func main() {
	// declare a new graph
	graph := bfs.NewGraph()
	// get data from file
	filename := "data.txt"
	if len(os.Args[1:]) != 0 {
		filename = os.Args[1]
	}
	nest, err := fl.FillTheNest(filename)
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

	fmt.Println("again start: ", graph.Start)
	fmt.Println("again end: ", graph.End)
	if err != nil {
		fmt.Println(err)
		return
	}
	// creat start and end
	graph.PrintGraph()
	graph.Traverse()
	all := graph.FindAllWays()
	fmt.Println(all)
	mat, err := devide.Devide(all, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	//////////////////////
	fmt.Println("/////////////////////////")
	devide.Print(mat)
}

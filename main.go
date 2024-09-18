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
	fmt.Println(graph.Start)
	fmt.Println(graph.End)
	// creat vertexes

	// snap is a map name *vertex.
	snap := graph.NewVerteces(nest.Rooms)
	fmt.Println("new rooms: ", snap)

	// creat edges relations betwen vertexes
	err = bfs.ConnectRooms(nest.Tunels, snap)
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

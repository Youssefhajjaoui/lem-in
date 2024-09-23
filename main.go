package main

import (
	"fmt"
	graphs "lem-in/graphs"
	fl "lem-in/parse_file"
	"os"
)

func main() {
	// declare a new graph
	file_name, err := fl.GetFileName(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}
	graph := graphs.NewGraph()
	// get data from file
	nest, err := fl.FillTheNest(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}

	// snap is a map name *vertex.
	// snap :=
	graph.NewVerteces(nest.Rooms)
	// graph.Start = snap[nest.Start]
	// graph.End = snap[nest.End]
	graph.Start = graph.Verteces[nest.Start]
	graph.End = graph.Verteces[nest.End]

	// creat edges relations betwen vertexes
	err = graph.ConnectRooms(nest.Tunels)
	if err != nil {
		fmt.Println(err)
		return
	}
	//////////////////////////////////////////////////////////////////
	// #############################################################

	all := graph.AllPaths(graph.Start, graph.End)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(all)

	maxFlow, _ := graph.EdmondsKarp()
	fmt.Println("max flow is: ", maxFlow)
	fmt.Println(all)
	fmt.Println("night light")
	fmt.Println(graph.AllPaths(graph.End, graph.Start))
}

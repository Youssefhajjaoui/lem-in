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

	all, err := graph.FindAllWays()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(all)

	maxFlow, _ := graph.EdmondsKarp()
	fmt.Println("max flow is: ", maxFlow)
	all, err = graph.BackFindAllWays()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(all)
}

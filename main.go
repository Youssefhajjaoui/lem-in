package main

import (
	"fmt"
	devide "lem-in/devide_ants"
	graphs "lem-in/graphs"
	fl "lem-in/parse_file"
	"os"
)

func main() {
	/*############ Parsing the File ###############*/
	// declare a new graph
	file_name, err := fl.GetFileName(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}
	// get data from file
	nest, err := fl.FillTheNest(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	/*############ Making the Graph ##############*/
	graph := graphs.NewGraph()
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
	/*############# Find the Best Paths ##############*/
	simple_paths := graph.AllPaths(graph.Start, graph.End, false)
	//fmt.Println("simple paths: ", simple_paths)
	simple, fsteps, err := devide.Devide(simple_paths, nest.Ants, graph.End.Name)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("steps of simple path: ", fsteps)

	/*maxFlow*/
	graph.EdmondsKarp()
	//fmt.Println("max flow is: ", maxFlow)
	edmonds := graph.AllPaths(graph.End, graph.Start, true)
	//fmt.Println("edmons paths: ", edmonds)
	carp, lsteps, err := devide.Devide(edmonds, nest.Ants, graph.End.Name)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("steps of edmonds are: ", lsteps)

	/*############# Devide The Ants ##################*/
	if fsteps <= lsteps {
		devide.Print(simple)
	} else {
		devide.Print(carp)
	}
	//////////////////////
}

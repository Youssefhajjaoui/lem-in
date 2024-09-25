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
	simple_paths := graph.AllPaths(graph.Start, graph.End)
	fmt.Println("simple paths: ", simple_paths)
	weight := devide.Weight(simple_paths, nest.Ants)
	fmt.Println("wieght of simple path: " , weight)

	maxFlow, _ := graph.EdmondsKarp()
	fmt.Println("max flow is: ", maxFlow)
	edmonds := graph.AllPaths(graph.End, graph.Start)
	fmt.Println("edmons paths: ",edmonds)
	weight = devide.Weight(edmonds, nest.Ants)
	fmt.Println("the steps taking are: ", weight)
	/*############# Devide The Ants ##################*/
	/*mat, _, err := devide.Devide(all, nest.Ants)
	if err != nil {
		fmt.Println(err)
		return
	}
	devide.Print(mat)*/
	//////////////////////
}

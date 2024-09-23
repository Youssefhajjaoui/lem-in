package main

import (
	"fmt"
	"os"

	graphs "lem-in/graphs"
	devide "lem-in/devide_ants"
	fl "lem-in/parse_file"
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

	/*oussama*/ 
	all, err  := graph.FindAllWays()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("oussama: roods found: ", all)
	mat, err := devide.Devide(all, nest.Ants)
	if err != nil {
		fmt.Println(err)
		return
	}
	devide.Print(mat)
	maxFlow, paths  := graph.EdmondsKarp()

	fmt.Printf("The maximum possible flow is %d\n", maxFlow)
	fmt.Println("paths taking")
	fmt.Println(paths)
	graph.AllPaths() 
	fmt.Println("try")
	all, err  = graph.FindAllWays()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("oussama: roods found: ", all)
	mat, err = devide.Devide(all, nest.Ants)
	if err != nil {
		fmt.Println(err)
		return
	}
	devide.Print(mat)
	//maxFlow, paths  = graph.EdmondsKarp()
	fmt.Println("paths")
	fmt.Println(paths)

	/*youcef*/ /* 
	graph.AllPaths()
	all := graph.All
	fmt.Println("youcef: roods found: ", all)
	mat, err := devide.Devide(all, nest.Ants)
	if err != nil {
		fmt.Println(err)
		return
	}
	devide.Print(mat)
	maxFlow, paths  := graph.EdmondsKarp()
	fmt.Printf("The maximum possible flow is %d\n", maxFlow)
	fmt.Println("paths taking")
	fmt.Println(paths)
	*/
}

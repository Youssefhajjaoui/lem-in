package main

import (
	bfs "lem-in/BFS"
	fl "lem-in/parse_file"
	devide "lem-in/devide_ants"
	"fmt"
)

func main() {
	// declare a new graph
	graph := bfs.NewGraph()
	// get data from file
	data, _ := fl.Getdata("data.txt")
	// creat vertexes
	Vertexes, _ := fl.GetNodes(data)
	// snap is a map name *vertex.
	snap := graph.CreatNodes(Vertexes)
	Edges, _ := fl.GetEdges(data)
	// creat edges relations betwen vertexes
	bfs.CreatEdge(Edges, snap)
	// creat start and end
	graph.SetStartEnd(fl.GetStart(data), fl.GetEnd(data))
	graph.PrintGraph()
	all := graph.FindAllWays()
	fmt.Println(all)
	mat := devide.Devide(all , 3)
	//////////////////////
	fmt.Println("/////////////////////////")
	devide.Print(mat)
}
 


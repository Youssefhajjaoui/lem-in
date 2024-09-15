package main

import (
	bfs "lem-in/BFS"
	fl "lem-in/parse_file"
	"fmt"
)

func main() {
	// declare a new graph
	graph := bfs.NewGraph()
	// get data from file
	data, _ := fl.Getdata("data.txt")
	// creat vertexes
	Vertexes, _ := fl.GetNodes(data)
	graph.CreatNodes(Vertexes)
	// creat edges relations betwen vertexes
	Edges, _ := fl.GetEdges(data)
	graph.CreatEdge(Edges)
	// creat start and end
	graph.SetStartEnd(fl.GetStart(data), fl.GetEnd(data))
	graph.PrintGraph()
	all := graph.FindAllWays()
	fmt.Println(all)
}


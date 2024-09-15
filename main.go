package main

import (
	bfs "lem-in/BFS"
	fl "lem-in/parse_file"
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
	// gr := bfs.NewGraph()
	// // this is all
	// gr.Start = graph.Start
	// found := graph.Search("4")
	// NotFound := graph.Search("1000")
	// fmt.Println("found the element with name v10: ", found)
	// fmt.Println("no elemen with name v100 found: ", NotFound)
	// ////////////////////////////////////////
	// fmt.Println("#########################")
	// fmt.Println("this is the shortest way")
	// // fmt.Println(graph.FindAllWays("v10"))
	// ss := graph.FirstSet("v10", map[string]bool{})
	// fmt.Println(ss)
	// // s := bfs.Domino(ss, "v10")
	// fmt.Println("#########################")
	// fmt.Println("these are all ways")
	// d := graph.FindAllWays("v10")
	// fmt.Println(d)
	// fmt.Println("#########################")
	// fmt.Println("#########################")
	// fmt.Println("#########################")
	// fmt.Println("#########################")
	// fmt.Println("#########################")
	// fmt.Println("#########################")
	// fmt.Println("#########################")
	// fmt.Println("#########################")
	// // for _, v := range graph.Verteces {
	// // 	fmt.Printf("all: %s <->", v.Name)
	// // 	for _, edje := range v.ad {
	// // 	}
	// // }
	// // fmt.Println("end:", graph.End.Name)
	// // fmt.Println("start:", graph.Start.Name)
	graph.PrintGraph()
	graph.Traverse()
}

// data , _ := fl.Getdata("data.txt")
// 	Nodes:=fl.Graphs
//
// 	gr.Start = bfs.NewVertex(data[0])
// 	for _, node := range nodes {
// 		v := bfs.NewVertex(node)
// 		gr.Add(v)
// 	}

package main 

import "fmt"
import bfs "lem-in/BFS"

func main(){
	graph := bfs.NewGraph()
	
	v1 := bfs.NewVertex("v1")
	v2 := bfs.NewVertex("v2")
	v2.Add_adjacent_vertex(v1)
	
	graph.Add(v2)
	graph.Add(v1)

	fmt.Println(graph.Verteces[0].Name)
	fmt.Println(graph.Verteces[1].Name)
}

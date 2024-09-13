package main 

import bfs "lem-in/BFS"

func main(){
	graph := bfs.NewGraph()
	
	v1 := bfs.NewVertex("v1")
	v2 := bfs.NewVertex("v2")
	v3 := bfs.NewVertex("v3")
	v4 := bfs.NewVertex("v4")

	v2.Add_adjacent_vertex(v1)
	v1.Add_adjacent_vertex(v3)
	v3.Add_adjacent_vertex(v4)

	graph.Start = v1

	graph.Add(v2)
	graph.Add(v1)
	
	graph.Traverse()
}

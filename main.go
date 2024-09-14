package main 

import bfs "lem-in/BFS"
import "fmt"

func main(){
	graph := bfs.NewGraph()
	
	v1 := bfs.NewVertex("v1")
	v2 := bfs.NewVertex("v2")
	v3 := bfs.NewVertex("v3")
	v4 := bfs.NewVertex("v4")
	v5 := bfs.NewVertex("v5")
	v6 := bfs.NewVertex("v6")
	v7 := bfs.NewVertex("v7")
	v8 := bfs.NewVertex("v8")
	v9 := bfs.NewVertex("v9")
	v10 := bfs.NewVertex("v10")
	v11 := bfs.NewVertex("v11")

	v1.Add_adjacent_vertex(v2)
	v1.Add_adjacent_vertex(v3)
	v1.Add_adjacent_vertex(v4)

	v2.Add_adjacent_vertex(v3)
	v2.Add_adjacent_vertex(v5)

	v3.Add_adjacent_vertex(v4)
	v3.Add_adjacent_vertex(v5)
	v3.Add_adjacent_vertex(v7)

	v4.Add_adjacent_vertex(v7)

	v5.Add_adjacent_vertex(v9)
	v5.Add_adjacent_vertex(v6)

	v7.Add_adjacent_vertex(v8)
	v7.Add_adjacent_vertex(v11)

	v6.Add_adjacent_vertex(v11)
	v6.Add_adjacent_vertex(v10)

	graph.Start = v1

	graph.Add(v1)
	
	graph.Traverse()
	
	found := graph.Search("v10")
	NotFound := graph.Search("v100")
	fmt.Println("found: ", found)
	fmt.Println("no found: ", NotFound)
}

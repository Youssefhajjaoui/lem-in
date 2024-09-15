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

	v20 := bfs.NewVertex("v20")
	v30 := bfs.NewVertex("v30")
	v40 := bfs.NewVertex("v40")
	v50 := bfs.NewVertex("v50")
	v60 := bfs.NewVertex("v60")
	v70 := bfs.NewVertex("v70")
	v80 := bfs.NewVertex("v80")

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
	v5.Add_adjacent_vertex(v10)
	v5.Add_adjacent_vertex(v6)

	v7.Add_adjacent_vertex(v8)
	v7.Add_adjacent_vertex(v11)

	v6.Add_adjacent_vertex(v11)
	v6.Add_adjacent_vertex(v10)
	v11.Add_adjacent_vertex(v10)
	
	v1.Add_adjacent_vertex(v20)
	v20.Add_adjacent_vertex(v30)
	v30.Add_adjacent_vertex(v8)
	v8.Add_adjacent_vertex(v40)
	v40.Add_adjacent_vertex(v50)
	v50.Add_adjacent_vertex(v60)
	v60.Add_adjacent_vertex(v70)
	v70.Add_adjacent_vertex(v80)
	v80.Add_adjacent_vertex(v10)

	graph.Start = v1

	graph.Add(v1)
	
	
	found := graph.Search("v10")
	NotFound := graph.Search("v100")
	fmt.Println("found: ", found)
	fmt.Println("no found: ", NotFound)
	////////////////////////////////////////
	fmt.Println("/////////////")
	fmt.Println("#########################")
	fmt.Println("these are all ways")
	//fmt.Println(graph.FindAllWays("v10"))
	ss := graph.FirstSet("v10", map[string]bool{})
	//s := bfs.Domino(ss, "v10")
	fmt.Println(ss)
	fmt.Println("#########################")
	d := graph.FindAllWays("v10")
	fmt.Println(d)
}

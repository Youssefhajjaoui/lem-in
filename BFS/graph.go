package bfs 

import "fmt"


type Graph struct{
	Verteces []*Vertex
	Start *Vertex // the starting room
	End *Vertex // the ending room
}

func NewGraph()*Graph{
	return &Graph{Verteces : []*Vertex{}, Start : nil , End : nil }
}


func (g *Graph)Add(v *Vertex){
	g.Verteces = append(g.Verteces, v)
}


func (g *Graph)Traverse(){
	for _, e := range g.Start.adjacentVerteces{
		fmt.Println(e.Name)
	}
}

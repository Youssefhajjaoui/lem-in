package bfs 

type Graph struct{
	Verteces []*Vertex
	Start *Vertex // the starting room
	End *Vertex // the ending room
}

func NewGraph()*Graph{
	return &Graph{Verteces : []*Vertex{}, Start : nil , End : nil }
}


func (g *Graph)AddRoom(v *Vertex){
	g.Verteces = append(g.Verteces, v)
}











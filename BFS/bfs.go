package bfs 


type Vertex struct{
	// name of the room
	Name string
	// cordonate of the room
	//crd [2]int
	// liks from the current room to other rooms
	adjacentVerteces []*Vertex 
}
// still don't know how to make this works 

func NewVertex(name string)*Vertex{
    return &Vertex{Name : name}
}

// this is to relate verticies
func (v *Vertex)Add_adjacent_vertex(vertex *Vertex){
	if v.include(vertex){
		return
	}
	v.adjacentVerteces = append(v.adjacentVerteces, vertex)
	vertex.Add_adjacent_vertex(v)
}
// check if if the v is already related to vertex
func (v *Vertex)include(vertex *Vertex)bool{
    for _, e := range v.adjacentVerteces {
        if e == vertex {
            return true
        }
    }
    return false
}
//////////////////////////////////////////////////////////////
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











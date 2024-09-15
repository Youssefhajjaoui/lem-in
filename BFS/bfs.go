package bfs

import (
	"errors"
)

type Vertex struct {
	// name of the room
	Name string
	// cordonate of the room
	// crd [2]int
	// liks from the current room to other rooms
	adjacentVerteces []*Vertex
}

// still don't know how to make this works
func NewVertex(name string) *Vertex {
	return &Vertex{Name: name}
}

// this is to relate verticies
func (g *Graph) Add_adjacent_vertex(node1 string, node2 string) error {
	vertex1, err := g.GetnodbyValue(node1)
	vertex2, err := g.GetnodbyValue(node2)
	if err != nil {
		return errors.New("add adjacent for no existing vertex !! ")
	}
	if !vertex1.include(vertex2) {
		vertex1.adjacentVerteces = append(vertex1.adjacentVerteces, vertex2)
	}
	if !vertex2.include(vertex1) {
		vertex2.adjacentVerteces = append(vertex2.adjacentVerteces, vertex1)
	}
	return nil
}

// check if if the v is already related to vertex
func (v *Vertex) include(vertex *Vertex) bool {
	for _, e := range v.adjacentVerteces {
		if e == vertex {
			return true
		}
	}
	return false
}

func (g *Graph) SetStartEnd(start string, end string) error {
	Start, err := g.GetnodbyValue(start)
	End, err := g.GetnodbyValue(end)
	if Start == nil || End == nil || err != nil {
		return errors.New("start or End is not valid")
	}
	g.Start = Start
	g.End = End
	return nil
}

func (g *Graph) GetnodbyValue(value string) (*Vertex, error) {
	for _, node := range g.Verteces {
		if node.Name == value {
			return node, nil
		}
	}

	return nil, errors.New("no vertex with this name")
}

func (g *Graph) CreatNodes(Nodes []string) {
	for _, Node := range Nodes {
		vertex := NewVertex(Node)
		g.Verteces = append(g.Verteces, vertex)
	}
}

func (g *Graph) CreatEdge(edges [][]string) {
	for _, cols := range edges {
		g.Add_adjacent_vertex(cols[0], cols[1])
	}
}

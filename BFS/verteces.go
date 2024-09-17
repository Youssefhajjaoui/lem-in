package bfs

import (
	"errors"
)

type Vertex struct {
	Name             string
	adjacentVerteces []*Vertex
}

func NewVertex(name string) *Vertex {
	return &Vertex{
		Name:             name,
		adjacentVerteces: []*Vertex{},
	}
}

// ///////////////////////////////////////////////////////////////////////////////
// this is to relate verticies
func (v *Vertex) Add_adjacent_vertex(vertex *Vertex) error {
	does, err := v.include(vertex)
	if err != nil {
		return err
	}
	if !does {
		v.adjacentVerteces = append(v.adjacentVerteces, vertex)
	}
	vertex.adjacentVerteces = append(vertex.adjacentVerteces, v)
	return nil
}

/////////////////////////////////////////////////////////////////////////////////

// check if if the v is already related to vertex
func (v *Vertex) include(vertex *Vertex) (bool, error) {
	var err error
	if v == nil || vertex == nil {
		if v == nil {
			err = errors.New("source pointer is nil")
		}
		if vertex == nil {
			err = errors.New("distination pointer is nil")
		}
		return false, err
	}
	for _, e := range v.adjacentVerteces {
		if e == vertex {
			return true, nil
		}
	}
	return false, nil
}

func (g *Graph) SetStartEnd(start string, end string) error {
	Start, err := g.GetnodbyValue(start)
	if err != nil {
		return err
	}
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

// the name of a vertex can not be repeated
func (g *Graph) CreatNodes(Nodes []string) map[string]*Vertex {
	snap := make(map[string]*Vertex)
	for _, Node := range Nodes {
		vertex := NewVertex(Node)
		snap[Node] = vertex
		g.Verteces = append(g.Verteces, vertex)
	}
	return snap
}

func CreatEdge(edges [][]string, snap map[string]*Vertex) {
	for _, cols := range edges {
		// name name
		snap[cols[0]].Add_adjacent_vertex(snap[cols[1]])
	}
}

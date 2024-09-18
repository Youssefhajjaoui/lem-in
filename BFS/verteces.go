package bfs

import (
	"errors"
	"fmt"
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

// Check if the vertex is already related to the given vertex.
func (v *Vertex) include(vertex *Vertex) (bool, error) {
	if v == nil && vertex == nil {
		return false, errors.New("both source and destination pointers are nil")
	}
	if v == nil {
		return false, errors.New("source pointer is nil")
	}
	if vertex == nil {
		return false, errors.New("destination pointer is nil")
	}

	// Check if the vertex is in the adjacent vertices list
	for _, e := range v.adjacentVerteces {
		if e == vertex {
			return true, nil
		}
	}
	return false, nil
}

// the name of a vertex can not be repeated
func (g *Graph) NewVerteces(names []string) map[string]*Vertex {
	var snap = make(map[string]*Vertex)
	for _, name := range names {
		vertex := NewVertex(name)
		snap[name] = vertex
		g.Verteces = append(g.Verteces, vertex)
	}
	return snap
}

func ConnectRooms(edges [][2]string, snap map[string]*Vertex) error {
	for _, cols := range edges {
		// Ensure both vertices exist in the snap map
		if vertex1, ok1 := snap[cols[0]]; ok1 {
			if vertex2, ok2 := snap[cols[1]]; ok2 {
				if err := vertex1.Add_adjacent_vertex(vertex2); err != nil {
					return fmt.Errorf("%w: %s -> %s", err, cols[0], cols[1])
				}
			} else {
				return fmt.Errorf("vertex %s not found in snap", cols[1])
			}
		} else {
			return fmt.Errorf("vertex %s not found in snap", cols[0])
		}
	}
	return nil
}

func (graph *Graph) PrintGraph() {
	for _, n := range graph.Verteces {
		fmt.Printf("Room %s: ", n.Name)
		for _, neighbor := range n.adjacentVerteces {
			fmt.Printf("%s <-> ", neighbor.Name)
		}
		fmt.Println("nil")
	}
	fmt.Printf("START: %s\n", graph.Start.Name)
	fmt.Printf("END: %s\n", graph.End.Name)
}

package bfs

import (
	"fmt"
	"strings"
)

type Vertex struct {
	// name of the room
	Name string
	// cordonate of the room
	// crd [2]int
	// liks from the current room to other rooms
	adjacentVerteces []*Vertex
}

type Graphs struct {
	Verteces []*Vertex
	Start    *Vertex
	End      *Vertex
}

func (graph *Graphs) addNode(value string) *Vertex {
	node := &Vertex{
		Name:             value,
		adjacentVerteces: []*Vertex{},
	}
	graph.Verteces = append(graph.Verteces, node)
	return node
}

func (graph *Graphs) addEdge(U, V *Vertex) {
	U.adjacentVerteces = append(U.adjacentVerteces, V)
	V.adjacentVerteces = append(V.adjacentVerteces, U)
}

func (graph *Graphs) PrintGraph() {
	for _, n := range graph.Verteces {
		fmt.Printf("Room %d: ", n.Name)
		for _, neighbor := range n.adjacentVerteces {
			fmt.Printf("%d <-> ", neighbor.Name)
		}
		fmt.Println("nil")
	}
}

func (g *Graphs) ParsetoNode(values []string) []*Vertex {
	Nodes := []*Vertex{}
	for i := 0; i < len(values); i++ {
		Node := g.addNode(values[i])
		Nodes = append(Nodes, Node)
	}
	return Nodes
}

func (g *Graphs) GetEdges(arr []string, Nodes []*Vertex) error {
	for _, v := range arr {
		if Len := len(strings.Split(v, "-")); Len == 2 {
			Node1 := g.GetnodbyValue(strings.Split(v, "-")[0])
			Node2 := g.GetnodbyValue(strings.Split(v, "-")[1])
			g.addEdge(Node1, Node2)
		}
	}
	return nil
}

func (g *Graphs) GetnodbyValue(value string) *Vertex {
	for _, node := range g.Verteces {
		if node.Name == value {
			return node
		}
	}
	return nil
}

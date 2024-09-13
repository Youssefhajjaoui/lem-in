package bfs

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Edges []*Node
}

type Graphs struct {
	Nodes []*Node
}

func (graph *Graphs) addNode(value int) *Node {
	node := &Node{
		Value: value,
	}
	graph.Nodes = append(graph.Nodes, node)
	return node
}

func (graph *Graphs) addEdge(U, V *Node) {
	U.Edges = append(U.Edges, V)
	V.Edges = append(V.Edges, U)
}

func (graph *Graphs) printGraph() {
	for _, n := range graph.Nodes {
		fmt.Printf("Room %d: ", n.Value)
		for _, neighbor := range n.Edges {
			fmt.Printf("%d <-> ", neighbor.Value)
		}
		fmt.Println("nil")
	}
}

func (g *Graphs) ParsetoNode(values []int) []*Node {
	Nodes := []*Node{}
	for i := 0; i < len(values); i++ {
		Node := g.addNode(values[i])
		Nodes = append(Nodes, Node)
	}
	return Nodes
}

func (g *Graphs) GetEdges(arr []string, Nodes []*Node) error {
	for _, v := range arr {
		if Len := len(strings.Split(v, " ")); Len == 2 {
			node1, err := strconv.Atoi(strings.Split(v, "")[0])
			if err != nil {
				return errors.New("can't convert it")
			}
			node2, err := strconv.Atoi(strings.Split(v, "")[1])
			if err != nil {
				return errors.New("can't convert it")
			}
			Node1 := g.GetnodbyValue(node1)
			Node2 := g.GetnodbyValue(node2)
			g.addEdge(Node1, Node2)
		}
	}
	return nil
}

func (g *Graphs) GetnodbyValue(value int) *Node {
	for _, node := range g.Nodes {
		if node.Value == value {
			return node
		}
	}
	return nil
}

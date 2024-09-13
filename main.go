package main

import "fmt"

type node struct {
	value int
	edges []*node
}

type graph struct {
	nodes []*node
}

func (graph *graph) addNode(value int) *node {
	node := &node{
		value: value,
	}
	graph.nodes = append(graph.nodes, node)
	return node
}

func (graph *graph) addEdge(U, V *node) {
	U.edges = append(U.edges, V)
	V.edges = append(V.edges, U)
}

func (graph *graph) printGraph() {
	for _, n := range graph.nodes {
		fmt.Printf("Room %d: ", n.value)
		for _, neighbor := range n.edges {
			fmt.Printf("%d <-> ", neighbor.value)
		}
		fmt.Println("nil")
	}
}

func main() {
	// Create a new graph
	g := &graph{}

	// Add nodes (rooms)
	node0 := g.addNode(0)
	node1 := g.addNode(1)
	node2 := g.addNode(2)
	node3 := g.addNode(3)
	node4 := g.addNode(4)
	node5 := g.addNode(5)
	node6 := g.addNode(6)
	node7 := g.addNode(7)

	// Add edges (tunnels between rooms)
	g.addEdge(node0, node4)
	g.addEdge(node0, node6)
	g.addEdge(node1, node3)
	g.addEdge(node4, node3)
	g.addEdge(node5, node2)
	g.addEdge(node3, node5)
	g.addEdge(node4, node2)
	g.addEdge(node2, node1)
	g.addEdge(node7, node6)
	g.addEdge(node7, node2)
	g.addEdge(node7, node4)
	g.addEdge(node6, node5)

	// Print the adjacency list of the graph
	g.printGraph()
}

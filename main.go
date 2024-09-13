package main

import (
	"os"

	bfs "lem-in/BFS"
	"lem-in/parse_file"
)

func main() {
	// read files and get data

	// Create a new graph
	g := &bfs.Graphs{}
	parse_file.ProcessInput(os.Args[1], g)
	// Add nodes (rooms)
	// node0 := g.addNode(0)
	// node1 := g.addNode(1)
	// node2 := g.addNode(2)
	// node3 := g.addNode(3)
	// node4 := g.addNode(4)
	// node5 := g.addNode(5)
	// node6 := g.addNode(6)
	// node7 := g.addNode(7)

	// // Add edges (tunnels between rooms)
	// g.addEdge(node0, node4)
	// g.addEdge(node0, node6)
	// g.addEdge(node1, node3)
	// g.addEdge(node4, node3)
	// g.addEdge(node5, node2)
	// g.addEdge(node3, node5)
	// g.addEdge(node4, node2)
	// g.addEdge(node2, node1)
	// g.addEdge(node7, node6)
	// g.addEdge(node7, node2)
	// g.addEdge(node7, node4)
	// g.addEdge(node6, node5)

	// // Print the adjacency list of the graph
}

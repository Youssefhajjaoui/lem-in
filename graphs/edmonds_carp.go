package graphs

import (
	"lem-in/queue"
)

// BFS to find an augmenting path
func (g *Graph) bfs(parent map[*Vertex]*Vertex) bool {
	visited := make(map[*Vertex]bool)
	q := queue.New() // Using a simple slice as a queue
	q.Enqueue(g.Start)
	visited[g.Start] = true

	for !q.IsEmpty() {
		current := q.Dequeue().Item.(*Vertex)
		for _, neighbor := range current.adjacentVerteces {
			if !visited[neighbor] { // Not visited
				q.Enqueue(neighbor) // Enqueue
				visited[neighbor] = true
				parent[neighbor] = current

				if neighbor == g.End {
					return true
				}
			}
		}
	}

	return false
}

// Helper function to check if a vertex is in the slice
func contains(slice []*Vertex, vertex *Vertex) bool {
	for _, v := range slice {
		if v == vertex {
			return true
		}
	}
	return false
}

// Edmonds-Karp algorithm to find maximum flow
func (g *Graph) EdmondsKarp() (int, [][]string) {
	totalFlow := 0
	parent := make(map[*Vertex]*Vertex)
	paths := [][]string{}

	for g.bfs(parent) {
		// Found an augmenting path
		totalFlow++
		path := []string{}

		// Construct the path from the end to the start
		for v := g.End; v != nil; v = parent[v] {
			path = append([]string{v.Name}, path...) // Prepend to maintain order
			if v == g.Start {
				break
			}
		}

		// Update the edges in the path
		for v := g.End; v != g.Start; v = parent[v] {
			u := parent[v]

			// Remove the forward edge
			for i, adjacent := range u.adjacentVerteces {
				if adjacent == v {
					u.adjacentVerteces = append(u.adjacentVerteces[:i], u.adjacentVerteces[i+1:]...)
					break
				}
			}

			// Add the backward edge (if necessary)
			if !contains(u.adjacentVerteces, parent[v]) {
				u.adjacentVerteces = append(u.adjacentVerteces, parent[v])
			}
		}

		paths = append(paths, path)
	}

	return totalFlow, paths
}
func (g *Graph) NightLight() [][]string {
	visited := make(map[*Vertex]bool)
	paths := [][]string{}
	for path := g.BackwardBFS(copyMapV(visited)); len(path) != 0; {
		paths = append(paths, path)
		for _, v := range path {
			visited[g.Verteces[v]] = true
		}
	}
	return paths
}
func copyMapV(original map[*Vertex]bool) map[*Vertex]bool {
	newMap := make(map[*Vertex]bool)
	for key, value := range original {
		newMap[key] = value
	}
	return newMap
}

// BFS function to find the shortest path from End to Start
func (g *Graph) BackwardBFS(visited map[*Vertex]bool) []string {
	// Queue to store the nodes for BFS
	q := queue.New()
	q.Enqueue(g.End)

	// Map to store the parent of each node in the BFS tree
	parent := make(map[*Vertex]*Vertex)
	//visited := make(map[*Vertex]bool)

	visited[g.End] = true

	// BFS loop
	for !q.IsEmpty() {
		// Get the current node
		node := q.Dequeue().Item.(*Vertex)

		// If we reach the Start node, we can construct the path
		if node == g.Start {
			return constructPath(parent, g.Start, g.End)
		}

		// Explore all adjacent nodes
		for _, neighbor := range g.End.adjacentVerteces {
			if !visited[neighbor] {
				visited[neighbor] = true
				parent[neighbor] = node
				q.Enqueue(neighbor)
			}
		}
	}

	// Return an empty slice if no path is found
	return []string{}
}

// Helper function to reconstruct the path from parent map
func constructPath(parent map[*Vertex]*Vertex, start, end *Vertex) []string {
	var path []string
	for node := start; node != end; node = parent[node] {
		path = append([]string{node.Name}, path...) // Prepend the node
	}
	path = append([]string{end.Name}, path...) // Add the end node at the start
	return path
}

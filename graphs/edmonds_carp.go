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

func copyMapV(original map[*Vertex]bool) map[*Vertex]bool {
	newMap := make(map[*Vertex]bool)
	for key, value := range original {
		newMap[key] = value
	}
	return newMap
}

// Helper function to reconstruct the path from parent map
func constructPath(parent map[*Vertex]*Vertex, start, end *Vertex) []string {
	var path []string
	for node := end; node != nil; node = parent[node] { // Fix here: start from end
		path = append([]string{node.Name}, path...) // Prepend the node
	}
	return path
}

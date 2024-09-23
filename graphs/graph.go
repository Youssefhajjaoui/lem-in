package graphs

import (
	"fmt"
	"lem-in/queue"
	Q "lem-in/queue"
	"lem-in/utils"
)

type Graph struct {
	Verteces map[string]*Vertex
	Start    *Vertex // the starting room
	End      *Vertex // the ending room
}

func NewGraph() *Graph {
	return &Graph{Verteces: make(map[string]*Vertex), Start: nil, End: nil}
}

// add a vertex to the graph
func (g *Graph) Add(v *Vertex) {
	g.Verteces[v.Name] = v
}

// this is used to track the path i used to the end
func assemble(parts [][2]string, exit string) []string {
	var find string
	path := []string{}
	done := false
	for i := len(parts) - 1; i >= 0; i-- {
		if parts[i][1] == exit && !done {
			path = append(path, exit)
			path = append(path, parts[i][0])
			done = true
			find = parts[i][0]
		}
		if done {
			if parts[i][1] == find {
				path = append(path, parts[i][0])
				find = parts[i][0]
			}
		}

	}
	return path
}

// BFS to find an augmenting path between from and to
func (g *Graph) BFS(from, to *Vertex, visited map[string]bool) []string {
	//parent := make(map[*Vertex]*Vertex)
	parent := [][2]string{}
	q := queue.New() // Using a simple slice as a queue
	q.Enqueue(from)
	visited[from.Name] = true
	// If we reach the Start node, we can construct the path

	for !q.IsEmpty() {
		current := q.Dequeue().Item.(*Vertex)
		if current == to {
			return assemble(parent, g.Start.Name)
		}
		for _, neighbor := range current.adjacentVerteces {
			if !visited[neighbor.Name] { // Not visited
				q.Enqueue(neighbor) // Enqueue
				visited[neighbor.Name] = true
				//parent[neighbor] = current
				parent = append(parent, [2]string{current.Name, neighbor.Name})
			}
		}
	}
	return nil
}

func (g *Graph) AllPaths(from, to *Vertex) [][]string {
	visited := make(map[string]bool)
	paths := [][]string{}
	for {
		path := g.BFS(from, to, utils.CopyMap(visited))
		if len(path) < 1 {
			break
		}
		paths = append(paths, path)
		for _, v := range path {
			if v != g.Start.Name {
				visited[v] = true
			}
		}
	}
	return paths
}

/*#####################################################################*/
// function of debuging
/*#####################################################################*/

// this is a function that traverse the graph from -> from tell there
// is nothing more to traverse
func (g *Graph) Traverse(from *Vertex) {
	fmt.Println("start traversing the graph")
	fmt.Println("this is the start: ", from.Name)

	// Initialize the queue
	q := Q.New()
	q.Enqueue(from)
	// Create a map to track visited vertices
	visited := make(map[*Vertex]bool)
	visited[from] = true
	// Start traversing the graph
	for !q.IsEmpty() {

		dequeuedItem := q.Dequeue()
		e, ok := dequeuedItem.Item.(*Vertex)
		if !ok {

			continue
		}
		// Process all adjacent vertices

		for _, adjVertex := range e.adjacentVerteces {

			if visited[adjVertex] {
				continue
			}

			visited[adjVertex] = true
			q.Enqueue(adjVertex)
			fmt.Print(adjVertex.Name + "->")
		}
	}
	fmt.Println("end traversing the graph")
}

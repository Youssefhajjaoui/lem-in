package graphs

import (
	"errors"
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

// BFS to find an augmenting path between from and to
func (g *Graph) BFS(from, to *Vertex, visited map[string]bool) []string {

	//parent := make(map[*Vertex]*Vertex)
	//parent := [][2]string{}
	parent := make(map[*Vertex]*Vertex)
	q := queue.New() // Using a simple slice as a queue
	q.Enqueue(from)
	visited[from.Name] = true
	// If we reach the Start node, we can construct the path

	for !q.IsEmpty() {

		current := q.Dequeue().Item.(*Vertex)
		if current == to {
			//return assemble(parent, to.Name)
			path := constructPath(parent, to, from)
			// case of the start connected to the end
			if len(path) == 2 {
				// break the connection forward
				err := g.breakEndStart()
				fmt.Println(err)
				return nil
			}
			return path

		}

		for _, neighbor := range current.adjacentVerteces {
			if !visited[neighbor.Name] { // Not visited
				q.Enqueue(neighbor) // Enqueue
				visited[neighbor.Name] = true
				//parent[current] = neighbor
				parent[neighbor] = current

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
		if len(path) != 2 {
			for _, v := range path {
				if v != to.Name {
					visited[v] = true
				}
			}
		}
	}
	return paths
}

// Helper function to reconstruct the path from parent map
func constructPath(parent map[*Vertex]*Vertex, from, to *Vertex) []string {
	var path []string
	for v := from; v != nil; v = parent[v] { // Fix here: start from end
		path = append([]string{v.Name}, path...) // Prepend the node
	}
	return path
}

func (g *Graph) breakEndStart() error {
	for i := 0; i < len(g.End.adjacentVerteces); i++ {
		if g.End.adjacentVerteces[i] == g.Start {
			g.End.adjacentVerteces = append(g.End.adjacentVerteces[:i], g.End.adjacentVerteces[i+1:]...)
			return nil
		}
	}
	return errors.New("not start end connection found")
}

/*##############################################################*/
// function of debuging
/*##############################################################*/

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

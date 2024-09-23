package graphs

import (
	"errors"
	"fmt"
	Q "lem-in/queue"
)

/*
	type Graph struct {
		Verteces []*Vertex
		Start    *Vertex // the starting room
		End      *Vertex // the ending room
	}
*/
type Graph struct {
	//Verteces []*Vertex
	Verteces map[string]*Vertex
	Start    *Vertex // the starting room
	End      *Vertex // the ending room
	Paths    []*Path
	Aints    int
	Target   int
	All      [][]string
}

func NewGraph() *Graph {
	return &Graph{Verteces: make(map[string]*Vertex), Start: nil, End: nil}
}

func (g *Graph) Add(v *Vertex) {
	g.Verteces[v.Name] = v
}

// Traverse traverses all the vertices in the graph.
func (g *Graph) Traverse() {
	fmt.Println("start traversing the graph")
	fmt.Println("this is the start: ", g.Start.Name)

	// Initialize the queue
	q := Q.New()
	q.Enqueue(g.Start)

	// Create a map to track visited vertices
	visited := make(map[*Vertex]bool)
	visited[g.Start] = true
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

func (g *Graph) FirstSet(name string, pas map[string]bool) ([]string, error) {

	// make a que
	visited := copyMap(pas)
	q := Q.New()
	q.Enqueue(g.Start)
	// make a map

	//visited := make(map[*Vertex]bool)

	visited[g.Start.Name] = true
	var from [][2]string
	if g.Start.Name == g.End.Name {
		return nil, errors.New("the starting and ending rooms are the same")
	}

	for !q.IsEmpty() {
		found := false
		save := []*Vertex{}

		dequeuedItem := q.Dequeue()
		e, ok := dequeuedItem.Item.(*Vertex)
		if !ok {
			continue
		}
		///////////////////////////////////////
		for i, l := range e.adjacentVerteces {
			if visited[l.Name] {
				continue
			}
			if l.Name != name {
				visited[l.Name] = true
			} else {
				found = true
				// break the bind
				if e == g.Start {
					e.adjacentVerteces = append(e.adjacentVerteces[:i], e.adjacentVerteces[i+1:]...)

				}
			}
			save = append(save, l)
			from = append(from, [2]string{e.Name, l.Name})
			if found {
				return assemble(from, name), nil
			}
		}
		if !found {
			for _, el := range save {
				q.Enqueue(el)
			}
		}
	}
	return assemble(from, name), nil
}

// this is used so i don't pass by refrence
func copyMap(original map[string]bool) map[string]bool {
	newMap := make(map[string]bool)
	for key, value := range original {
		newMap[key] = value
	}
	return newMap
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

func (g *Graph) FindAllWays() ([][]string, error) {
	// find the first set
	name := g.End.Name
	var paths [][]string
	block := make(map[string]bool)
	var stop = true
	for stop {
		ss, err := g.FirstSet(name, block)
		if err != nil {
			return nil, err
		}
		if len(ss) == 0 {
			stop = false
			continue
		}
		paths = append(paths, ss)
		for _, s := range ss {
			if s != name {
				block[s] = true
			}
		}
	}
	return paths, nil
}

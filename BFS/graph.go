package bfs

import (
	"fmt"

	Q "lem-in/queue"
)

type Graph struct {
	Verteces []*Vertex
	Start    *Vertex // the starting room
	End      *Vertex // the ending room
}

func NewGraph() *Graph {
	return &Graph{Verteces: []*Vertex{}, Start: nil, End: nil}
}

func (g *Graph) Add(v *Vertex) {
	g.Verteces = append(g.Verteces, v)
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

// this traverse all the graph
func (g *Graph) Traverse() {
	fmt.Println("start traversing")
	// make a que
	q := Q.New()
	q.Enqueue(g.Start)
	fmt.Println(g.Start.Name)
	// make a map
	visited := make(map[*Vertex]bool)
	visited[g.Start] = true
	// start from the q and gethem all
	// e is a node
	for !q.IsEmpty() {
		dequeuedItem := q.Dequeue()
		e, ok := dequeuedItem.Item.(*Vertex)
		if !ok {
			continue
		}
		for _, l := range e.adjacentVerteces {
			if visited[l] {
				continue
			}
			visited[l] = true
			q.Enqueue(l)
			fmt.Print(l.Name + "->")
		}
	}
	fmt.Println("end traversing the graph")
}

// this is searching an intem and returning pointer to the vertex
func (g *Graph) Search(name string) *Vertex {
	// make a que
	q := Q.New()
	q.Enqueue(g.Start)
	// make a map
	visited := make(map[*Vertex]bool)
	visited[g.Start] = true
	if g.Start.Name == name {
		return g.Start
	}
	// start from the q and gethem all
	// e is a node
	for !q.IsEmpty() {
		dequeuedItem := q.Dequeue()
		e, ok := dequeuedItem.Item.(*Vertex)
		if !ok {
			continue
		}
		for _, l := range e.adjacentVerteces {
			if l.Name == name {
				return l
			}
			if visited[l] {
				continue
			}
			visited[l] = true
			q.Enqueue(l)
		}
	}
	return nil
}

// start from the end,
// get all the rooms pointing to the end
// if a room from those room points somewhere else
// remove that link ?
// this is just a normal search, you can  use it to undestand
// the next method that is build on this one
// actually i don't even rememver what i did by this function
func (g *Graph) ValidPaths(end string) [][2]string {
	// make a que
	q := Q.New()
	q.Enqueue(g.Start)
	// make a map
	visited := make(map[*Vertex]bool)
	visited[g.Start] = true
	from := [][2]string{}
	// from[g.Start.Name] = g.Start.Name
	from = append(from, [2]string{g.Start.Name, g.Start.Name})
	// start from the q and gethem all
	// e is a node
	for !q.IsEmpty() {
		dequeuedItem := q.Dequeue()
		e, ok := dequeuedItem.Item.(*Vertex)
		if !ok {
			continue
		}
		for _, l := range e.adjacentVerteces {
			if visited[l] {
				continue
			}
			if l.Name != end {
				visited[l] = true
			}
			q.Enqueue(l)
			// from[l.Name] = e.Name
			from = append(from, [2]string{l.Name, e.Name})
		}
	}
	return from
}

// ///////////////////////////////////////////////////////
// this is bfs that allows me to find one path
func (g *Graph) FirstSet(name string, pas map[string]bool) []string {
	// make a que
	visited := copyMap(pas)
	q := Q.New()
	q.Enqueue(g.Start)
	// make a map
	// visited := make(map[*Vertex]bool)
	visited[g.Start.Name] = true
	var from [][2]string
	if g.Start.Name == name {
		return assemble(from, name)
	}
	// start from the q and gethem all
	// e is a node
	// i need a data type to store the path in.
	// var found = false
	for !q.IsEmpty() {
		found := false
		save := []*Vertex{}

		dequeuedItem := q.Dequeue()
		e, ok := dequeuedItem.Item.(*Vertex)
		if !ok {
			continue
		}
		///////////////////////////////////////
		for _, l := range e.adjacentVerteces {
			if visited[l.Name] {
				continue
			}
			if l.Name != name {
				visited[l.Name] = true
			} else {
				found = true
			}
			save = append(save, l)
			from = append(from, [2]string{e.Name, l.Name})
			if found {
				return assemble(from, name)
			}
		}
		if !found {
			for _, el := range save {
				q.Enqueue(el)
			}
		}
	}
	return assemble(from, name)
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

func (g *Graph) FindAllWays() [][]string {
	// find the first set
	name := g.End.Name
	var paths [][]string
	block := make(map[string]bool)
	stop := true
	for stop {
		ss := g.FirstSet(name, block)
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
	return paths
}

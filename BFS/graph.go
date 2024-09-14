package bfs 

import "fmt"
import Q "lem-in/queue"


type Graph struct{
	Verteces []*Vertex
	Start *Vertex // the starting room
	End *Vertex // the ending room
}

func NewGraph()*Graph{
	return &Graph{Verteces : []*Vertex{}, Start : nil , End : nil }
}


func (g *Graph)Add(v *Vertex){
	g.Verteces = append(g.Verteces, v)
}


func (g *Graph)Traverse(){
	// make a que
	q := Q.New()
	q.Enqueue(g.Start)
	// make a map 
	visited := make(map[*Vertex]bool)
	visited[g.Start] = true
	// start from the q and gethem all
	// e is a node
	for  !q.IsEmpty()  {
		dequeuedItem := q.Dequeue()
		e , ok := dequeuedItem.Item.(*Vertex)
		if !ok{
			continue
		}
		for _, l := range e.adjacentVerteces {
			if visited[l] {
				continue
			}
			visited[l] = true
			q.Enqueue(l)
			fmt.Println(l.Name)
		}
	}
}


func (g *Graph)Search(name string)*Vertex{
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
	for  !q.IsEmpty()  {
		dequeuedItem := q.Dequeue()
		e , ok := dequeuedItem.Item.(*Vertex)
		if !ok{
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
			fmt.Println(l.Name)
		}
	}
	return nil
}

// start from the end, 
// get all the rooms pointing to the end 
// if a room from those room points somewhere else 
// remove that link ? 

func (g *Graph)ValidPaths(end string)[][2]string{
	// make a que
	q := Q.New()
	q.Enqueue(g.Start)
	// make a map 
	visited := make(map[*Vertex]bool)
	visited[g.Start] = true
	from := [][2]string{}
	//from[g.Start.Name] = g.Start.Name
	from = append(from , [2]string{g.Start.Name, g.Start.Name})
	// start from the q and gethem all
	// e is a node
	for  !q.IsEmpty()  {
		dequeuedItem := q.Dequeue()
		e , ok := dequeuedItem.Item.(*Vertex)
		if !ok{
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
			//from[l.Name] = e.Name
			from = append(from , [2]string{l.Name, e.Name})
			fmt.Println(l.Name)
		}
	}
	return from
}




/////////////////////////////////////////////////////////
func (g *Graph)FirstSet(name string, visited map[string]bool)[][2]string{
	// make a que
	q := Q.New()
	q.Enqueue(g.Start)
	// make a map 
	//visited := make(map[*Vertex]bool)
	visited[g.Start.Name] = true
	var from [][2]string
	if g.Start.Name == name {
		return from
	}
	// start from the q and gethem all
	// e is a node
	// i need a data type to store the path in.
	//var found = false
	for  !q.IsEmpty() { 
		found := false
		save := []*Vertex{}

		dequeuedItem := q.Dequeue()
		e , ok := dequeuedItem.Item.(*Vertex)
		if !ok{
			continue
		}
		///////////////////////////////////////
		for _, l := range e.adjacentVerteces {
			if visited[l.Name] {
				continue
			}
			if l.Name != name {
				visited[l.Name] = true
			}else{
				found = true
			}
			save = append(save, l)
			from = append(from, [2]string{e.Name, l.Name})
		}
		if !found {
			for _, el := range save {
				q.Enqueue(el)
			}
		}
	}
	// instead of returning assemple and return 
	return from
	////////////////////////////////////////////
}

func Domino(parts [][2]string, exit string)[][]string{
	// start from the end 
	// look for the exit
	// check the first element of the last 
	// look for it's pair 
	// add to the path
	// remove the element
	var paths [][]string
	for i := len(parts) -1 ; i >= 0 ; i -- {
		if parts[i][1] == exit {
			paths = append(paths, assemble(parts[:i+1],exit))
		}
	}	


	return paths
}

func assemble(parts [][2]string, exit string)[]string{
	var find  string
	path := []string{}
	done := false
	for i := len(parts) - 1 ; i >= 0 ; i -- {
		if parts[i][1] == exit && !done {
			path = append(path, exit)
			path = append(path, parts[i][0])
			done = true
			find = parts[i][0]
			//parts = parts[:len(parts)-1]
		}
		if done {
			if parts[i][1] == find {
				path = append(path, parts[i][0])
				find = parts[i][0]
			//	parts = append(parts[:i] , parts[i+1:]...)
			}
		}

	}
	return path

}


func (g *Graph)FindAllWays(name string)[][]string{
	// find the first set 
	var paths [][]string
	block := make(map[string]bool)
	//block["v2"] = true 
	//block["v3"] = true 
	//block["v4"] = true
	var stop = true 	
	for stop {
		ss := Domino(g.FirstSet(name ,block), name) 
		if len(ss) != 0 {
			stop = false
		}
		paths = append(paths , ss... )
		for _, s := range ss {
			for _, t := range s {
				if t != name {
					block[t] = true
				}
			}
		}
	}
	return paths
}

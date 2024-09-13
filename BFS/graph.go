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

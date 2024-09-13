package queue


// define a type
type Node struct{
	Item any 
	Prev *Node
	Next *Node
}

type Queue struct {
	Head *Node
	Tail *Node
}

// return pointer or actual variable ? 
func New()*Queue{ // {addr 1001}// head
	return &Queue{Head: nil , Tail: nil }
} 

func (q *Queue)Remove(){
	if q.Head == nil {
		return 
	}	
	if q.Head.Next == nil {
		return 
	}
	q.Head = q.Head.Next
}

func (q *Queue)Add(item any){
	node := Node{Item : item , Prev : nil, Next : nil }
	if q.Head == nil {
		q.Head = &node
		q.Tail = &node
	}else{
		q.Tail.Next = &node
		node.Prev = q.Tail
		q.Tail = &node
		node.Next = nil
	}
}





















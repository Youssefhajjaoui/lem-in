package queue

import "testing"

func TestQueue(t *testing.T){
	q := New()
	// Test adding items 
	q.Add(1)
	q.Add([]int{2,3,4})
	q.Add("five")
	
	// 

}

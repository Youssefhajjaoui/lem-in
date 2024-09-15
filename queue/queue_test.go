package queue

import (
    "testing"
)

// TestQueue tests the Queue implementation
func TestQueue(t *testing.T) {
    q := New()

    // Test adding items to the queue
    q.Enqueue(0)
    q.Enqueue(1)
    q.Enqueue(2)

    // Define a helper function to check queue contents
    checkQueueContents := func(expected []any) {
        current := q.Head
        for i, exp := range expected {
            if current == nil {
                t.Errorf("Expected item %v at position %d, but got nil", exp, i)
                return
            }
            if current.Item != exp {
                t.Errorf("Expected item %v at position %d, but got %v", exp, i, current.Item)
            }
            current = current.Next
        }
        if current != nil {
            t.Error("Queue has more items than expected")
        }
    }

    // Check the contents of the queue
    checkQueueContents([]any{0, 1 , 2})
}

// TestEmptyQueue tests the behavior of the queue when empty
func TestEmptyQueue(t *testing.T) {
    q := New()

    // Check the queue contents
    if q.Head != nil {
        t.Error("Expected queue to be empty, but Head is not nil")
    }
    if q.Tail != nil {
        t.Error("Expected queue to be empty, but Tail is not nil")
    }
}


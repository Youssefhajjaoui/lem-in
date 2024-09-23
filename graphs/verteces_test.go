package graphs

import (
    "testing"
)

func TestNewVertex(t *testing.T) {
    v := New("Room A")
    if v.Name != "Room A" {
        t.Errorf("Expected Name to be 'Room A', but got %s", v.Name)
    }
}

func TestAddAdjacentVertex(t *testing.T) {
    v1 := New("Room A")
    v2 := New("Room B")

    v1.Add_adjacent_vertex(v2)

    if len(v1.adjacentVerteces) != 1 || v1.adjacentVerteces[0] != v2 {
        t.Errorf("Expected Room B to be adjacent to Room A")
    }

    if len(v2.adjacentVerteces) != 1 || v2.adjacentVerteces[0] != v1 {
        t.Errorf("Expected Room A to be adjacent to Room B")
    }
}

func TestInclude(t *testing.T) {
    v1 := New("Room A")
    v2 := New("Room B")
    v1.Add_adjacent_vertex(v2)

    if !v1.include(v2) {
        t.Errorf("Expected Room B to be included in Room A's adjacent vertices")
    }
    if v2.include(v1) != true {
        t.Errorf("Expected Room A to be included in Room B's adjacent vertices")
    }
}


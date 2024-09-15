package parse_file

import (
	"bufio"
	"errors"

	//	"fmt"
	"os"
	"strings"
	// bfs "lem-in/BFS"
)

// this file contains function to:
// 1- get the name of the file from the user
// 2- open and read the file
// 3- split the data of the file into 3
// // starting room
// // middle rooms
// // end room
// // links between rooms
// every function should return the error type
// so it can be used by others

func Getdata(filename string) ([]string, error) {
	result := []string{}
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("error from open")
	}
	Scanner := bufio.NewScanner(file)
	defer file.Close()
	for Scanner.Scan() {
		line := Scanner.Text()
		result = append(result, line)
	}
	return result, nil
}

func GetNodes(arr []string) ([]string, error) {
	Nodes := []string{}
	for _, v := range arr {
		if Len := len(strings.Split(v, " ")); Len == 3 {
			Nodes = append(Nodes, strings.Split(v, "")[0])
		}
	}
	return Nodes, nil
}

// func ParsetoNode(arr []string) ([]*Vertex, error) {
// 	values, err := GetNodes(arr)
// 	if err != nil {
// 		return nil, errors.New("NO Vertexes")
// 	}
// 	Nodes := []*Vertex{}
// 	for i := 0; i < len(values); i++ {
// 		g.Verteces = append(g.Verteces, &Vertex{
// 			Name: values[i],
// 		})
// 	}
// 	if len(Nodes) == 0 {
// 		return nil, errors.New("error in the Nodes details")
// 	}
// 	return Nodes, nil
// }

func GetEdges(arr []string) ([][]string, error) {
	cols := [][]string{}
	rows := []string{}
	for _, v := range arr {
		if Len := len(strings.Split(v, "-")); Len == 2 {
			rows = append(rows, strings.Split(v, "-")[0])
			rows = append(rows, strings.Split(v, "-")[1])
			cols = append(cols, rows)
			rows = []string{}
		}
	}
	if len(cols) == 0 {
		return nil, errors.New("maybe no relation here")
	}
	return cols, nil
}

func GetStart(arr []string) string {
	for i, v := range arr {
		if v == "##start" {
			return strings.Split(arr[i+1], "")[0]
		}
	}
	return ""
}

func GetEnd(arr []string) string {
	for i, v := range arr {
		if v == "##end" {
			return strings.Split(arr[i+1], "")[0]
		}
	}
	return ""
}

// func ProcessInput(filename string) (*Graphs, error) {
// 	g := &Graphs{}
// 	input, err := Getdata(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	Nodes, err := GetNodes(input)
// 	if err != nil {
// 		return nil, err
// 	}
// 	Rooms, err := g.ParsetoNode(Nodes)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err := g.GetEdges(input, Rooms); err != nil {
// 		return nil, err
// 	}
// 	//	fmt.Println(g->adjacentVerteces)
// 	//
// 	// fmt.Println(g)
// 	//
// 	return g, nil
// }

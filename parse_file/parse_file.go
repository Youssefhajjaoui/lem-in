package parse_file

import (
	"bufio"
	"errors"
//	"fmt"
	"os"
	"strings"

	//bfs "lem-in/BFS"
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

/*func ProcessInput(filename string, g *bfs.Graphs) {
	input, err := Getdata(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	Nodes, err := GetNodes(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	Rooms := g.ParsetoNode(Nodes)
	g.GetEdges(input, Rooms)
	//	fmt.Println(g->adjacentVerteces)
	//
	// fmt.Println(g)
	//
	g.PrintGraph()
}*/

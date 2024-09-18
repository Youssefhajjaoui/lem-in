package parse_file

import (
	"bufio"
	"errors"
	//	"fmt"
	"os"
	"strings"
)

type Nest struct {
	Rooms []string
	Tunels [2]string
	Start []string
	End []string
	Ants int
}


func FillTheNest(filename string) ([]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("error from open")
	}
	defer file.Close()

	result := []string{}
	Scanner := bufio.NewScanner(file)
	for Scanner.Scan() {
		line := strings.TrimSpace(Scanner.Text())
		if len(line) == 0 {continue}
		result = append(result, line)
	}
	nest := Parse(result)
	//return result, nil
}

func Parse(result []string)(Nest, error){
	var nest Nest
	if len(result == 0){
		return nest , errors.New("empty file")
	}
	nest.Ants , err := strconv.Atoi(result[0])
	if err != nil{
		return nest , errors.New("the first argument of the file should be the number of ants. Not Found !"
	}
	for i := 1 ; i < len(result []string) ; i ++ {
		arg := result[i]
		arg := strings.Fields(result[i])
		
	}
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


func GetEdges(arr []string) ([][]string, error) {
	cols := [][]string{}
	for _, v := range arr {
		rows := []string{}
		Len := strings.Split(v, "-")
		if len(Len) == 2 {
			rows = append(rows, Len[0])
			rows = append(rows, Len[1])
			cols = append(cols, rows)
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
			//return strings.Split(arr[i+1], " ")[0]
			return arr[i+1]
		}
	}
	return ""
}

func GetEnd(arr []string) string {
	for i, v := range arr {
		if v == "##end" {
			return strings.Split(arr[i+1], " ")[0]
		}
	}
	return ""
}

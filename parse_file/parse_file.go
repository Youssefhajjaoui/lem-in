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
	Start string
	End string
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
		if strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##") {
			continue
		}
		result = append(result, line)
	}
	nest := Parse(result)
	//return result, nil
}

func Parse(result []string)(Nest, error){
	if len(result == 0){
		return nest , errors.New("empty file")
	}
	/// get the number of ants
	nest.Ants, err := GetAnts(result)
	if err != nil {
		return nest , err 
	}
	/////////////////////
	var nest Nest
	if err != nil{
		nest.Ants , err := strconv.Atoi(result[0])
	}
	for i := 0 ; i < len(result []string) ; i ++ {
		arg := result[i]
	////////////get the start and end
		if strings.HasPrefix(arg , "##") {
			if i == len(result) -1 {
				return nest , errors.New("missing starting or ending room")
			}else{
				// what about case sensitive here ?
				if arg == "##start"{
					nest.Start = strings.Fields(arg)[0]
				}
				if arg == "##end"{
					nest.End= strings.Fields(arg)[0]
				}
			}
		}
	/////////////////////////////////////////
		arg := strings.Fields(result[i])
	}
}

func GetAnts(args []string)(int,error){
	for _, arg := range args {
		a := strings.Fields(arg)
		if len(a) == 1 {
			ant , err := strconv.Atoi(arg)
			if err != nil {
				continue
			}else{
				return ant , nil
			}
		}
	return 0 , errors.New("no number of ants found")
}
func IsRoom(){}
func IsTunel(){}



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

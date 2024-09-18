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
			continue
		}
	/////////////////////////////////////////
		Tor := strings.Fields(arg)
		if len(Tor) == 3 {
			Tunel , err := GetRoom(Tor)
			if err != nil{
				continue
			}
			// case room
			nest.Rooms = append(nest.Rooms, GetRoom(Tor))

		}else if len(arg) == 1 {
			Tunel , err := GetTunel(Tor)
			if err != nil{
				continue
			}
			nest.Tunels= append(nest.Tonels, GetTunel(Tor))
				
		}
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
func GetRoom(room []string)(string, error){
	if len(room)!= 3 {
		return "", errors.New("not room")
	}
	_, err := string.atoi(room[1])
	if err!= nil{
		return "", errors.new("not room")
	}
	_, err := string.atoi(room[2])
	if err!= nil{
		return "", errors.new("not room")
	}
	return room[0], nil
}
func GetTunel(tunel string)([2]string, error){
	t := strings.Split(tunel, "-")
	if len(tunel)!= 2 {
		return "", errors.New("not tunel")
	}
	return [2]string{t[0], t[1]}
}

package parse_file

import (
	"bufio"
	"errors"
	"strconv"

	"fmt"
	"os"
	"strings"
)

type Nest struct {
	Rooms  []string
	Tunels [][2]string
	Start  string
	End    string
	Ants   int
}

// fix the case of starnt and end
// remove them after you use them.
func FillTheNest(filename string) (Nest, error) {
	var nest Nest
	file, err := os.Open(filename)
	if err != nil {
		return nest, errors.New("error from open")
	}
	defer file.Close()

	result := []string{}
	Scanner := bufio.NewScanner(file)
	for Scanner.Scan() {
		line := strings.TrimSpace(Scanner.Text())
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "##") {
			continue
		}
		result = append(result, line)
	}
	nest, err = Parse(result)
	return nest, err
}

func Parse(result []string) (Nest, error) {
	fmt.Println(len(result))
	var nest Nest
	if len(result) == 0 {
		return nest, errors.New("empty file")
	}
	/// get the number of ants
	var err error
	nest.Ants, err = GetAnts(result)
	if err != nil {
		return nest, err
	}
	/////////////////////
	for i := 0; i < len(result); i++ {
		arg := result[i]

		// Check for start and end indicators
		if strings.HasPrefix(arg, "##") {
			// Ensure there's a next element to check
			if i == len(result)-1 {
				return nest, errors.New("missing starting or ending room")
			}

			nextArg := result[i+1] // Store next argument for easier access
			switch strings.ToLower(arg) {
			case "##start":
				nest.Start = strings.Fields(nextArg)[0]
				// Remove the start indicator and its following element
				result = append(result[:i], result[i+2:]...)
				i-- // Adjust index after modification
			case "##end":
				nest.End = strings.Fields(nextArg)[0]
				// Remove the end indicator and its following element
				result = append(result[:i], result[i+2:]...)
				i-- // Adjust index after modification
			}
			continue
		}

		/////////////////////////////////////////
		Tor := strings.Fields(arg)
		if len(Tor) == 3 {
			tunel, err := GetRoom(Tor)
			if err != nil {
				continue
			}
			// case room
			nest.Rooms = append(nest.Rooms, tunel)

		} else if len(arg) == 1 {
			tunel, err := GetTunel(arg)
			if err != nil {
				continue
			}
			nest.Tunels = append(nest.Tunels, tunel)

		}
	}
	return nest, nil
}

func GetAnts(args []string) (int, error) {
	for _, arg := range args {
		a := strings.Fields(arg)
		if len(a) == 1 {
			ant, err := strconv.Atoi(arg)
			if err != nil {
				continue
			} else {
				return ant, nil
			}
		}
	}
	return 0, errors.New("no number of ants found")
}
func GetRoom(room []string) (string, error) {
	if len(room) != 3 {
		return "", errors.New("not room")
	}
	_, err := strconv.Atoi(room[1])
	if err != nil {
		return "", errors.New("not room")
	}
	_, err = strconv.Atoi(room[2])
	if err != nil {
		return "", errors.New("not room")
	}
	return room[0], nil
}
func GetTunel(tunel string) ([2]string, error) {
	t := strings.Split(tunel, "-")
	if len(tunel) != 2 {
		return [2]string{}, errors.New("not tunel")
	}
	return [2]string{t[0], t[1]}, nil
}

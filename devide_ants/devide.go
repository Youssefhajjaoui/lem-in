package devide 

import "fmt"

/*L1-3 L2-2
L1-4 L2-5 L3-3
L1-0 L2-6 L3-4
L2-0 L3-0*/
//x, z, r represents the ants numbers (going from 1 to number_of_ants) and y, w, o represents the rooms names.



//[[0 4 3 1] [0 6 5 2 1]]
////////////////////////////////////////////////////////////
type Path struct {
	Rooms []string
	Passenger int
} 
func NewPath()*Path{
	return &Path{Rooms : []string{}, Passenger : 0 }
}
func (p *Path)Add(room string){
	(*p).Rooms = append((*p).Rooms, room)
}
func (p *Path)Pass(){
	p.Passenger ++
}

//////////////////////////////////////////////////////////
type Paths []*Path 
// add to the path ? 
// declare Paths 
func NewPaths()*Paths{
	return &Paths{}
}
func (p *Paths)Append(path *Path){
	*p = append(*p , path)
}
func (p *Paths)Shortest()(*Path, bool){
	signle := false
	// this has to be the first 
	var shortest = (*p)[0]
	for _, path := range *p{
		if path.Passenger <= shortest.Passenger {
			signle = true
			shortest = path
		}		
	}
	return shortest, signle
}
//////////////////////////////////////////////////////////


/////////////////////////////////////////////////////////


func Devide( ways [][]string, ants int)[][]string{
	// make new paths
	paths := NewPaths()
	for _ , p := range ways{
		// make new path 
		path := NewPath()
		for _, room := range p{
			path.Add(room)
			path.Pass()
		}
		paths.Append(path)
	}

	
	show := [][]string{}
	for i := 1 ; i <= ants ; i ++ {
		short , _ := paths.Shortest()
		short.Pass()
		took := antpath(i, short.Rooms)
		show = append(show , took)	
	}
	mat := Retate(show)
	return mat
}


func antpath(ant int, path []string)[]string{
	took := []string{}
	for i:= len(path) - 2 ; i >= 0 ; i -- {
		room := path[i]
		took = append(took, fmt.Sprintf("L%d-%s", ant, room))

	}
	return took
}
func Retate(matrix [][]string)[][]string{
	result := [][]string{}
	stop := true
	y := 0 
	for stop {
		stop = false
		line := []string{}
		for i := 0 ; i < len(matrix)  ; i ++ {
			branch := matrix[i]
			//fmt.Println(branch)
			if len(branch) > y {
				stop = true
				if !Check(line , branch[y]){
					line= append(line, branch[y])
				}else{
					matrix[i] = append([]string{""} , branch...)
				}
			}
		}
		result = append(result , line)
		y ++
	}
	return result
}


func Check(line []string, s string)bool{
	r := s[len(s)-1] 
	for i := 0 ; i < len(line) -1  ; i ++ {
		b := line[i][len(line[i])-1] 
		if b == r {
			return true
		}
	}
	return false
}

func Print(mat [][]string){
	for _, line := range mat {
		if len(line) == 0 {
			continue
		}
		for i , word := range line {
			if word != "" {
				fmt.Print(word)
				if i != len(line) -1 {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}

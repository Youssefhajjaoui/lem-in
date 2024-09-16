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
func (p *Paths)Shortest()*Path{
	// this has to be the first 
	var shortest = (*p)[0]
	for _, path := range *p{
		if path.Passenger <= shortest.Passenger {
			shortest = path
		}		
	}
	return shortest
}
//////////////////////////////////////////////////////////


/////////////////////////////////////////////////////////


func Devide( ways [][]string, ants int)/*map[string]int*/{
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

	
	for i := 1 ; i <= ants ; i ++ {
		short := paths.Shortest()
		short.Pass()
		antpath(i, short.Rooms)
	}






/*	shortest := len(paths[0])
	// choos the path for the ant

	// send the ant
	ant := 1
	take := paths[0]
	for _, path := range paths {
		if len(path) < shortest {
			take = path
		}
	}
	shortest ++*/
}


func antpath(ant int, path []string){
	for i:= len(path) - 2 ; i >= 0 ; i -- {
		room := path[i]
		fmt.Printf("L%d-%s", ant , room )
		fmt.Println()
	}
}


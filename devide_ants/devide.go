package devide 

import "fmt"

/*L1-3 L2-2
L1-4 L2-5 L3-3
L1-0 L2-6 L3-4
L2-0 L3-0*/
//x, z, r represents the ants numbers (going from 1 to number_of_ants) and y, w, o represents the rooms names.



//[[0 4 3 1] [0 6 5 2 1]]

func Devide( paths [][]string, ants int)/*map[string]int*/{
	// send all the ans in the same derection// 
	for j := 1 ; j <= ants ; j ++ {
		for i:= len(paths[0]) -1 ; i >= 0 ; i --  {
			fmt.Printf("L%d-%s ", j , paths[0][i] )	
			fmt.Println()
		}
	}
}



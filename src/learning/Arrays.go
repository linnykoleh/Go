package main

import "fmt"

func main(){
	var array1 [2]bool
	var array2 =[...]int {1, 2, 3, 4, 5, 6}
	array3 :=[2][2]int {{1, 2}, {}}

	fmt.Println("array1: ", len(array1), array1)
	fmt.Println("array2: ", len(array2), array2)
	fmt.Println("array3: ", len(array3), array3)

	fmt.Println("access by index: ", array3[1][0])

}

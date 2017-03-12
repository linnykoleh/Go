package main

//Срезы
//имеют переменную длину

import "fmt"

func main()  {

	var firstSlice = make([]byte, 5, 15) // []byte-type, 5-size, 15-capacity
	var secondSlice = make([]byte, 10)   // []byte-type, 10-size, 10-capacity
	var thirdSlice []byte
	var fourthSlice = []int{1,2,3,4,5}

	fmt.Println("firstSlice:", len(firstSlice), cap(firstSlice), firstSlice)
	fmt.Println("secondSlice:", len(secondSlice), cap(secondSlice), secondSlice)
	fmt.Println("thirdSlice:", len(thirdSlice), cap(thirdSlice), thirdSlice)
	fmt.Println("fourthSlice:", len(fourthSlice), cap(fourthSlice), fourthSlice)

	fmt.Println("\n\t Part of slice")

	part1 := fourthSlice[:2]  // from index = 0 to index 2
	part2 := fourthSlice[2:3] // from index = 2 to index 3
	part3 := fourthSlice[3:]  // from index = 3 to length - 1
	part4 := fourthSlice[:]   // take all

	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
	fmt.Println("part3:", part3)
	fmt.Println("part4:", part4)

	fmt.Println("\n\t Copy of slice")

	copyOfFourth := make([]int, 10)
	copiedElementsNumber := copy(copyOfFourth, fourthSlice)

	fmt.Println("copiedElementsNumber:", copiedElementsNumber)
	fmt.Println("copyOfFourth:", copyOfFourth)

	fmt.Println("\n\t Append  of slice")

	copyOfFourth  = append(copyOfFourth, fourthSlice[:3]...)
	fmt.Println("copyOfFourth:", copyOfFourth)

	copyOfFourth  = append(copyOfFourth, 100, 200, 300)
	fmt.Println("copyOfFourth:", copyOfFourth)

}

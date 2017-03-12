package main

import "fmt"

func main() {
	first := "first line1 \t\n line2"
	second := `second line \t\n line`

	fmt.Println("first : ", first)
	fmt.Println("second : ", second)

	stringOperation()
}

func stringOperation(){

	fmt.Println("\n\t String operations")

	first := "first"
	fmt.Println("first[1] : ", first[1]) // 105
	fmt.Println("string(first[1]) : ", string(first[1])) // i

	fmt.Println("first[:2] : ", first[:2]) // fi
	fmt.Println("first[3:] : ", first[3:]) // st
	fmt.Println("first[1:3] : ", first[3:]) // ir
}

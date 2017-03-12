package main

import "fmt"

func main() {

	var first = make(map[bool]float32, 20)
	var second = make(map[bool]uint)
	third := map[string]int{"one" : 1, "two" : 2, "three" : 3}
	fourth := map[int]string{}


	fmt.Println("first : ", first)
	fmt.Println("second : ", second)
	fmt.Println("third : ", third)
	fmt.Println("fourth : ", fourth)

	third["fourth"] = 4

	fmt.Println("third added : ", third)

	delete(third, "one")

	fmt.Println("third deleted: ", third)


}

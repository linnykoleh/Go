package main

import "fmt"

func main() {

	fmt.Println("First: ")
	first := 1
	for{ // infitity loop
		fmt.Print(".")
		if first > 10{
			break
		}
		first++
	}

	fmt.Println("\nSecond: ")
	for second := 10; second < 20; second++ {
		fmt.Print(". ")
	}

	fmt.Println("\nThird: ")
	third :=[]string{"a", "b", "c", "d", "e"}

	for index := range third{
		fmt.Println(index, "->", third[index])
	}

	fmt.Println("\nFourth: ")
	fourth := map[string]string{
			"1" : "a",
			"2" : "b",
			"3" : "c",
			"4" : "d"}

	for key, value := range fourth{
		fmt.Println(key, "->", value)
	}

	fmt.Println("\nFifth: ")

	for key, value := range third{
		fmt.Println(key, "->", value)
	}
}

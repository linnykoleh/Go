package main

import "fmt"

func init(){
	//That function executes before main
	fmt.Println("Init")
}

func main() {
	fmt.Println("Main")
}

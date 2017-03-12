package main

import (
	"fmt"
	"time"
)

func first(){
	time.Sleep(100)
	fmt.Println("first")
}

func second(){
	time.Sleep(200)
	fmt.Println("second")
}

func third(){
	time.Sleep(300)
	fmt.Println("third")
}

func fourth(){
	time.Sleep(400)
	fmt.Println("fourth")
}

func main() {

	fmt.Println("Main begin")

	defer first()
	defer second()
	defer third()
	defer fourth()

	fmt.Println("Main end")
}

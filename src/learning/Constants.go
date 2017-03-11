package main

import (
	"fmt"
	)

const SingleLineConst  =  "Constant"

const (
	MultiFirstConstant = 1 //1
	MultiSecondConstant    //1
)

const (
	MultiFirstConstantIOTA = iota //0
	MultiSecondConstantIOTA       //1
)

func main() {
	fmt.Println(SingleLineConst)
	fmt.Println(MultiFirstConstant)
	fmt.Println(MultiSecondConstant)

	fmt.Println(MultiFirstConstantIOTA)
	fmt.Println(MultiSecondConstantIOTA)
}
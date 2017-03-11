package main

import (
	"fmt"
	"math"
	"reflect"
)

func main()  {

	var floatDefault32 float32
	var floatDefault64 float64
	floatNumber := 0.1

	fmt.Println("float32: ", floatDefault32)
	fmt.Println("float64: ", floatDefault64)

	fmt.Println("floatNumber: ", reflect.TypeOf(floatNumber), floatNumber)

	fmt.Println("Max float: ", math.MaxFloat32)
	fmt.Println("MIn float: ", math.SmallestNonzeroFloat32)

}


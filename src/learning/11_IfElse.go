package main

import (
	"fmt"
	"os"
)

func main() {

	trueVar := true
	falseVar := false

	if trueVar {
		fmt.Println("True")
	}

	if !falseVar {
		fmt.Println("True")
	}

	if falseVar {
		fmt.Println("false")
	}else if trueVar {
		fmt.Println("True")
	}else{
		fmt.Println("Error")
	}

	if 2 >= 1 && 1 == 1 {
		fmt.Println("2 >= 1 && 1 == 1\n")
	}

	if data, err := os.Stat("asasd"); data == nil{
		fmt.Println("data : ", data)
		fmt.Println("err : ", err)
	}
}

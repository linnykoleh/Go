package main

import "fmt"

// Указатели - переменные, которые хранят ссылки на другие переменные такие как
//
// * среза
// * отображения
// * каналы
// * функции
// * указатели
//
// & - взятие ссылки
// * - взятие значения

func main() {

	intVar := 10
	firstPointer := &intVar
	secondPointer := & firstPointer

	fmt.Println("intVar", &intVar, intVar)
	fmt.Println("firstPointer", firstPointer, &firstPointer, firstPointer)
	fmt.Println("firstPointer", *firstPointer, &*firstPointer, *firstPointer)

	fmt.Println("secondPointer", secondPointer, &secondPointer, secondPointer)
	fmt.Println("secondPointer", *secondPointer, &*secondPointer, *secondPointer)
	fmt.Println("secondPointer", **secondPointer, &**secondPointer, **secondPointer)

	fmt.Println("------------------")

	anotherPointer := 500

	fmt.Println("intVar : ", intVar)                  //10
	fmt.Println("anotherPointer : ", anotherPointer)  //500
	fmt.Println("*firstPointer : ", *firstPointer)    //10

	fmt.Println("------------------")

	firstPointer = &anotherPointer
	*firstPointer += 10

	fmt.Println("intVar : ", intVar)                 //10
	fmt.Println("anotherPointer : ", anotherPointer) //510
	fmt.Println("*firstPointer : ", *firstPointer)   //510

}

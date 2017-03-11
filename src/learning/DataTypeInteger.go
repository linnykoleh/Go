package main

import (
	"fmt"
	"math"
)

func main()  {
	var typeInt int64 = 5
	number := 10

	fmt.Println("typeInt", typeInt)
	fmt.Println("number", number)

	integersSizes()
}

func integersSizes()  {

	fmt.Println("\n \t Пределы значений целых чисел \n")
	fmt.Println("int8 :", math.MinInt8, "-", math.MaxInt8)
	fmt.Println("int16 :", math.MinInt16, "-", math.MaxInt16)
	fmt.Println("int32 : ", math.MinInt32, "-", math.MaxInt32)
	//fmt.Println("int64", math.MinInt64, "-", math.MaxInt64) // too big

	fmt.Println("\n uint8 :",0, "-", math.MaxUint8)
	fmt.Println("uint16 :", 0, "-", math.MaxUint16)
	//fmt.Println("uint32 : ", 0, "-", math.MaxUint32) // too big
	//fmt.Println("uint64", 0, "-", math.MaxUint64)    // too big

	fmt.Println("\n \t Синонимы целух типов \n")

	var intByte byte = 0
	var intRune rune = 0
	var intInt int = 0
	var intUInt uint = 0

	fmt.Println("byte  = int8 ", intByte)
	fmt.Println("rune  = int32 ", intRune)
	fmt.Println("int  = int32 or int64 (by system) ", intInt)
	fmt.Println("uint  = uint32 or uint64 (by system) ", intUInt)

}

package main

import (
	"fmt"
	"strings"
)

func printData(printData string){
	fmt.Println(printData)
}

func concatenation(agrs ...string) string{
	var res string
	for _, value:= range agrs{
		res += value
	}
	return res
}

//example return function
func wrap(value string) (func(userValue string)string){
	return func(userValue string) string {
		return value + userValue
	}
}

//example use function as arguments
func convert(value string, someFunctions ...func(data string) string) string{
	for _, function := range someFunctions{
		value = function(value)
	}
	return value
}

func main() {

	printData("Hello")

	concatenation := concatenation("He", "llo", " ,", "Wo", "rld")
	printData(concatenation)

	result := wrap("MyValue ")("User Value")
	printData(result)

	convertedValue := convert(" TestValue ", strings.ToLower, strings.TrimSpace)
	printData(convertedValue)
}

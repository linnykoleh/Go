package main

import "fmt"

func main() {

	fmt.Print("First: ")
	switch { // boolean switch
		case 1 < 2:
			fmt.Println("1 < 2")
	        case 2 < 3:
			fmt.Println("2 < 3")
	}

	fmt.Print("Second: ")
	imageVar := "jpg"
	switch imageVar{
	case "gif":
		fallthrough
	case "jpeg":
		fallthrough
	case "jpg":
		fmt.Println("Image")
	default:
		fmt.Println("Not image")
	}

	fmt.Print("Third: ")
	switch imageVarShort := "exe"; imageVarShort{
	case "gif":
		fallthrough
	case "jpeg":
		fallthrough
	case "jpg":
		fmt.Println("Image")
	default:
		fmt.Println("Not image")
	}

	fmt.Print("Fourth: ") // comparing by type
	var someObject interface{}

	switch someObject.(type){
		case bool:
			fmt.Print("Boolean")
		case int:
			fmt.Print("Integer")
		case string:
			fmt.Print("String")
		default:
			fmt.Print("not type")
	}
}

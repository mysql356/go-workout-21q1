package main

import "fmt"

func findType(i interface{}) {

	switch i.(type) {
	case string:
		fmt.Println(i.(string))

	case int:
		fmt.Println(i.(int))

	default:

		fmt.Println("undefined")
	}
}

func main() {

	findType("aaa")

	findType(10)

}

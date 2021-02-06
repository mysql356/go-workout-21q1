package main

import "fmt"

func main() {

	//interface-array
	arr := []interface{}{1, "aa", 2.5, nil}
	fmt.Println(arr)
	fmt.Printf("%T \n", arr)

	show(arr)

}

func show(b []interface{}) {
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}
}

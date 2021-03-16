package main

import "fmt"

func main() {

	//string
	var a interface{}
	a = "aa"
	fmt.Println(a.(string))
	fmt.Printf("a ======>%s %T \n", a, a)

	//int
	var b interface{}
	b = 20
	fmt.Println(b.(int))

	//string-1
	var s interface{} = "aa"
	fmt.Println(s.(string))

	//interface-array
	arr := [4]interface{}{1, "aa", 2.5, nil}
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	//interface-in-function
	arr2 := []interface{}{1, "aa", 2.5, nil}
	fmt.Println(arr2)
	fmt.Printf("%T \n", arr2)

	show(arr2)

}

func show(b []interface{}) {
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}
}

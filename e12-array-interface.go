package main

import "fmt"

func main() {
	a := [3][2]string{{"aa", "bb"}, {"mm", "nn"}, {"xx", "yy"}}
	arrPrint(a)

	b := [4]interface{}{1, "hello", 2.5, nil}
	infPrint(b)
}

func arrPrint(a [3][2]string) {

	for _, v := range a {
		for _, v1 := range v {
			fmt.Print(v1, "===")
		}
		fmt.Println()
	}
}

func infPrint(b [4]interface{}) {
	for k, v := range b {

		switch v.(type) {
		case string:
			fmt.Println(k, "=====", v.(string))

		case int:
			fmt.Println(k, "=====", v.(int))

		case bool:
			fmt.Println(k, "=====", v.(bool))

		case nil:
			fmt.Println(k, "=====", nil)

		}
	}
}

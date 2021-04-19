package main

import "fmt"

func main() {
	print := fmt.Println

	print("Hello " + "World")
	t := fmt.Sprintf("http://localhost/manojk/editorModule/%s", "abcd.txt")
	fmt.Println(t)
}

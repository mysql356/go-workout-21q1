package main

import "fmt"

func main() {
	var no int
	fmt.Println("Enter the no: ")
	fmt.Scanf("%d", &no)

	reminder := no % 4
	switch reminder {

	case 0:
		fmt.Printf("Reminde : 0")

	case 1:
		fmt.Printf("Reminder: 1")

	case 2:
		fmt.Printf("Reminder : 2")

	case 3:
		fmt.Printf("Reminder : 3")

	default:

		fmt.Println("default")
	}
}

package main

import "fmt"

func main() {
	var no int
	fmt.Println("Enter the no: ")
	//fmt.Scanf("%d", &no)
	fmt.Scan(&no)

	reminder := no % 4
	switch reminder {

	case 0:
		fmt.Println("Reminde : 0")

	case 1:
		fmt.Println("Reminder: 1")

	case 2:
		fmt.Println("Reminder : 2")

	case 3:
		fmt.Println("Reminder : 3")

	default:

		fmt.Println("default")
	}
}

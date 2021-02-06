package main

import "fmt"

type Address struct{ city, state string }
type Contact struct{ email, mobile string }
type Person struct {
	name    string
	age     int
	address Address
	contact []Contact
}

func main() {
	var p Person
	p.name = "aa"
	p.age = 20
	p.address = Address{
		city:  "ND",
		state: "Delhi",
	}

	p.contact = []Contact{
		{"user1@gmail.com", "9000000001"},
		{"user2@gmail.com", "9000000002"},
	}

	p.contact = []Contact{

		Contact{"user3@gmail.com", "9000000003"},
	}

	fmt.Println(p.name, p.age)
	fmt.Println(p.address.city)
	fmt.Println(p)

	/////////////////////
	p1 := Person{
		name: "bb",
		age:  30,
		address: Address{
			city:  "Patna",
			state: "Bihar",
		},

		contact : []Contact{
			{"user1@gmail.com", "9000000001"},
			{"user2@gmail.com", "9000000002"},
		},
	}

	fmt.Println(p1.name, p1.age)
	fmt.Println(p1.address.city)
	fmt.Println(p1)

}

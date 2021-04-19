package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Employee struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}

func main() {

	//map-to-json : output
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"a": 5, "b": 10}
	fmt.Println("-map-to-json-----------")
	enc.Encode(d)

	//by value
	emp1 := new(Employee)
	emp1.FirstName = "Manoj"
	emp1.LastName = "kumar"
	emp1.City = "Delhi"

	//struct-to-json
	fmt.Println("\n-struct-to-json-----------")
	enc.Encode(emp1)

	fmt.Println("\n-Marshal -----------")
	jsonStr, _ := json.Marshal(emp1)
	//fmt.Println(jsonStr)         //[123 34 102 105 114 ... 105 34 125]
	fmt.Println(emp1)            //&{Manoj kumar Delhi}
	fmt.Println(*emp1)           //{Manoj kumar Delhi}
	fmt.Println(string(jsonStr)) //{"firstname":"Manoj","lastname":"kumar","city":"Delhi"}

	fmt.Println("\n-Unmarshal -----------")
	emp2 := new(Employee)
	json.Unmarshal(jsonStr, emp2)
	fmt.Println(emp2)
	fmt.Println(emp2.FirstName, emp2.LastName, emp2.City)

	//by string
	json_string := `
    {
        "firstname": "Rahul",
        "lastname": "Singh",
        "city": "Mumbai"
    }`
	fmt.Println("\n-Unmarshal :by val-----------")
	emp3 := new(Employee)
	json.Unmarshal([]byte(json_string), emp3)
	fmt.Println(emp3)

	fmt.Println("\n-Unmarshal :by ref-----------")
	emp4 := Employee{}
	json.Unmarshal([]byte(json_string), &emp4)
	fmt.Println(emp4)

}

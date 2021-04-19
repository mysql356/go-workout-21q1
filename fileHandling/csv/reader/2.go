package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	// Load a TXT file.
	f, _ := os.Open("animal.csv")

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))

	s := ""
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		//fmt.Println(record)
		//fmt.Println(len(record))

		for _, v := range record {
			//fmt.Printf("  %v\n", record[value])
			s = fmt.Sprintf("%s,%s", s, v)
		}

	}

	fmt.Println(s[1:])
	fmt.Println(s)
}

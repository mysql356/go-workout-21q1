package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Simple static webserver:
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("/dev"))))

	fmt.Println(http.ListenAndServe(":3002", mux))
}

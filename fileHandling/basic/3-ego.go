package main

import (
 "flag"
 "log"
 "net/http"
)

func main() {
 port := flag.String("p", "3000", "port to serve on")
 directory := flag.String("d", ".", "the directory of static file to host")
 flag.Parse()

 handler := http.FileServer(http.Dir(*directory))
 http.Handle("/", handler)

 log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
 log.Fatal(http.ListenAndServe(":"+*port, nil))
}
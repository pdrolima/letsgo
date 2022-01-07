package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func printHelloWorld(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1>Hello World! The time is: %s </h1>", t)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.Handle("/hello", http.HandlerFunc(printHelloWorld))

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

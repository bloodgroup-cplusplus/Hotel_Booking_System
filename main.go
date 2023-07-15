package main

import (
	"fmt"
	"net/http"
)

var portNumber = ":8080"

// let's pretend we are building website of two pages

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "This is the home page")

}

// About is the about page handler

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2+2 is %d", sum))

}

// add values add two integers and returns the sum

func addValues(x, y int) int {
	return x + y
}

// main is the main entry point

func main() {
	// let's create a new function
	// use builtin http package

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	//n, err := fmt.Fprintf(w, "Hello, World!")
	//if err != nil {
	//fmt.Println(err)
	//}
	//fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))

	//})
	//_ = http.ListenAndServe(":8080", nil)

	// Routing

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	_, _ = fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)

}

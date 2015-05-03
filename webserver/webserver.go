package main

import (
	"fmt"
	//	"log"
	"net/http"
)

// Hello hello
type Hello struct{}

// String string
type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

// main
// func main() {
// 	/*	var h Hello
// 		http.Handle("/string", String("I'm a frayed knot."))
// 		http.Handle("/", h)
// 		err := http.ListenAndServe("localhost:4000", nil)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	*/
// 	m := martini.Classic()
// 	m.Get("/", func() string {
// 		return "Hello world!"
// 	})
// 	m.Run()
// }

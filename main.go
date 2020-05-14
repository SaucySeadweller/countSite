package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

//NewPage is struct hurr durr
type NewPage struct {
	Title   string
	Message string
}

func counter(w http.ResponseWriter, r *http.Request) {
	p := NewPage{Title: "Click the button"}
	t, err := template.ParseFiles("basictemplating.html")
	log.Println(err)
	t.Execute(w, p)
}
func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World <h1>")
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/counter/", counter)
	http.ListenAndServe(":8080", nil)
}

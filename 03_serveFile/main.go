package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/cat/", cat)
	http.HandleFunc("/cat.jpg", peach)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran\n")
}

func cat(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("cat.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "cat.gohtml", nil)
}

func peach(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "cat.jpg")
}
package main

import (
	"net/http"
	"html/template"
	"log"
)

var tpl *template.Template

// init will excute before main func
func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	/*fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", fs)*/
	// use StripPrefix way to do it
	http.HandleFunc("/", dogs)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("public"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dogs(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)


// using DefaultServeMux
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", bar)
	http.HandleFunc("/me/", myName)
	/* using Handle with HandlerFunc
	http.Handle("/", http.HandlerFunc(foo))
	http.Handle("/dog/", http.HandlerFunc(bar))
	http.Handle("/me/", http.HandlerFunc(mcleod))
	*/
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "this is index")
}

func bar(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "bar ran")
}

func myName(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("something.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}
	err = tpl.ExecuteTemplate(w, "something.gohtml", "Kevin")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

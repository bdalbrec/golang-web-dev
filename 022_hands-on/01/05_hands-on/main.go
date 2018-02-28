package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func index(w http.ResponseWriter, res *http.Request) {
	err := tpl.ExecuteTemplate(w, "tpl.gohtml", "home page")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func dog(w http.ResponseWriter, res *http.Request) {
	err := tpl.ExecuteTemplate(w, "tpl.gohtml", "dog dog doggy")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func me(w http.ResponseWriter, res *http.Request) {
	err := tpl.ExecuteTemplate(w, "tpl.gohtml", "Brian Albrecht")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}

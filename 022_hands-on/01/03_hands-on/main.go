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
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}

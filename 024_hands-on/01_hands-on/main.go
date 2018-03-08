package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/dog.jpg", handleDog)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>This is from dog</h1>")

	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(w, "dog.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func handleDog(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpg")
}

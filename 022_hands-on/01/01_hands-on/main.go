package main

import (
	"io"
	"net/http"
)

func index(w http.ResponseWriter, res *http.Request) {
	io.WriteString(w, "home page")
}

func dog(w http.ResponseWriter, res *http.Request) {
	io.WriteString(w, "dog dog doggy")
}

func me(w http.ResponseWriter, res *http.Request) {
	io.WriteString(w, "My name is Brian")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}

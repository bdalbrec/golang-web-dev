package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", cook)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func cook(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("counter")

	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "counter",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	c.Value = strconv.Itoa(count)
	http.SetCookie(w, c)

	fmt.Fprintln(w, "You have visited this site", count, "times.")
}

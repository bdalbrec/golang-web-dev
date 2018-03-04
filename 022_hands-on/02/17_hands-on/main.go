package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go serve(conn)

		// we never get here
		// we have an open stream connection
		// how does the above reader know when it's done?
		fmt.Println("Code got here.")

	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	i := 0
	var m, u, body string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			sf := strings.Fields(ln)
			m = sf[0]
			u = sf[1]
		}
		if ln == "" {
			// according to http protocol the last line of the headers is a blank line
			fmt.Println("***END OF HTTP REQUEST HEADERS***")
			fmt.Printf("Received request with METHOD: %s\n", m)
			fmt.Printf("Received request with URL: %s\n", u)
			break
		}
		i++
	}

	if m == "GET" && u == "/" {
		body = fmt.Sprintf(`<!DOCTYPE html><html lang="en">
			<head><meta charset="UTF-8"><title>A Webpage</title></head>
			<body><h1>HOLY COW THIS IS LOW LEVEL:</h1><h2>METHOD: %s</h2><h2>URL: %s</h2></body></html>`, m, u)
	}

	if m == "GET" && u == "/dog" {
		body = fmt.Sprintf(`<!DOCTYPE html><html lang="en">
			<head><meta charset="UTF-8"><title>A Webpage</title></head>
			<body><h1>HOLY COW THIS IS LOW LEVEL:</h1><h2>METHOD: %s</h2><h2>URL: %s</h2></body></html>`, m, u)
	}

	if m == "POST" && u == "/dog" {
		body = fmt.Sprintf(`<!DOCTYPE html><html lang="en">
			<head><meta charset="UTF-8"><title>A Webpage</title></head>
			<body><h1>HOLY COW THIS IS LOW LEVEL:</h1><h2>METHOD: %s</h2><h2>URL: %s</h2></body></html>`, m, u)
	}

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

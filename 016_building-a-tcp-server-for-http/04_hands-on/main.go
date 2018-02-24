package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	//read request
	url := request(conn)

	// write response
	respond(conn, url)
}

func request(conn net.Conn) string {
	i := 0
	var u string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			u = strings.Fields(ln)[1]
			fmt.Println("***URL", u)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
	return u
}

func respond(conn net.Conn, url string) {
	var bodyText string

	switch url {
	case "/":
		bodyText = home()
	case "/about":
		bodyText = about()
	case "/contact":
		bodyText = contact()
	default:
		bodyText = notFound()
	}

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Test http server</title></head><body>` + bodyText + `</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func home() string {
	return "Hello World"
}

func about() string {
	return "A little about me:"
}

func contact() string {
	return "My contact info:"
}

func notFound() string {
	return "404"
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
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
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			// according to http protocol the last line of the headers is a blank line
			fmt.Println("***END OF HTTP REQUEST HEADERS***")
			break
		}
	}
	io.WriteString(conn, "I see you connected.")
}

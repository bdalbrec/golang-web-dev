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
		io.WriteString(conn, "I see you connected.")

		conn.Close()
	}
}

func serve(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			// after the headers is a blank line according to http protocol
			fmt.Println("***END OF HTTP REQUEST HEADERS***")
			break
		}
		fmt.Println(ln)
	}
	defer conn.Close()
}

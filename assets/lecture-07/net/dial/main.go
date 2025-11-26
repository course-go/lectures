package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// START OMIT

func main() {
	conn, err := net.Dial("tcp", "course-go.dev:80")
	if err != nil {
		log.Fatalf("failed dialing: %v", err)
	}

	defer conn.Close()
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	resp, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatalf("failed reading from connection: %v", err)
	}

	fmt.Println(resp)
}

// END OMIT

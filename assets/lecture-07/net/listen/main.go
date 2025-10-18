package main

import (
	"fmt"
	"log"
	"net"
)

// START OMIT

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("could not listen on port 8000: %v", err)
	}

	fmt.Println("Listening on port 8080...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("could not accept new connection: %v", err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	log.Printf("Handling connection from: %s", conn.RemoteAddr().String())
	conn.Write([]byte("HTTP/1.0 200 OK\r\n\r\nHello\r\n"))
	conn.Close()
}

// END OMIT

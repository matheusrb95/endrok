package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	log.Printf("Client %s connected", c.RemoteAddr())

	buf := make([]byte, 1024)

	for {
		n, err := c.Read(buf)
		if err != nil {
			log.Printf("Error reading from client %s: %s", c.RemoteAddr(), err.Error())
			return
		}

		if n > 0 {
			log.Printf("%s: %s", c.RemoteAddr(), string(buf[:n]))
		}
	}
}
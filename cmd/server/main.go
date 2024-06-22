package main

import "log"

func main() {
	server := NewServer("192.168.15.4:8000")

	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}

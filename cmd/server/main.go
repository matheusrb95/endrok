package main

import "log"

func main() {
	server := NewServer(":8000")

	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}

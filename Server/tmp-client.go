package main

import (
	"net"
	"log"
)

func main() {
	conn, err := net.Dial("tcp", ":7000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("this is tmp client"))
	if err != nil {
		log.Println(err)
	}
}
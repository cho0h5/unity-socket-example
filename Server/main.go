package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":7000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		go func() {
			buffer := make([]byte, 1000)

			for {
				count, err := conn.Read(buffer)
				if err != nil {
					log.Println("Exit")
					break
				}
				log.Println(string(buffer[:count]))

			}
		}()
	}

}

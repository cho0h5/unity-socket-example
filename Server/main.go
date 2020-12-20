package main

import (
	"encoding/json"
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
				// Receive data
				count, err := conn.Read(buffer)
				if err != nil {
					log.Println("Exit")
					break
				}

				// Print data
				log.Println(string(buffer[:count]))

				// Unmarshal received data
				var recvData map[string]interface{}
				json.Unmarshal(buffer[:count], &recvData)

				// Marshal data
				recvData["Count"] = recvData["Count"].(float64) * 2
				data, _ := json.Marshal(recvData)

				// Response data
				_, err = conn.Write(data)
				if err != nil {
					log.Println(err)
				}
			}
		}()
	}

}

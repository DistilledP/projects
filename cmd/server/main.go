package main

import (
	"fmt"
	"log"
	"net"

	"github.com/DistilledP/lungfish/internal/libs"
)

func main() {
	config := libs.GetServices().GetConfig()

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", config.PublicPort))
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("ERR accepting connection:", err)
			continue
		}

		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New connection from", conn.RemoteAddr())

	for {
		var resp []byte
		for {
			var bit []byte = make([]byte, 1)
			fmt.Println("preread")

			n, err := conn.Read(bit)
			if err != nil {
				fmt.Println("ERR reading bit:", err)
				break
			}

			fmt.Print(n, " ")
			fmt.Print(bit[0] == 10, " ")
			if bit[0] == 10 { // \n
				fmt.Println("newline")
				break
			}
			fmt.Println(string(bit))

			resp = append(resp, bit...)
			if len(resp) > 256 {
				fmt.Println("Hit response break")
				break
			}
		}

		fmt.Println(string(resp))
	}
	///////////////////////////////////////

	// fmt.Println("New connection from", conn.RemoteAddr())
	// for {
	// 	var buffer [1024]byte

	// 	_, err := conn.Read(buffer[:])
	// 	if err != nil {
	// 		log.Printf("err while reading from conn: %v, exiting ...", err)
	// 		return
	// 	}
	// 	fmt.Println("message read: ", string(buffer[:]))
	// }
}

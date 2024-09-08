package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln , err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Servidor escuchando en el puerto 8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	message := "¡Hola, cliente!\n"
	conn.Write([]byte(message))
}
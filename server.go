package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}

func main() {
	listener, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	log.Println("Listening on 0.0.0.0:2000")

	for {
		conn, err := listener.Accept()
		log.Println("Recieved a connection")

		if err != nil {
			log.Fatalln("Unable to accept connection")
		}

		// go keyword for enable concurrent routine
		go echo(conn)
	}
}

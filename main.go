package main

import (
	"io"
	"net"
	"log"
	"fmt"
)

// var templates *template.Template

func main() {
	// Establishing a server and connection over http
	li, err := net.Listen("tcp", ":8080")

	if err!= nil{
		log.Panic(err)
	}

	defer li.Close()
	fmt.Println(li)

	for{
		conn, err := li.Accept()
		if err!= nil{
			log.Println(err)
		}

		io.WriteString(conn, "Hello there!, TCP connection has been establiched!")

		conn.Close()
	}
}

package main

import (
	"flag"
	"fmt"
	"net"
)

func main(){
	port := flag.String("port", "8080", "порт для подключения")
	listener, _ := net.Listen("tcp", ":" + *port)

	defer listener.Close()
	
	for{
	 conn, _ := listener.Accept()
	 go func(c net.Conn) {
			defer c.Close()
			fmt.Fprint(c, "OK\n")
	 }(conn)
	}
}
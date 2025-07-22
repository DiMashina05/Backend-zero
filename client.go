package main

import (
	"bufio"
	"log"
	"net"
)


func main(){
		conn, _ := net.Dial("tcp", "localhost:8080")
		defer conn.Close()
		reader := bufio.NewReader(conn)
	 	line, _ := reader.ReadString('\n')
		if line != "OK\n"{
		log.Print("Подключение возвращает что-то не то(")
		} else {
			log.Print("Корректный ответ сервера")
		}
}

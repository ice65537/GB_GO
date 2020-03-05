package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

var shutDownFlag bool = false

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	listener2, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}
	//
	fmt.Println("Server is running")
	//
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Print(err)
				continue
			}
			go handleConnWrite(conn)
		}
	}()
	//
	go func() {
		for {
			conn, err := listener2.Accept()
			if err != nil {
				log.Println(err)
			}
			go handleConnRead(conn)
		}
	}()
	for !shutDownFlag {
		time.Sleep(time.Second)
	}
	fmt.Println("Server is down")
}
func handleConnWrite(c net.Conn) {
	fmt.Println("handleConnWrite")
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
func handleConnRead(c net.Conn) {
	fmt.Println("handleConnRead")
	defer c.Close()
	reader := bufio.NewReader(c)
	for {
		fmt.Println("Awaiting for client command...")
		s, err := reader.ReadString(byte('\n'))
		if err != nil {
			fmt.Println("Ошибка чтения команды", err)
			return
		}
		fmt.Print("Command received: ", s)
		if strings.ToLower(s) == "exit\n" {
			c.Write([]byte("Start server shutdown...Connection closed\n"))
			shutDownFlag = true
			return
		}
		//Просто эхо команды
		_, err = c.Write([]byte(s + "\n"))
		if err != nil {
			fmt.Println("Ошибка эхо", err)
			return
		}
	}
}

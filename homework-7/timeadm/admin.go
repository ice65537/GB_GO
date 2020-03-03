package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	var request string
	var response string
	conn, err := net.Dial("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//
	for {
		fmt.Print(">")
		_, err := fmt.Scanln(&request)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = conn.Write([]byte(request + "\n"))
		if err != nil {
			fmt.Println(err)
			return
		}
		reader := bufio.NewReader(conn)
		response, err = reader.ReadString(byte('\n'))
		if err != nil {
			fmt.Println(err)
			return
		}
		response = response[0 : len(response)-1]
		fmt.Println(">>>", response)
		if strings.ToLower(request) == "exit" {
			break
		}
	}
}

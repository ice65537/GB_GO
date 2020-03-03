package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() {
		responses <- request("http://lib.ru/")
	}()
	go func() {
		responses <- request("https://vk.com/")
	}()
	go func() {
		responses <- request("https://www.rbc.ru/")
	}()
	return <-responses // возврат самого быстрого ответа
}

func request(hostname string) string {
	html, err := http.Get(hostname)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer html.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(html.Body)
	return buf.String()
}

func main() {
	fmt.Println(mirroredQuery())
}

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const bufferSize int = 65536

func main() {
	f1Name := flag.String("source", "", "Имя файла для копирования")
	f2Name := flag.String("target", "", "Имя копии (будет перезаписан при наличии)")
	needHelp := flag.Bool("help", false, "Вывод справки")
	flag.Parse()

	if *f1Name == "" || *f2Name == "" || *needHelp {
		flag.PrintDefaults()
		return
	}

	file1, err := os.Open(*f1Name)
	if err != nil {
		fmt.Println("Невозможно открыть файл: " + err.Error())
		return
	}
	defer file1.Close()

	file2, err := os.OpenFile(*f2Name, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Невозможно создать файл: " + err.Error())
		return
	}
	defer file2.Close()

	// reading file
	buffer := make([]byte, bufferSize)
	for {
		n, err := file1.Read(buffer)
		//fmt.Println("Прочитано: ", n)
		if err != nil {
			if err == io.EOF {
				//Можно выходить, т.к. io.EOF возвращается при n==0
				break
			} else {
				fmt.Println("Ошибка чтения файла: " + err.Error())
				return
			}
		}
		_, err = file2.Write(buffer[:n-1])
		if err != nil {
			fmt.Println("Ошибка записи файла: " + err.Error())
			return
		}
	}
	fmt.Println("Файл ", *f1Name, "скопирован успешно как ", *f2Name)
}

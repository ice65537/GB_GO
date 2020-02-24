package phonebook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var book map[string][]int
var bookfilename string = ""

//SetFileName -
func SetFileName(filename string) {
	bookfilename = filename
	if bookfilename != "" {
		if !saveToFile(bookfilename) {
			bookfilename = ""
		}
	}
}

//ReInitBook -
func ReInitBook(filename string) {
	bookfilename = filename
	book = nil
	initbook()
}

func initbook() {
	if book == nil {
		book = make(map[string][]int)
	}
	if bookfilename != "" {
		if !loadFromFile(bookfilename) {
			bookfilename = ""
		}
	}
}

func loadFromFile(filename string) bool {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Ошибка чтения файла [", filename, "]: ", err.Error())
		return false
	}
	//
	err = json.Unmarshal(buffer, &book)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON из файла [", filename, "]: ", err.Error())
		return false
	}
	//
	return true
}

//SetItem -
func SetItem(fio string, phones ...int) {
	initbook()
	book[fio] = phones
	if bookfilename != "" {
		if !saveToFile(bookfilename) {
			bookfilename = ""
		}
	}
}

func saveToFile(filename string) bool {
	//
	buffer, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Ошибка формирования JSON: ", err.Error())
		return false
	}
	//
	err = ioutil.WriteFile(filename, buffer, 0644)
	if err != nil {
		fmt.Println("Ошибка записи JSON в файл [", filename, "]: ", err.Error())
		return false
	}
	return true
}

//DelItem -
func DelItem(fio string) {
	if _, ok := book[fio]; !ok {
		panic(fmt.Sprint(fio, " not found!"))
	}
	delete(book, fio)
	if bookfilename != "" {
		if !saveToFile(bookfilename) {
			bookfilename = ""
		}
	}
}

//Print -
func Print(a ...interface{}) {
	initbook()
	fmt.Print(a...)
	fmt.Println(book)
}

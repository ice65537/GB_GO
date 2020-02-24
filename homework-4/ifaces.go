package main

import (
	"GBStudy/homework-4/phonebook"
	"fmt"
	"sort"
)

type v1 struct {
	x string
	y int
	z float32
}
type v2 struct {
	x float32
	y string
	z int
}
type anyVForV1 interface {
	getV1() v1
}

func (v v1) getV1() (vOut v1) {
	vOut = v
	return
}

func (v v2) getV1() (vOut v1) {
	vOut.x = v.y
	vOut.y = v.z
	vOut.z = v.x
	return
}

func sumV1float(x ...anyVForV1) (outF float32) {
	outF = 0
	for _, value := range x {
		outF += value.getV1().z
	}
	return
}

func main() {
	//свои структуры и интерфейс
	a1 := v1{"sdfsdf", 1, 0.25}
	a2 := v2{0.38, "kggkhgjhk", 45}
	a3 := v2{3.44, "sdfsdfsdfsdf/sdf", 65536}
	fmt.Println(sumV1float(a1, a2, a3))

	//работа с телефонным справочником
	var book phonebook.PhoneBook

	/*if err := book.LoadFromFile("c:/ice/phonebook.json"); err != nil {
		fmt.Println(err)
		return
	}*/
	book.SetItem("Коля", "kolya@uk.com", 89002223344, 88005556699, 89161130201)
	book.SetItem("Поликарп Поликарпыч Шариков", "v21@uk.com", 50502)
	book.SetItem("Вася", "vasya@uk.com", 88005559988, 87776665544)
	//
	fmt.Print("Телефонная книга, состояние 1:\n", book, book.IdxName, "\n")
	sort.Sort(&book)
	fmt.Print("Телефонная книга, состояние 1s:\n", book, book.IdxName, "\n")
	book.SaveToFile("c:/ice/phonebook1.json")
	//
	book.DelItem("Поликарп Поликарпыч Шариков")
	book.SetItem("Борис, Лорд Волан-де-Морт", "garry.potter.is.dead@gmail.com", 666666666)
	fmt.Print("Телефонная книга, состояние 2:\n", book, book.IdxName, "\n")
	sort.Sort(&book)
	fmt.Print("Телефонная книга, состояние 2s:\n", book, book.IdxName, "\n")
	book.SaveToFile("c:/ice/phonebook2.json")
}

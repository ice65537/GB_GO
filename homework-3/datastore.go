package main

import (
	"GBStudy/homework-3/moto"
	"GBStudy/homework-3/phonebook"
	"GBStudy/homework-3/queue"
	"fmt"
)

func main() {
	//работа со структурами
	var moto1 moto.Car
	var moto2 *moto.Car
	var moto3 moto.Truck

	moto1 = moto.Car{Chassis: moto.Moto{Brand: "Пежо", Model: "206s", ProdYear: 2006}, ChassisClass: moto.ChassisClassSedan}

	moto2 = new(moto.Car) // moto2 = &moto.Car{}
	moto2.Chassis.Brand = "Пежо"
	moto2.ChassisClass = moto.ChassisClassStationWagon
	moto2.Chassis.Model = "206sw"

	moto3.Chassis.Brand = "Камаз"
	moto3.BootClass = moto.BootClassRefrigerator

	fmt.Println(moto1)
	fmt.Println(moto2)
	fmt.Println(moto3)

	//работа с очередью
	queue.Push("Коля")
	queue.Push("Вася")
	queue.Push("Миша")
	fmt.Println("Обработан: ", queue.Pop())
	fmt.Println("Обработан: ", queue.Pop())
	fmt.Println("Обработан: ", queue.Pop())
	//fmt.Println(queue.Pop()) --panic

	//работа с телефонным справочником
	phonebook.SetItem("Коля", 89002223344, 88005556699, 89161130201)
	phonebook.SetItem("Вася", 88005559988, 87776665544)
	phonebook.SetItem("Поликарп Поликарпыч Шариков", 50502)
	phonebook.Print("Телефонная книга, состояние 1: ")
	phonebook.SetFileName("c:/ice/phonebook1.json")
	//
	phonebook.SetFileName("c:/ice/phonebook2.json")
	phonebook.DelItem("Поликарп Поликарпыч Шариков")
	phonebook.SetItem("Лорд Волан-де-Морт", 666666666)
	phonebook.Print("Телефонная книга, состояние 2: ")
	//
	phonebook.ReInitBook("c:/ice/phonebook1.json")
	phonebook.Print("Телефонная книга, состояние 3: ")

}

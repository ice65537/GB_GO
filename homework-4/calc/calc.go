package main

import (
	"fmt"
	"strconv"
	"strings"
)

var firstNum float32
var secondNum float32
var action string

type consoleInput struct {
	data   string
	isExit bool
	isHelp bool
}

func (x *consoleInput) inputData(msg string) {
	var data string
	fmt.Print(msg)
	fmt.Scanln(&data)
	x.data = data
	x.isExit = false
	x.isHelp = false
	switch strings.ToLower(data) {
	case "exit":
		x.isExit = true
	case "help":
		x.isHelp = true
	}
}

func enterNumber(msg string) float32 {
	var x consoleInput
	for {
		x.inputData(msg)
		num, err := strconv.ParseFloat(x.data, 32)
		if err == nil {
			// Ввод числа прошел без ошибок
			return float32(num)
		}
		panic("Не удалось распознать число")
	}
}

func calculate(firstNum float32, action string, secondNum float32) float32 {
	switch action {
	case "+":
		return firstNum + secondNum
	case "-":
		return firstNum - secondNum
	case "*":
		return firstNum * secondNum
	case "/":
		return firstNum / secondNum
	default:
		panic("Не удалось распознать действие")
	}
}

func printHelp() {
	fmt.Println("Доступные действия:")
	fmt.Println("Сложение: +")
	fmt.Println("Вычитание: -")
	fmt.Println("Умножение: *")
	fmt.Println("Деление: /")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Нештатное завершение программы: ", r)
		}
	}()

	var x consoleInput

	for {
		x.inputData("Программа \"Калькулятор\" (введите help для справки, exit для выхода" +
			" или нажмите Enter для продолжения работы): ")
		if x.isExit {
			break
		} else if x.isHelp {
			printHelp()
			continue
		}
		// Ввод первого числа
		firstNum := enterNumber("Введите первое число: ")

		x.inputData("Укажите действие: ")
		action := x.data

		// Ввод второго числа
		secondNum := enterNumber("Введите второе число: ")

		fmt.Printf("Результат: %v \n", calculate(firstNum, action, secondNum))
	}
}

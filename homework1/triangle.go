package main

import (
	"fmt"
	"math"
)

func main() {
	var sideA float64 = 0 //катет-1
	var sideB float64 = 0 //катет-2
	var sideC float64 = 0 //гипотенуза
	//
	fmt.Print("Введите величину первого катета:")
	fmt.Scanln(&sideA)
	fmt.Print("Введите величину второго катета:")
	fmt.Scanln(&sideB)
	sideC = math.Sqrt(sideA*sideA + sideB*sideB)
	fmt.Printf("Гипотенуза: %.2f\n", sideC)
	fmt.Printf("Периметр: %.2f\n", sideA+sideB+sideC)
	fmt.Printf("Площадь: %.2f\n", sideA*sideB*0.5)
}

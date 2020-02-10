package main

import (
	"fmt"
	"math"
)

func main() {
	var valueRub float64 = 0
	const ratio float64 = 75
	fmt.Print("Введите сумму в рублях:")
	fmt.Scanln(&valueRub)
	fmt.Printf("Сумма в долларах: %.2f\n", math.Round(100*valueRub/ratio)/100)
}

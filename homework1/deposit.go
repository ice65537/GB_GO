package main

import (
	"fmt"
)

func main() {
	var startAmount float64 = 1000 //размер депозита стартовый
	var amount float64 = 0         //размер депозита итоговый
	var interestRate float64 = 12  //процентная ставка в  %%
	var yearCount int = 5          //количество лет
	//
	fmt.Print("Введите размер депозита: ")
	fmt.Scanln(&startAmount)
	fmt.Print("Введите процентную ставку (%): ")
	fmt.Scanln(&interestRate)
	fmt.Print("Введите количество лет: ")
	fmt.Scanln(&yearCount)
	amount = startAmount

	for i := 1; i <= yearCount; i++ {
		amount += amount * interestRate / 100
	}
	fmt.Printf("Cумма депозита %.2f, спустя %d лет, при ставке %.2f процентов составит %.2f\n", startAmount, yearCount, interestRate, amount)
}

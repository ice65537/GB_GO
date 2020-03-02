package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x := 1; x <= 1000; x++ {
			fmt.Printf("naturals <- %d\n", x)
			naturals <- x
		}
		fmt.Println("close(naturals)")
		close(naturals)
	}()

	// возведение в квадрат
	go func() {
		for x, ok := <-naturals; ok; x, ok = <-naturals {
			fmt.Printf("squares <- %d\n", x*x)
			squares <- x * x
		}
		fmt.Println("close(squares)")
		close(squares)
	}()

	// печать
	for x, ok := <-squares; ok; x, ok = <-squares {
		fmt.Println(x)
	}
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go spinner(75*time.Millisecond, 10*time.Second, &wg)
	const n = 42
	//fibN := fibonacci(n)
	wg.Wait()
	//fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration, minSpinning time.Duration, wg *sync.WaitGroup) {
	tStart := time.Now()
	for time.Now().Sub(tStart) < minSpinning {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
	(*wg).Done()
}

func fibonacci(x int) int {
	if x < 2 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

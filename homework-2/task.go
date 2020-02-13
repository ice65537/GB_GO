package main

import (
	"fmt"
	"GBStudy/homework-2/flib"

)

func main() {
	fmt.Println("Тест функций проверки кратности:")
	testMult(234523542345,[]int64{2,3,5,6,7,48,25})
	//
	fmt.Println("Тест функции вычисления чисел Фибоначчи:")
	testFiboBig(100)
}

func testFiboBig(n int) {
	var i int
	for i=0; i<n; i++ {
		fmt.Printf("%v",flib.FiboNumberBigStr(i)); if i<n-1 {fmt.Print(", ")}
	}
}

func testMult (x int64, bases []int64) {
	var result bool
	for _,y := range bases {
		switch y {
		case 2:
			result = flib.IsMultOf2(x)
		case 3:
			result = flib.IsMultOf3(x)
		default:
			result = flib.IsMultOfN(x,y)
		}
		fmt.Printf("%d >> IsMultOf%d >> %v\n",x,y,result)
	}
}

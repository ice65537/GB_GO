package flib

import (
"math/big"
"strconv"

)

var fiboCache map[int]int64 //Кэш int64 фибоначчей
var fiboCacheBig map[int]*big.Int //Кэш big.Int фибоначчей
						
//Тупая регрессия (работает адски медленно)
func FiboNumberSlow (n int) (out int64) {
	if n == 1 || n == 0 {
		out = int64(n)
	} else {
		out = FiboNumberSlow(n-1) + FiboNumberSlow(n-2)
		if out<0 {panic("Переполнение n="+strconv.Itoa(n))}
	}	
	return
}

//Кэшированная регрессия (падает на 93)
func FiboNumber (n int) (out int64) {
	var ok bool
	
	if fiboCache==nil {
		fiboCache = make(map[int]int64)
		fiboCache[0]=0
		fiboCache[1]=1
	}
	
	out, ok = fiboCache[n]
	if !ok {
		out = FiboNumber(n-1) + FiboNumber(n-2)
		if out<0 {panic("Переполнение n="+strconv.Itoa(n))}
		fiboCache[n] = out
	}	
	return
}

//Кэшированная регрессия с big.Int
func FiboNumberBig (n int) (out *big.Int) {
	var ok bool
	
	if fiboCacheBig==nil {
		fiboCacheBig = make(map[int]*big.Int)
		fiboCacheBig[0]=big.NewInt(0)
		fiboCacheBig[1]=big.NewInt(1)
	}

	out, ok = fiboCacheBig[n]
	if !ok {
		out = new(big.Int).Add(FiboNumberBig(n-1),FiboNumberBig(n-2))
		fiboCacheBig[n] = out
	}	
	return
}

//Человечий формат
func FiboNumberBigStr(n int) (out string) {
	x:= FiboNumberBig(n)
	out = x.String()
	return
}


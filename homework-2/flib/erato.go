package flib

import (
	"fmt"
)

var eratoCache []int64               //Кэш c простыми числами
const newNumberCount = 15            //Количество новых чисел, которые будут обрабатываться на каждой итерации
var newNumberCountMultiplier int = 1 //Множитель порции (необходим для случая, когда в порции заданного размера уже нет новых простых чисел)

//PrimeNumberSlice возвращает массив из n простых чисел (2,3...)
func PrimeNumberSlice(n int) (out []int64) {
	var ok bool
	var currentPrime, i, j int64
	var x []int64          //Рабочий массив
	var z map[int64]string //map для хранения вычеркнутых не-простых чисел
	var batchSize int

	if n <= 0 {
		panic("Невозможно сформировать менее 1 простого числа")
	}

	batchSize = newNumberCount * newNumberCountMultiplier //Текущий размер обрабатываемой порции

	if eratoCache == nil {
		//Инициализация
		eratoCache = make([]int64, 1)
		eratoCache[0] = 2
		fmt.Println("Поиск новых простых чисел будет вестить порциями по ", batchSize)
	}

	for len(eratoCache) < n {
		//Надо увеличить количество простых чисел в кэше
		currentPrime = eratoCache[len(eratoCache)-1] //Максимальное известное простое число

		//Формируем временный срез x = известные_простые + порция_новых
		for i = 0; i < int64(batchSize); i++ {
			if i == 0 {
				x = append(eratoCache, currentPrime+1)
			} else {
				x = append(x, i+currentPrime+1)
			}
		}
		//fmt.Println("Временный массив = ", x)
		z = make(map[int64]string, batchSize) //пустой map для "вычеркивания"

		//Основной проход по алгоритму Эратосфена
		for i = 0; i < int64(len(x)); i++ {
			for j = i + 1; j < int64(len(x)); j++ {
				_, ok = z[x[j]]
				if j <= int64(len(eratoCache)-1) || ok { //если число лежит в зоне простых чисел или вычеркнуто, то skip-нуть
					continue
				}
				if x[j]%x[i] == 0 {
					z[x[j]] = "x" //вычеркнуть x[j], если оно кратно предыдущему простому
				}
			}
		}
		//fmt.Println("Итого вычеркнуто = ", z)

		//Сброс невычеркнутых чисел в кэш
		if len(z) == batchSize {
			newNumberCountMultiplier++
			batchSize = newNumberCount * newNumberCountMultiplier
			fmt.Println("Порция обрабатываемых данных увеличена до ", batchSize)
		} else {
			for i = int64(len(eratoCache)); i < int64(len(x)); i++ {
				if _, ok = z[x[i]]; !ok {
					eratoCache = append(eratoCache, x[i])
				}
			}
			//fmt.Println("Кэш = ", eratoCache)
			fmt.Println("Кэш содержит ", len(eratoCache), " простых чисел")
		}
	}
	out = eratoCache[0 : n-1]
	return
}

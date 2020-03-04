package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

const distanceTics int = 1000

type raceResult struct {
	name       string
	finishTime time.Time
}

type raceResults []raceResult

func (x raceResults) Len() int { return len(x) }
func (x raceResults) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
func (x raceResults) Less(i, j int) bool {
	return x[j].finishTime.After(x[i].finishTime)
}

var results raceResults
var wgStartPrepare sync.WaitGroup
var wgStartFlag sync.WaitGroup
var wgFinish sync.WaitGroup
var resultsLock sync.Mutex

func car(name string) {
	var velocity int
	velocity = rand.Intn(100)
	time.Sleep(time.Millisecond * time.Duration(velocity))
	fmt.Println("Автомобиль:", name, "[v =", velocity, "] вышел на старт", time.Now().Format("15:04:05.000000000"))
	wgStartPrepare.Done()
	wgStartFlag.Wait()
	fmt.Println("Автомобиль:", name, "[v =", velocity, "] стартовал в", time.Now().Format("15:04:05.000000000"))
	for i := 0; i < distanceTics; i += velocity {
		time.Sleep(time.Millisecond)
	}
	finishTime := time.Now()
	resultsLock.Lock()
	results = append(results, raceResult{name, finishTime})
	resultsLock.Unlock()
	fmt.Println("Автомобиль:", name, "[v =", velocity, "] финишировал в", finishTime.Format("15:04:05.000000000"))
	wgFinish.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//
	wgStartFlag.Add(1)
	//
	cars := [4]string{"Красный", "Синий", "Желтый", "Фиолетовый"}
	wgStartPrepare.Add(len(cars))
	wgFinish.Add(len(cars))
	//
	results = make([]raceResult, 0, len(cars))
	for i := range cars {
		go car(cars[i])
	}
	//
	wgStartPrepare.Wait() //Ждем когда все машины подготовятся к старту
	fmt.Println("*****Старт*****")
	wgStartFlag.Done() //Старт гонки
	wgFinish.Wait()    //Ждем когда все машины финишируют
	fmt.Println("*****Финиш*****")
	//
	sort.Sort(results)
	for i, v := range results {
		fmt.Println("Место", i+1, "автомобиль", v.name)
	}
}

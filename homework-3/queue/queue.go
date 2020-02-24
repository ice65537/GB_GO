package queue

import "fmt"

var x []string

// Push добавит новый элемент в очередь
func Push(str string) {
	x = append(x, str)
	Print()
}

// Pop вернет первый элемент очереди
func Pop() string {
	// Когда очередь будет пустой, она устроит панику
	if len(x) == 0 {
		panic("Очередь пустая!")
	}
	popElem := x[0]
	x = x[1:]
	Print()
	return popElem
}

//Print todo
func Print() {
	fmt.Println("Queue: ", x)
}

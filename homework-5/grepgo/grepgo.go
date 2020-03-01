package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func main() {
	f1Name := flag.String("source", "", "Имя текстового файла для анализа")
	regExp := flag.String("regexp", "", "Регулярное выражение")
	needHelp := flag.Bool("help", false, "Вывод справки")
	flag.Parse()

	if *f1Name == "" || *regExp == "" || *needHelp {
		flag.PrintDefaults()
		return
	}

	file1, err := os.Open(*f1Name)
	if err != nil {
		fmt.Println("Невозможно открыть файл: " + err.Error())
		return
	}
	defer file1.Close()

	fmt.Println("Построчный анализ файла ", *f1Name, " на соответствие регулярному выражению: ", *regExp)
	scanner := bufio.NewScanner(file1)
	var i int64 = 0
	var foundCount int64 = 0
	for scanner.Scan() {
		s := scanner.Text()
		i++
		if ok, err := regexp.MatchString(*regExp, s); ok {
			fmt.Println("[", i, "]: ", s)
			foundCount++
		} else if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	fmt.Println("Файл ", *f1Name, "проанализирован успешно, найдено соответсвий: ", foundCount)
}

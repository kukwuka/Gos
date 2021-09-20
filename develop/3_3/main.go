package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	sortByRow := flag.Int("k", -1, "Указание колонки для сортировки")
	sortByNumbers := flag.Bool("n", false, "Сортировать по числовому значению")
	sortReverse := flag.Bool("r", false, "Сортировать в обратном порядке")
	dontShowRepeat := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	flag.Parse()

	b, err := ioutil.ReadFile("./text.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'

	strs := strings.Split(str, "/n")

	fmt.Println(strs[0]) // print the content as a 'string'
}

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	var after int
	flag.IntVar(&after, "A", 0, "печатать +N строк после совпадения")
	var before int
	flag.IntVar(&before, "B", 0, "печатать +N строк до совпадения")
	var contextText int
	flag.IntVar(&contextText, "C", 0, "печатать ±N строк вокруг совпадения")
	var countBool bool
	var count int
	flag.BoolVar(&countBool, "c", false, "количество строк")
	var ignoreCase bool
	flag.BoolVar(&ignoreCase, "i", false, "игнорировать регистр")
	var invert bool
	flag.BoolVar(&invert, "v", false, "вместо совпадения, исключать")
	var fixed bool
	flag.BoolVar(&fixed, "F", false, "точное совпадение со строкой, не паттерн")
	var lineNum bool
	flag.BoolVar(&lineNum, "n", false, "печатать номер строки")
	flag.Parse()

	str := os.Args[len(os.Args)-1]

	t, err := ioutil.ReadFile("./text.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	text := string(t)

	fmt.Println(text)
	fmt.Println("=============")

	if lineNum {
		fmt.Println("Find on", LineNum(str, text)+1, "line")
	}
	if countBool {
		count = Count(text)
		fmt.Println("Count of lines", count)
	}
	if after > 0 {
		text = After(str, text, after)
	}
	if before > 0 {
		text = Before(str, text, before)
	}
	if contextText > 0 {
		text = ContextText(str, text, contextText)
	}
	if ignoreCase {
		text = IgnoreCase(str, text)
	}
	if invert {
		text = Invert(str, text)
	}
	if fixed {
		fmt.Println(Fixed(str, text))
	}

	//fmt.Println(text)

}

// After возвращает искомую строку и +count строк после нее
func After(str, file string, count int) string {
	validID := regexp.MustCompile(str)
	var res string
	rows := strings.Split(file, "\n")
	for i, v := range rows {
		if validID.MatchString(v) {
			res += v + "\n"
			for j, k := range rows {
				if j > i && j <= i+count {
					if j == len(rows)-1 || j == i+count {
						res += k
					} else {
						res += k + "\n"
					}
				}
			}
		}
	}
	return res
}

// Before возвращает искомую строку и +count строк до нее
func Before(str, file string, count int) string {
	validID := regexp.MustCompile(str)
	var res string
	rows := strings.Split(file, "\n")
	for i, v := range rows {
		if validID.MatchString(v) {
			for j, k := range rows {
				if j < i && j >= i-count {
					res += k + "\n"
				}
			}
			res += v
		}
	}
	return res
}

// ContextText возвращает искому строку и +-count строк вокруг нее
func ContextText(str, file string, count int) string {
	validID := regexp.MustCompile(str)
	var res string
	rows := strings.Split(file, "\n")
	for i, v := range rows {
		if validID.MatchString(v) {
			for j, k := range rows {
				if j < i && j >= i-count {
					res += k + "\n"
				}
			}
			res += v + "\n"
			for j, k := range rows {
				if j > i && j <= i+count {
					if j == len(rows)-1 || j == i+count {
						res += k
					} else {
						res += k + "\n"
					}
				}
			}
		}
	}
	return res
}

// Count подсчитывает количество строк
func Count(file string) int {
	var count int
	rows := strings.Split(file, "\n")
	for range rows {
		count++
	}
	return count
}

// IgnoreCase возвращает искомую строку игнорируя регистр
func IgnoreCase(str, file string) string {
	var res string
	str = strings.ToLower(str)
	file = strings.ToLower(file)
	validID := regexp.MustCompile(str)
	rows := strings.Split(file, "\n")
	for _, v := range rows {
		if validID.MatchString(v) {
			res += v
			break
		}
	}
	return res
}

// Invert вощвращает текст без искомой строки
func Invert(str, file string) string {
	var res string
	validID := regexp.MustCompile(str)
	rows := strings.Split(file, "\n")
	for i, v := range rows {
		if validID.MatchString(v) {
			continue
		}
		if i == len(rows)-1 {
			res += v
		} else {
			res += v + "\n"
		}
	}
	return res
}

// Fixed возвращает строку, если имеется точное совпадение
func Fixed(str, file string) bool {
	rows := strings.Split(file, "\n")
	for _, v := range rows {
		if v == str {
			return true
		}
	}
	return false
}

// LineNum возвращает номер исходной строки
func LineNum(str, file string) int {
	validID := regexp.MustCompile(str)
	rows := strings.Split(file, "\n")
	for i, v := range rows {
		if validID.MatchString(v) {
			return i
		}
	}
	return -1
}

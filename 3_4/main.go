package main

import (
	"log"
	"sort"
	"strings"
)

func main() {

	ListOfWords := []string{
		"пятак",
		"пятка",
		"тяпка",
		"листок",
		"слиток",
		"столик",
		"хрень",
	}
	log.Print(ListOfWords)

	log.Print(GetAnnograms(ListOfWords))

}

func GetAnnograms(dummyDict []string) map[string][]string {

	list := make(map[string][]string)
	for _, word := range dummyDict {
		key := sortStr(word)
		list[key] = append(list[key], word)
	}
	result := make(map[string][]string)
	for _, value := range list {
		if len(value) < 2 {
			continue
		}
		result[value[0]] = value
	}

	return result
}

func sortStr(k string) string {
	s := strings.Split(k, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

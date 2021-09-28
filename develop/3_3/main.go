package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sortByRow := flag.Int("k", 0, "Указание колонки для сортировки")
	sortByNumbers := flag.Bool("n", false, "Сортировать по числовому значению")
	sortReverse := flag.Bool("r", false, "Сортировать в обратном порядке")
	dontShowRepeat := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	colSeparator := flag.String("t", " ", "таб сепаратор")
	spaceIgnore := flag.Bool("b", false, "игнорировать хвостовые пробелы")
	sorted := flag.Bool("c", false, "проверять отсортированы ли данные")

	flag.Parse()

	file, err := ioutil.ReadFile("text.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	result := Sort(file,
		*sortByRow,
		*colSeparator,
		*sortByNumbers,
		*sortReverse,
		*spaceIgnore,
		*dontShowRepeat,
		*sorted)

	for i := 0; i < len(result); i++ {

		fmt.Println(result[i])
	}
}

func Sort(
	content []byte,
	sortByRow int,
	colSeparator string,
	sortByNumbers, sortReverse, spaceIgnore, dontShowRepeat, sorted bool,
) []string {
	if spaceIgnore {
		content = []byte(strings.TrimSpace(string(content)))
	}
	rez := strings.Split(string(content), "\n")

	if sortByRow > 0 {
		buf := make([]string, 0, len(rez))
		for _, r := range rez {
			buf = append(buf, strings.Split(r, colSeparator)[sortByRow-1])
		}
		rez = buf
	}

	if dontShowRepeat {
		un := map[string]struct{}{}
		for _, r := range rez {
			un[r] = struct{}{}
		}
		rez = make([]string, 0, len(un))
		for k := range un {
			rez = append(rez, k)
		}
	}
	if sortByNumbers {
		nums := make([]int, 0, len(rez))
		if sorted {
			if !sort.IntsAreSorted(nums) {
				return []string{"Need to sort"}
			} else {
				return []string{"No need to sort"}
			}
		}
		for k := range rez {
			n, err := strconv.Atoi(rez[k])
			if err != nil {
				fmt.Println("error sort numeric:", err)
				os.Exit(1)
			}
			nums = append(nums, n)
		}
		if sortReverse {
			sort.Sort(sort.Reverse(sort.IntSlice(nums)))
		} else {
			sort.Ints(nums)
		}

		rez = make([]string, 0, len(nums))
		for _, n := range nums {

			rez = append(rez, strconv.Itoa(n))
		}
	}

	if sortByRow > 0 {
		buf := make(map[string]string, len(rez))
		for _, c := range strings.Split(string(content), "\n") {

			line := strings.Split(c, colSeparator)
			buf[line[sortByRow-1]] = c
		}

		for i := 0; i < len(rez); i++ {
			rez[i] = buf[rez[i]]
		}

	}

	return rez
}

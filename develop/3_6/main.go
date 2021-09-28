package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	fields    string
	delim     = "\t"
	separated bool
)

func main() {
	flag.StringVar(&fields, "f", "-1", "выбрать поля (колонки)")
	flag.StringVar(&delim, "d", "	", "использовать другой разделитель")
	flag.BoolVar(&separated, "s", false, "только строки с разделителем")

	flag.Parse()
	data := enter()
	if false {
		data = parse([]string{})
	}

	rez := cut(data)
	for _, r := range rez {
		fmt.Print(r)
	}

}

func enter() (data [][]string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		line, _ := reader.ReadString('\n')

		if line == "\n" {
			break
		}
		data = append(data, strings.Split(line[:len(line)-1], delim))
	}
	return
}
func parse(lines []string) (data [][]string) {
	for _, line := range lines {
		data = append(data, strings.Split(line[:len(line)-1], delim))
	}
	return
}

func cut(data [][]string) []string {
	rez := []string{}
	f := strings.Split(fields, ",")
	coloumn := []int{}
	for _, c := range f {
		num, err := strconv.Atoi(c)
		if err != nil {
			return nil
		}
		coloumn = append(coloumn, num-1)
	}
	shift := 0
	for i := range data {
		if !separated || len(data[i]) > 1 {
			rez = append(rez, "")
			for j := range data[i] {
				for _, col := range coloumn {
					if j == col || col == -1 {
						rez[i-shift] += data[i][j] + delim
					}
				}
			}
		} else {
			shift++
		}
	}
	return rez
}

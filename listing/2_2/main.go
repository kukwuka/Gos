package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	str := "a4bc2d5e"
	//fmt.Println(str)
	//str, err := unzipStr(str)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(str)
	unzipStr2(str)
}

func unzipStr(str string) (string, error) {
	if str == "" {
		return "", fmt.Errorf("incorrect string")
	}

	strRune := []rune(str)
	var res string

	for i := 0; i < len(strRune); i++ {
		if i < len(strRune)-1 {
			if unicode.IsDigit(strRune[i]) {
				return "", fmt.Errorf("incorrect string")
			}
			if strRune[i] == []rune("\\")[0] {
				i++
			}
			if i < len(strRune)-1 && unicode.IsDigit(strRune[i+1]) {
				count, _ := strconv.Atoi(string(strRune[i+1]))
				for k := 0; k < count; k++ {
					res += string(strRune[i])
				}
				i++
			} else {
				res += string(strRune[i])
			}
		} else {
			res += string(strRune[len(strRune)-1])
		}
	}
	return res, nil
}

func unzipStr2(str string) (string, error) {
	letters := []rune(str)
	if unicode.IsDigit(letters[0]) {
		return "", fmt.Errorf("incorrect string")
	}

	// var res []rune
	for i := 0; i < len(letters)-1; i++ {
		if letters[i] == []rune("\\")[0] {

		}
	}
	return "", nil
}

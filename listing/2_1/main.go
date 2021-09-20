package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b)
	//выдаст [77 78 79] с 1 до 4го, 4 не включает
}

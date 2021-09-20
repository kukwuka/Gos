package main

import (
	"fmt"
	"os"
)

type MyError struct {
}

func (me *MyError) Error() string {
	return "My error"
}

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
	fmt.Print(err.Error())
	// var err error
	// fmt.Println(err)
	// fmt.Println(err == nil)

	// var my *MyError = nil
	// fmt.Println(my)

	// err = my
	// fmt.Println(err == nil)

	// err = Foo()
	// fmt.Println(err.Error())

	// fmt.Println(my == nil)
}

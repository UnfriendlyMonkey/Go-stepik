package main

import (
	"fmt"
	"errors"
)

func divide(a, b int) int {
	return a / b
}

func ErrorExample() {
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Check input data types")
	} else {
		fmt.Println(divide(input, 5))
	}
}

func OwnError() {
	err := errors.New("my error")
	fmt.Println("", err)
}

func divideWithErr(a, b int) (int, error) {
	res := a / b
	err := errors.New("just error")
	fmt.Printf("%T\n", err)
	return res, err
}

func CatchError() {
	var a, b int
	fmt.Scan(&a, &b)
	res, err := divideWithErr(a, b)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(res)
	}
}

func specPrint(s uint32) uint32 {
	// fmt.Print(s)
	fmt.Printf("Printing from deferred: %v\n", s)
	s++
	return s
}

func testDefer() {
	var a uint32 = 5
	defer specPrint(a)
	a = 6
	// fmt.Print(a)
	fmt.Printf("Printing from test: %v\n", a)
}

func main() {
	// ErrorExample()
	// OwnError()
	// CatchError()
	testDefer()
}

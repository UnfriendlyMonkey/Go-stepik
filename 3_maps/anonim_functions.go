package main

import (
	"fmt"
	"unicode"
	"strings"
)

func invertCase(r rune) rune {
	if unicode.IsLower(r) {
		return unicode.ToUpper(r)
	}
	return unicode.ToLower(r)
}

func ExampleFirstClassFunctionArgument() {
	str := "SoMe StRiNg"
	test := "sOmE sTrInG"

	str = strings.Map(invertCase, str)
	fmt.Printf("inverted string: %s. Asserted: %v.\n", str, str == test)
}

func ExampleFirstClassReturn() func(rune) rune {
	return invertCase
}

func externalFunction() func() {
	text := "TEXT"
	fn := func() {
		fmt.Println("text: ", text)
	}
	text = "NEW"
	return fn
}

func ExampleEnvironment() {
	fn := externalFunction()
	text := "TRY THIS"
	fn()
	fmt.Println(text)
}

func ExampleClosure() {
	fn := func() func(int) int {
		count := 0
		return func(i int) int {
			// closure can not only use but change its values as well
			count++
			return i * count
		}
	}()
	fmt.Printf("%T\n", fn) // func(int) int

	for i := 1; i <= 5; i++ {
		fmt.Println(fn(i))
	}
}

func onlyEvens() {

	delNotEvens := func(n uint) uint {
		var sum uint = 0
		addNext := func(b uint) {
			sum = sum * 10 + b
		}

		func() {
			for n > 0 {
				if d := n % 10; d != 0 && d % 2 == 0 {
					defer addNext(d)
					// defer works only after function it is declared in
				}
				n /= 10
			}
		}()

		if sum == 0 {return 100}
		return sum
	}

	fn := func(n uint) (res uint) {
		for j := uint(1); n > 0; n /= 10 {
			if d := n % 10; d != 0 && d&1 == 0 {
				res += j * d
				// declared in signature
				j *= 10
			}
		}
		if res == 0 {res = 100}
		return
		// var to return is declared in signature
	}

	fmt.Println(delNotEvens(727178))
	fmt.Println(fn(727178))
}

func main() {
	// ExampleFirstClassFunctionArgument()
	// ExampleEnvironment()
	// ExampleClosure()
	onlyEvens()
}

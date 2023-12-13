package main

import (
	"fmt"
	"math"
	"strings"
)

func findHypot() {
	var a, b float64
	fmt.Scan(&a, &b)
	res := math.Sqrt(a * a + b * b)
	// return res
	fmt.Println(res)
}

func wildcardedString() {
	var s string
	fmt.Scan(&s)
	// fmt.Println(strings.ReplaceAll(s, "", "*"))
	rs := strings.Split(s, "")
	// fmt.Println(rs)
	res := strings.Join(rs, "*")
	fmt.Println(res)
}

func largestDigit() {
	var s string
	fmt.Scan(&s)
	// fmt.Printf("%v: %T, %T\n", s[0], fmt.Sprintf("%c", s[0]), s[0])
	var m uint8
	for i := 0; i < len(s); i++ {
		if s[i] >= m {
			m = s[i]
		}
	}
	// fmt.Println(m)
	fmt.Printf("%c\n", m)
}

func SquaredDigits() {
	// var s string
	// fmt.Scan(&s)
	// for i := 0; i < len(s); i++ {
	// 	fmt.Print((s[i] - '0') * (s[i] - '0'))
	// }
	// var_2 via int and deferred
	var n int
	// fmt.Scan(&n)
	for fmt.Scan(&n); n != 0; n /= 10 {
		defer fmt.Print((n % 10) * (n % 10))
	}
}

const k float64 = 1296
const p float64 = 6
const v float64 = 6

func M() float64 {
	return p * v
}

func W() float64 {
	m := M()
	return math.Sqrt(k / m)
}

func T() float64 {
	w := W()
	return 6 / w
}

func wtfPendulum() {
	fmt.Println(T())
}

func main() {
	// findHypot()
	// wildcardedString()
	// largestDigit()
	// SquaredDigits()
	wtfPendulum()
}

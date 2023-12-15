package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"strconv"
	"unicode"
)

func clearString(s string) string {
	var r string
	rs := []rune(s)
	for _, el := range rs {
		if unicode.IsDigit(el) {
			r += string(el)
		}
	}
	return r
}

func sumDirtyStrings() int64 {
	var a, b string
	fmt.Scan(&a, &b)
	a = clearString(a)
	b = clearString(b)
	fmt.Println(a, b)
	ai, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		panic(err)
	}
	bi, err := strconv.ParseInt(b, 10, 64)	
	if err != nil {
		panic(err)
	}
	return ai + bi
}

func sumFromCSV() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	// fmt.Println(input, err)
	input = strings.TrimRight(input, "\n")
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, ",", ".")
	
	as := strings.Split(input, ";")
	i1, err := strconv.ParseFloat(as[0], 64)
	if err != nil {
		panic(err)
	}
	i2, err := strconv.ParseFloat(as[1], 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%.4f\n", i1 / i2)
}

func main() {
	// fmt.Println(sumDirtyStrings())
	sumFromCSV()
}

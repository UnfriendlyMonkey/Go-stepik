package main

import (
	// "errors"
	"fmt"
	"strings"
)

// func checkFloat64(val interface{}) (float64, error) {
// 	switch t := val.(type) {
// 	case float64: return t, nil
// 	default: return float64(0), errors.New("not float64")
// 	}
// }

func makeResult(val1, val2 float64, op string) float64 {
	switch op {
	case "+": return val1 + val2
	case "-": return val1 - val2
	case "*": return val1 * val2
	default: return val1 / val2
	}
}

func calc(v1, v2, op interface{}) {
	val1, ok := v1.(float64)
	if !ok {fmt.Printf("value=%v: %T\n", v1, v1); return}

	val2, ok := v2.(float64)
	if !ok {fmt.Printf("value=%v: %T\n", v2, v2); return}

	operand, ok := op.(string)
	if !ok || !strings.Contains("+-/*", operand) {
		fmt.Println("неизвестная операция")
		return
	}

	fmt.Printf("%.4f\n", makeResult(val1, val2, operand))
}

type Battery struct {
	report string
}

func (b Battery) String() string {
	full := strings.Count(b.report, "1")
	empty := strings.Count(b.report, "0")
	res := "["
	res += strings.Repeat(" ", empty)
	res += strings.Repeat("X", full)
	// for i := 0; i < empty; i++ {res += " "}
	// for j := 0; j < full; j++ {res += "X"}
	res += "]"
	return res
}

func main() {
	// var value1, value2, operation interface{} = 10.5, 23.4, "+"
	// value1, value2, operation := float64(10), bool(true), "o"
	// value1, value2, operation := float64(10), float64(23.4), false
	// calc(value1, value2, operation)
	var report string
	fmt.Scan(&report)
	batteryForTest := Battery{report: report,}
	fmt.Println(batteryForTest)
}

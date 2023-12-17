package main

import (
	// "errors"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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

type Book struct {
	Title string
	Author string
	Year uint16
}

func (b Book) String() string {
	return fmt.Sprintf("Book: %s - %s; %d", b.Title, b.Author, b.Year)
}

type Count int

func (c Count) String() string {
	return strconv.Itoa(int(c))
}

func WriteLog(s fmt.Stringer) {
	log.Print(s.String())
}

func fromAlexedwards() {
	book := Book{"HeadFirst Go", "Jay McGavren", 2019}
	WriteLog(book)

	count := Count(3)
	WriteLog(count)
}

type Customer struct {
	Name string
	Age int
}

func (c *Customer) WriteJSON(w io.Writer) error {
	js, err := json.Marshal(c)
	if err != nil {return err}
	_, err = w.Write(js)
	return err
}

func fromAlexedwards2() {
	c := &Customer{Name: "Alice", Age: 21}

	// We can call the WriteJSON method using a buffer...
	var buf bytes.Buffer
	err := c.WriteJSON(&buf)
	if err != nil {log.Fatal(err)}

	// or using a file
	f, err := os.Create("/tmp/customer")
	if err != nil {log.Fatal(err)}
	defer f.Close()
	err = c.WriteJSON(f)
	if err != nil {log.Fatal(err)}
}

func emptyInterfaceExample() {
	person := make(map[string]interface{}, 0)

	person["name"] = "Alice"
	person["age"] = 21
	person["height"] = 167.4

	log.Printf("%+v", person)

	// person["age"] = person["age"] + 1 // misamtched types interface{} and int

	age, ok := person["age"].(int)
	if !ok {log.Fatal("could not assert value to int"); return}
	person["age"] = age + 1
	log.Printf("%+v", person)
}

func main() {
	// var value1, value2, operation interface{} = 10.5, 23.4, "+"
	// value1, value2, operation := float64(10), bool(true), "o"
	// value1, value2, operation := float64(10), float64(23.4), false
	// calc(value1, value2, operation)

	// var report string
	// fmt.Scan(&report)
	// batteryForTest := Battery{report: report,}
	// fmt.Println(batteryForTest)

	// fromAlexedwards()
	// fromAlexedwards2()
	emptyInterfaceExample()
}

package main

import "fmt"

func string_byte() {
  fmt.Println("Hello"[0])
  fmt.Println(string("Hello"[0]))
}

func hello_go() {
  fmt.Println("Hello, Go!")
}

func variables() {
  var hello string
  hello = "Hello Go!"

  var a int = 2023

  simple_string := "As simple as that"
  var simpe_int = 123

  fmt.Println(hello)
  fmt.Println(a)

  fmt.Println(simple_string, simpe_int)
}

func symbols() {
  fmt.Println("func_symbols")
  var (
    symbol int32 = 'c'
    s_2 rune = 'y'
  )
  fmt.Println(symbol, s_2)
  fmt.Println(string(symbol), string(s_2))
}

func int_as_string() {
  fmt.Println("int as strings")
  var ex int = 'r'
  // var longex int = 'example'
  // illegal rune literal
  var ex_2 int = 'e'
  fmt.Println(ex, string(ex))
  fmt.Println(ex_2, string(ex_2))
}

func math_operations() {
  var t int = 93
  a := 100
  i := a - t
  var j = a - t
  // var f = 100.0
  var k float32 = 100.0 / 10
  var m float32 = 100 / 10.0
  fmt.Println(i, j, k, m)
}

func scan_input() {
  var name string
  var age int
  fmt.Println("Input your name: ")
  fmt.Scan(&name)
  fmt.Println("Enter you age: ")
  fmt.Scan(&age)

  fmt.Println(name, age)
}

func simple_math() {
  var a int
  fmt.Scan(&a)
  b := a * 2
  c := b + 100
  fmt.Println(c)
}

func main() {
  // hello_go()
  // string_byte()
  // variables()
  // symbols()
  // int_as_string()
  // math_operations()
  scan_input()
  // simple_math()
}

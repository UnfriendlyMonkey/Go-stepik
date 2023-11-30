package main

import "fmt"

func switch_case() {
  v := 42
  // v := 100
  switch v {
  case 100:
    fmt.Println(100)
    // fallthrough
  case 42:
    fmt.Println(42)
    fallthrough
  case 1:
    fmt.Println(1)
    fallthrough
  default: fmt.Println("default")
  }
}

func s_w_2() {
  var c uint32
  fmt.Scan(&c)
  switch {  // may not point what variable to chek in this block
  case 1 <= c && c <= 9:
    fmt.Println("from 1 to 9")
  case 100 <= c && c <= 250:
    fmt.Println("from 100 to 250")
  default: fmt.Println("don't know")
}
}

func is_positive() {
  var a int
  fmt.Scan(&a)
  switch {
  case a > 0:
    fmt.Println("Число положительное")
  case a < 0:
    fmt.Println("Число отрицательное")
  default:
    fmt.Println("Ноль")
  }
}

func is_digits_different() {
  var a int
  fmt.Scan(&a)
  c := a / 100
  d := a % 100
  d = d / 10
  i := a % 10
  fmt.Println(c, d, i)
  if i == d || d == c || i == c {
    fmt.Println("NO")
  } else {
    fmt.Println("YES")
  }
}

func first_digit() {
  var a int
  fmt.Scan(&a)
  switch {
  case a == 10000: fmt.Println(1)
  case a > 999:
    fmt.Println(a / 1000)
  case a > 99:
    fmt.Println(a / 100)
  case a > 9:
    fmt.Println(a / 10)
  default:
    fmt.Println(a)
  }
}

func is_lucky() {
  var a int
  fmt.Scan(&a)
  first := a / 100000
  second := a % 100000 / 10000
  third := a % 10000 / 1000
  fourth := a % 1000 / 100
  fifth := a % 100 / 10
  last := a % 10
  fmt.Println(first, second, third, fourth, fifth, last)
  first_three := first + second + third
  last_three := fourth + fifth + last
  if first_three == last_three {
    fmt.Println("YES")
  } else {
    fmt.Println("NO")
  }
}

func is_leap() {
  var year int
  fmt.Scan(&year)
  var message string
  if year % 400 == 0 {
    message = "YES"
  } else if year % 4 == 0 && year % 100 != 0 {
    message = "YES"
  } else {
    message = "NO"
  }
  fmt.Println(message)
}

func main() {
  // switch_case()
  // s_w_2()
  // is_positive()
  // is_digits_different()
  // first_digit()
  // is_lucky()
  is_leap()
}

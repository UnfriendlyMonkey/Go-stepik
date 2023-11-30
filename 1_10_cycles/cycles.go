package main

import "fmt"

func first_cycle() {
  sum := 0
  for i := 1; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)
}

func infinite_scan() {
  var n int
  for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
    fmt.Println(n)
  }
}

func squares() {
  for i := 1; i <= 10; i++ {
    fmt.Println(i * i)
  }
}

func sums_a_b() {
  sum := 0
  var a, b int
  fmt.Scan(&a)
  fmt.Scan(&b)
  for ; a <= b; a++ {
    sum += a
  }
  fmt.Println(sum)
}

func divisible_by_8() {
  sum := 0
  var a, b int
  fmt.Scan(&a)
  for i := 1; i <= a; i++ {
    fmt.Scan(&b)
    if b < 100 && b > 9 && b % 8 == 0 {
      sum += b
    }
  }
  fmt.Println(sum)
}

func max_counter() {
  var n, max, count int
  max = 0
  count = 0
  for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
    if n > max {
      max = n
      count = 1
    } else if n == max {
      count += 1
    }
  }
  fmt.Println(count)
}

func n_c_d() {
  var n, c, d int
  fmt.Scan(&n, &c, &d)
  for i := 1; i <= n; i++ {
    if i % d == 0 {
      continue
    }
    if i % c == 0 {
      fmt.Println(i)
      break
    }
  }
}

func ten_hundred() {
  var n int
  for {
    fmt.Scan(&n)
    if n < 10 {
      continue
    }
    if n > 100 {
      break
    }
    fmt.Println(n)
  }
}

func bank_account() {
  var init float32
  var interest float32
  var goal int
  fmt.Scan(&init, &interest, &goal)
  for i := 1; ; i++ {
    fmt.Println(i, ": ", init)
    per_interest := interest / 100
    grow := per_interest * init
    init = init + grow
    fmt.Println(init)
    // init := int(init)
    fmt.Println(init)
    if int(init) >= goal {
      fmt.Println(i)
      break
    }
  }
}

func same_digits() {
  var a, b string
  fmt.Scan(&a, &b)
  // fmt.Println(len(a), string(a[0]))
  for i := 0; i < len(a); i++ {
    a_letter := string(a[i])
    for j := 0; j < len(b); j++ {
      b_letter := string(b[j])
      if a_letter == b_letter {
        fmt.Print(a_letter, " ")
        break
      }
    }
  }
}

func formatted_out() {
  var a float64
  fmt.Scan(&a)
  switch {
  case a <= 0:
    fmt.Printf("число %2.2f не подходит", a)
  case a > 10000:
    fmt.Printf("%e", a)
  default:
    r := a * a
    fmt.Printf("%.4f", r)
  }
}

func main() {
  // first_cycle()
  // infinite_scan()
  // squares()
  // sums_a_b()
  // divisible_by_8()
  // max_counter()
  // n_c_d()
  // ten_hundred()
  // bank_account()
  // same_digits()
  formatted_out()
}

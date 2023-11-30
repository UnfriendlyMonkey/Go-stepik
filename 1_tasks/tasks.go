package main

import "fmt"

func sum_of_digits() {
  var n int
  var s int
  fmt.Scan(&n)
  for n > 0 {
    s += n % 10
    n = n / 10
    // fmt.Println(n, s)
  }
  fmt.Println(s)
}

func revers_num() {
  var n string
  fmt.Scan(&n)
  fmt.Printf("%c%c%c", n[2], n[1], n[0])
}

func secs_to_time() {
  var secs int
  fmt.Scan(&secs)
  secsInHour := 3600
  secsInMin := 60
  hours := secs / secsInHour
  secsLeft := secs % secsInHour
  mins := secsLeft / secsInMin
  fmt.Printf("It is %d hours %d minutes.", hours, mins)
}

func is_triangle_rectangular() {
  var a, b, c int
  fmt.Scan(&a, &b, &c)
  isRectangular := a * a + b * b == c * c
  // fmt.Println(isRectangular)
  switch isRectangular {
  case true: fmt.Println("Прямоугольный")
  case false: fmt.Println("Непрямоугольный")
  }
}

func is_triangle_possible() {
  var a, b, c int
  fmt.Scan(&a, &b, &c)
  impossible := a + b < c || a + c < b || b + c < a
  switch impossible {
  case true: fmt.Println("Не существует")
  case false: fmt.Println("Существует")
  }
}

func half_sum() {
  var a, b int
  fmt.Scan(&a, &b)
  c := a + b
  divider := 2.0
  res := float64(c) / divider
  fmt.Println(res)
}

func count_zeros() {
  var n, sum, el int
  fmt.Scan(&n)
  for i := 0; i < n; i++ {
    fmt.Scan(&el)
    if el == 0 {
      sum++
    }
  }
  fmt.Println(sum)
}

func count_mins() {
  var n, sum, min, el int
  fmt.Scan(&n)
  for i := 0; i < n; i++ {
    fmt.Scan(&el)
    if i == 0 {
      min = el
      sum++
    } else if el == min {
      sum++
    } else if el < min {
      min = el
      sum = 1
    }
  }
  fmt.Println(sum)
}

func digital_root() {
  var n, sum uint
  fmt.Scan(&n)
  for n > 9 {
    sum = 0
    for n > 0 {
      sum += n % 10
      n /= 10
    }
    n = sum
  }
  fmt.Println(n)
}

func largest_by_7() {
  var a, b, res int
  fmt.Scan(&a, &b)
  for i := b; i >= a - 1; i-- {
    if i % 7 == 0 {
      res = i
      break
    }
  }
  if res < a {
    fmt.Println("NO")
  } else {
    fmt.Println(res)
  }
}

func main() {
  // sum_of_digits()
  // revers_num()
  // secs_to_time()
  // is_triangle_rectangular()
  // is_triangle_possible()
  // half_sum()
  // count_zeros()
  // count_mins()
  // digital_root()
  largest_by_7()
}

package main

import "fmt"

func main() {
  // var age, name = add(4, 5, "Tom", "Simpson")
  // fmt.Println(age)
  // fmt.Println(name)
  // fmt.Println(minimumFromFour())
  // fmt.Println(vote(1, 0, 1))
  // fmt.Println(fibonacci(10))
  // pack_unpack()
  fmt.Println(sumInt(1, 0, 3, 2345696))
}


func add(x, y int, firstName, lastName string) (int, string) {
  var r int = x + y
  fullName := firstName + " " + lastName
  return r, fullName
}

func minimumFromFour() int {
  var workArray [4]int = [4]int{}
  for i := range workArray {
    fmt.Scan(&workArray[i])
  }
  m := workArray[0]
  for _, el := range workArray {
    if el < m {
      m = el
    }
  }
  return m
}

func vote(x int, y int, z int) int {
  s := x + y + z
  switch {
  case s > 1:
    return 1
  default:
    return 0
  }
}

func fibonacci(n int) int {
  a, b, i := 0, 1, 0
  for ; i < n; i++ {
    a, b = b, a + b
    // fmt.Println(i, ": ", a, b)
  }
  return a
}

func pack_unpack() {
  s1 := []int{1, 2, 3, 4, 5}
  s2 := []int{6, 7, 8, 9, 10}

  s3 := append(s1, s2...)
  fmt.Println(s3)
}

func sumInt(a ...int) (int, int) {
  sum := 0
  for _, el := range a {
    sum += el
  }
  return len(a), sum
}

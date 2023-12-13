package main

import (
	"fmt"
	"slices"
)

func first_arrays() {
  var numbers [5]int = [5]int{1, 2, 3, 4, 5}

  fmt.Println(numbers)
  numbers[0] = 89
  fmt.Println(numbers[0])

  // for i := 0; i < len(numbers); i++ {
  //   fmt.Println(numbers[i])
  // }

  for idx, elem := range numbers {
    fmt.Printf("Elem with index %d: %d\n", idx, elem)
  }
}

func two_vars() {
  a, b := 1, 2
  fmt.Printf("A = %d; B = %d\n", a, b)
  fmt.Printf("%#v, %T\n", a, a)
  fmt.Printf("%c", 12345)
}

func arr_from_input() {
  workArray := [10]uint8{}
  for i := 0; i < len(workArray); i++ {
    fmt.Scan(&workArray[i])
  }
  // fmt.Println(workArray)
  var a, b uint8
  for j := 0; j < 3; j++ {
    fmt.Scan(&a, &b)
    // fmt.Println(a, b)
    // fmt.Println(workArray[a], workArray[b])
    workArray[a], workArray[b] = workArray[b], workArray[a]
    // fmt.Println(workArray[a], workArray[b])
  }
  for _, el := range workArray {
    fmt.Printf("%v ", el)
  }
}

func base_slices() {
  baseArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
  fmt.Printf("Base Array: %v\n", baseArray)

  baseSlice := baseArray[5:8]
  fmt.Printf(
    "Slice based upon Base Array with len %d and capacity %d: %v\n",
    len(baseSlice),
    cap(baseSlice),
    baseSlice,
  )
  pointer := fmt.Sprintf("%p", baseSlice)
  fmt.Println(pointer)

  baseSlice = append(baseSlice, 10)
  // slice doesn't extend array's len, so it stays in the same array
  // by adding new element to slice, we changed the relevant element of array (10 instead of 8)
  fmt.Printf("Base Array: %v\n", baseArray)
  fmt.Printf( "Slice with len %d and capacity %d: %v\n", len(baseSlice), cap(baseSlice), baseSlice,)
  fmt.Println(pointer == fmt.Sprintf("%p", baseSlice))

  baseSlice = append(baseSlice, 11, 12, 13)
  // basic array's len is not enough for new elements so new slice start using new array
  fmt.Printf("Base Array: %v\n", baseArray)
  fmt.Printf( "Slice with len %d and capacity %d: %v\n", len(baseSlice), cap(baseSlice), baseSlice,)
  fmt.Println(pointer == fmt.Sprintf("%p", baseSlice))
}

func del_from_slice() {
  a := []int{1, 2, 3, 4, 5, 6, 7}
  fmt.Printf("%p\n", a)
  a = append(a[0:2], a[3:]...)  // unpacking of second slice
  fmt.Println(a)
  fmt.Printf("%p\n", a)  // the same pointer
  a = slices.Delete(a, 1, 2)
  fmt.Printf("%v\n", a)
}

func copy_slice() {
  a := []int{1, 2, 3}
  b := make([]int, 3, 3)  // three zeros
  fmt.Println(b)
  fmt.Printf("A: %v - %#v - %T\n", a, a, a)
  fmt.Printf("B: %v - %#v - %T\n", b, b, b)

  n := copy(b, a)  // copy from a to b and return quantity of copied elements
  fmt.Printf("B: %v - %#v - %T\n", b, b, b)
  fmt.Printf("Copied %d elements\n", n)
  fmt.Println(fmt.Sprintf("%p", a) == fmt.Sprintf("%p", b))  // derives from another array
}

func print_fourth_el() {
  var a uint8
  fmt.Scan(&a)
  workArray := make([]int, a)
  for idx := range workArray {
    fmt.Scan(&workArray[idx])
  }
  fmt.Println(workArray[3])
}

func find_max_in_arr() {
  array := [5]int{}
  for i := 0; i < len(array); i++ {
    fmt.Scan(&array[i])
  }
  var max int
  max = array[0]
  for _, el := range array {
    if el > max {
      max = el
    }
  }
  fmt.Println(max)
}

func only_evens() {
  var l int
  fmt.Scan(&l)
  var workArray [100]int = [100]int{}
  for idx := 0; idx < l; idx++ {
    fmt.Scan(&workArray[idx])
  }
  // fmt.Println(workArray)
  for idx, el := range workArray {
    if idx >= l {
      break
    }
    if idx % 2 == 0 {
      fmt.Printf("%d ", el)
    }
  }
}

func sum_of_positives() {
  var n, i uint8
  fmt.Scan(&n)
  var s, cur int
  for i = 0; i < n; i++ {
    fmt.Scan(&cur)
    if cur > 0 {
      s ++
    }
  }
  fmt.Println(s)
}


func main() {
  // first_arrays()
  // two_vars()
  // arr_from_input()
  // base_slices()
  del_from_slice()
  // copy_slice()
  // print_fourth_el()
  // find_max_in_arr()
  // only_evens()
  // sum_of_positives()
}

package main

import (
    "fmt"
)

func scanToPointers() {
    var p1 int
    var p2 int
    fmt.Scan(&p1, &p2)
    // testPointers(&p1, &p2)
    changePointersValues(&p1, &p2)
}

func testPointers(x1 *int, x2 * int) {
    res := *x1 * *x2
    fmt.Println(res)
}

func changePointersValues(x1 *int, x2 *int) {
    *x1, *x2 = *x2, *x1
    fmt.Println(*x1, *x2)
}

func pointerToPointer() {
    a := 100
    b := &a
    *b++
    c := &b
    **c++
    *b++
    fmt.Println(a)
}

func main() {
    scanToPointers()
}

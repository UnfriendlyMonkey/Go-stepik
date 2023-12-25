package main

import (
	"fmt"
	"container/list"
)

func printList(l *list.List) {
	for item := l.Front(); item != nil; item = item.Next() {
		fmt.Printf("%v ",item.Value)
	}
	fmt.Println("")
}

func listExample() {
	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	myList.PushFront(4)
	myList.PushFront(5)
	elem := myList.PushFront(6)
	printList(myList)
	myList.Remove(elem)
	// should be the pointer to element, not it's value
	deletedItem := myList.Remove(myList.Back())
	fmt.Println(deletedItem)
	printList(myList)

	myList.PushBack(11)
	myList.PushBack(12)

	for el := myList.Front(); el != nil; {
		next := el.Next()
		if el.Value.(int) % 2 == 0 {
			myList.Remove(el)
		}
		el = next
	}

	printList(myList)
}

func ReverseList(l *list.List) *list.List {
	reversedList := list.New()
	for i := l.Front(); i != nil; i = i.Next() {
		reversedList.PushFront(i.Value)
	}
	return reversedList
}

func testReverseList() {
	checkList := list.New()
	for i:=0; i<10; i++ {
		checkList.PushFront(i)
	}
	printList(checkList)
	reversedList := ReverseList(checkList)
	printList(reversedList)
}

func main() {
	// listExample()
	testReverseList()
}

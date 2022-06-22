// From singly linked list remove Nth element from the end.
// Из односвязного списка удалить N-й элемент с конца.
package main

import (
	"container/list"
	"fmt"
)

// Result:
// Input:  [1,2,3,4,5], n = 2
// Output: [1,2,3,5]

func main() {
	// Creating a list with some numbers in it.
	// Создание списка с числами.
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)
	l.InsertAfter(5, e4)
	l.Remove(e4)

	// Iterate through list and print its contents.
	// Перебор списка и вывод его содержимого.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

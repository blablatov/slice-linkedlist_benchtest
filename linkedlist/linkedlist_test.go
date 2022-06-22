package main

import (
	"container/list"
	"fmt"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < 1000; i++ {
		l := list.New()
		e4 := l.PushBack(4)
		e1 := l.PushFront(1)
		l.InsertBefore(3, e4)
		l.InsertAfter(2, e1)
		l.InsertAfter(5, e4)
		l.Remove(e4)

		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}
	}
}

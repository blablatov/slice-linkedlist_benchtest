package main

import (
	"fmt"
	"testing"
)

func BenchmarkRemove(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < 1000; i++ {
		s := []int{1, 2, 3, 4, 5}
		fmt.Println(remove(s, 3))
	}
}

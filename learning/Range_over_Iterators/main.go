package main

import (
	"fmt"
	"iter"
)

func iterate[T any](list []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for _, val := range list {
			if !yield(2, val) {
				return
			}
		}
	}

}

func main() {
	intlist := []int{1, 2, 3}
	for i, v := range iterate(intlist) {
		fmt.Println(i, "->", v)
	}
}

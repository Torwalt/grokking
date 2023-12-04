package main

import (
	"fmt"
	"grokking/exercises/quicksort"
)

func main() {
	in := []int{2, 5, 12, 1, 7, 9, 111, 4, 300000, 12}
	out := quicksort.QuickSort(in)

	fmt.Println(out)
}

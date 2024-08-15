package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [10]int{2, 4, 6, 8, 9, 10}
	target := 4

	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	if index < len(arr) && arr[index] == target {
		fmt.Printf("элемент %d найден на позиции %d\n", target, index)
	} else {
		fmt.Printf("не пизди, нет такого %d", target)
	}
}

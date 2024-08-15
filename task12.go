//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

func main() {
	my_array := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})
	for _, v := range my_array {
		set[v] = struct{}{}
	}

	for v := range set {
		fmt.Println(v)
	}
}

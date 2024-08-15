package main

import "fmt"

func main() {
	arr := [8]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tempGroups := make(map[int][]float32)

	for _, digit := range arr {
		key := int(digit/10) * 10
		//fmt.Println(digit)
		tempGroups[key] = append(tempGroups[key], digit)
	}

	for k, v := range tempGroups {
		fmt.Println(k, v)
	}
}

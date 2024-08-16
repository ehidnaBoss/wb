//Удалить i-ый элемент из слайса.

package main

import "fmt"

func removeElement(anySlice []int, rmNum int) []int {
	var newSlice []int
	for i := range anySlice {
		if i+1 != rmNum {
			newSlice = append(newSlice, anySlice[i])
		}
	}
	return newSlice
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var rm int

	fmt.Println("Введите i-ый элемент, для удаления из слайса: ")
	_, err := fmt.Scan(&rm)
	if err != nil {
		fmt.Println("Ошибка!", err)
		return
	}
	if rm < 0 || rm > len(slice) {
		fmt.Println("Нельзя удалить такой элемент!")
		return
	}

	resultSlice := removeElement(slice, rm)

	fmt.Println("Обновленный слайс: ", resultSlice)

}

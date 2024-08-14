package main

import "fmt"

func setBitToOne(value int64, bitIndex int) int64 {
	return value | (1 << bitIndex)
}
func setBitToZero(value int64, bitIndex int) int64 {
	return value &^ (1 << bitIndex)
}

func main() {
	var digit int64
	fmt.Println("Введите число: ")
	_, err := fmt.Scan(&digit)
	if err != nil {
		fmt.Println("Ошибка ввода числа:", err)
		return
	}

	var bitIndex int
	fmt.Println("Введите индекс бита (от 0 до 63): ")
	_, err = fmt.Scan(&bitIndex)
	if err != nil {
		fmt.Println("Ошибка ввода индекса бита:", err)
		return
	}

	if bitIndex < 0 || bitIndex > 63 {
		fmt.Println("Индекс бита должен быть в диапазоне от 0 до 63")
		return
	}
	var bitValue int
	fmt.Println("Введите значение бита(0 или 1): ")
	_, err = fmt.Scan(&bitValue)
	if err != nil {
		fmt.Println("Ошибка ввода значения бита:", err)
		return
	}
	if bitValue != 0 && bitValue != 1 {
		fmt.Println("Только 0 или 1!")
		return
	}
	if bitValue == 1 {
		digit = setBitToOne(digit, bitIndex)
	} else {
		digit = setBitToZero(digit, bitIndex)
	}
	fmt.Println(digit)
}

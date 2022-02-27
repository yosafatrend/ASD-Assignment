package main

import (
	"fmt"
	"strconv"
	"strings"
)

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func bubbleSort(array []int, size int) {
	for i := 1; i < size; i++ {
		for j := size - 1; j >= i; j-- {
			if array[j] < array[j-1] {
				swap(&array[j], &array[j-1])
			}
		}
	}
}

func printArray(array []int, size int) {
	for i := 0; i < size; i++ {
		fmt.Print(array[i], " ")
	}
}

func arrayToInt(array []string, size int) []int {
	s := make([]int, size)

	for i := 0; i < size; i++ {
		value, err := strconv.Atoi(array[i])
		if err != nil {
			panic("Inputan salah. Ex : 4,8,23,5,12,43")
		}
		s[i] = value
	}

	return s
}

func main() {
	fmt.Println("---- Bubble Sort Algorithm ----")
	fmt.Print("Array Values (comma separated) : ")

	var input string
	fmt.Scanln(&input)

	arrayString := strings.Split(input, ",")

	size := len(arrayString)
	arrayInt := arrayToInt(arrayString, size)
	bubbleSort(arrayInt, size)

	fmt.Print("Sorted : ")
	printArray(arrayInt, size)
}

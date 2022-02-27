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

func selection_sort(array []int, size int) {
	for i := 0; i < size; i++ {
		pos := i
		for j := i; j < size; j++ {
			if array[j] < array[pos] {
				pos = j
			}
		}
		array[i], array[pos] = array[pos], array[i]
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
	fmt.Println("---- Selection Sort Algorithm ----")
	fmt.Print("Array Values (comma separated) : ")

	var input string
	fmt.Scanln(&input)

	arrayString := strings.Split(input, ",")

	size := len(arrayString)
	arrayInt := arrayToInt(arrayString, size)
	selection_sort(arrayInt, size)

	fmt.Print("Sorted : ")
	printArray(arrayInt, size)
}

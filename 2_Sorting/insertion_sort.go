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

func insertion_sort(array []int, size int) {
	for i := 1; i < size; i++ {
		j := i
		for j > 0 {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
			j = j - 1
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
	fmt.Println("---- Insertion Sort Algorithm ----")
	fmt.Print("Array Values (comma separated) : ")

	var input string
	fmt.Scanln(&input)

	arrayString := strings.Split(input, ",")

	size := len(arrayString)
	arrayInt := arrayToInt(arrayString, size)
	insertion_sort(arrayInt, size)

	fmt.Print("Sorted : ")
	printArray(arrayInt, size)
}

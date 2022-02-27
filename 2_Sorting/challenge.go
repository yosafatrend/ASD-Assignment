package main

import "fmt"

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

func exchangeSort(array []int, size int) {
	for i := 0; i < size-1; i++ {
		for j := (i + 1); j < size; j++ {
			if array[i] > array[j] {
				swap(&array[i], &array[j])
			}
		}
	}
}

func printArray(array []int, size int) {
	for i := 0; i < size; i++ {
		fmt.Print(array[i], " ")
	}
}

func main() {
	array := []int{
		23, 65, 30, 8, 33, 24, 76, 7,
	}

	size := len(array)

	fmt.Print("Array : ")
	printArray(array, size)
	fmt.Println("\nChoose algorithm :")
	fmt.Println("1. Bubble Sort ")
	fmt.Println("2. Exchange Sort ")
	fmt.Print("Pilihan : ")

	var input int
	fmt.Scan(&input)

	switch input {
	case 1:
		bubbleSort(array[:], size)
		printArray(array, size)

	case 2:
		exchangeSort(array[:], size)
		printArray(array, size)
	}
}

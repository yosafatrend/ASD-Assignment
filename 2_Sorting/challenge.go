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
	bubbleSort(array[:], size)
	//exchangeSort(array[:], size) //uncomment to sort with exchange sort method
	printArray(array, size)
}

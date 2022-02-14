package main

import "fmt"

func main() {
	var array = []int{
		1, 4, 7, 9, 16, 56, 70,
	}
	n := 7
	element := 56
	found_index := iterativeBinarySearch(array, 0, n-1, element)
	if found_index == -1 {
		fmt.Println("Angka tersebut TIDAK DITEMUKAN")
	} else {
		fmt.Println("Angka ditemukan pada lokasi index :", found_index)
	}
}

func iterativeBinarySearch(array []int, start_index int, end_index int, element int) int {
	for start_index <= end_index {
		middle := start_index + (end_index-start_index)/2
		if array[middle] == element {
			return middle
		}
		if array[middle] < element {
			start_index = middle + 1
		} else {
			end_index = middle - 1
		}
	}
	return -1
}

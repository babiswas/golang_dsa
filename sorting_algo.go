package main

import (
	"fmt"
)

// Function to perform selection sort of the array
func sort_arr(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

// Function to perform bubble sort of the array.
func bubble_sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
		fmt.Println("After", " ", i, " stage", arr)
	}
}

// Function to perform quick sort of the array.
func partition(arr []int, start_index int, end_index int) int {
	pivot := arr[end_index]
	wall := start_index - 1
	for i := start_index; i < end_index+1; i++ {
		if arr[i] < pivot {
			wall += 1
			arr[wall], arr[i] = arr[i], arr[wall]
		}
	}
	wall += 1
	arr[wall], arr[end_index] = arr[end_index], arr[wall]
	return wall
}

func qsort(arr []int, start_index int, end_index int) {
	if start_index < end_index {
		index := partition(arr, start_index, end_index)
		fmt.Println(arr)
		qsort(arr, start_index, index-1)
		qsort(arr, index+1, end_index)
	}
}

// Performing inserion sort in golang
func insertion_sort(arr []int, start_index int, end_index int) {
	for i := 1; i < len(arr); i++ {
		fix := arr[i]
		hole := i
		for hole > 0 && arr[hole-1] > fix {
			arr[hole] = arr[hole-1]
			hole = hole - 1
		}
		arr[hole] = fix
	}
}

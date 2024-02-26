package main

import "fmt"

const size1 = 4
const size2 = 5
const mergedSize = size1 + size2

func arrMerge(array1 [size1]int, array2 [size2]int) [mergedSize]int {
	var mergedArray [mergedSize]int
	for i := 0; i < len(array1); i++ {
		mergedArray[i] = array1[i]
	}
	for j := 0; j < len(array2); j++ {
		mergedArray[len(array1)+j] = array2[j]
	}
	return mergedArray
}

func arrSort(input [mergedSize]int) [mergedSize]int {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(input); i++ {
			if input[i-1] > input[i] {
				input[i], input[i-1] = input[i-1], input[i]
				swapped = true
			}
		}
	}
	return input
}

func main() {
	arr1 := [size1]int{14, 22, 86, 6}
	arr2 := [size2]int{53, 18, 1, 145, 55}
	finalArray := arrMerge(arr1, arr2)
	fmt.Println(arrSort(finalArray))
}
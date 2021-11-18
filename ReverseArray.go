package main

import "fmt"

func reverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func main() {
	fmt.Println(reverseArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(reverseArray([]int{3, 5, 7, 2, 1}))
	fmt.Println(reverseArray([]int{9, 8, 6, 1, 0}))
}

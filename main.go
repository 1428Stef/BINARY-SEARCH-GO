package main

import (
	"fmt"
)

func binarySearch(list []int, item int) *int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]

		if guess == item {
			return &mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return nil
}

func main() {
	myList := []int{1, 3, 5, 7, 9}

	fmt.Println(binarySearch(myList, 3))
	fmt.Println(binarySearch(myList, -1))

}

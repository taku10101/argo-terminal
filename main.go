package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	target := 5
	fmt.Println(binarySearch(arr, target))
}
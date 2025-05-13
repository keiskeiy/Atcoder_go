package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	field := []rune(scanner.Text())

	result := string(mergeSort(field))
	fmt.Println(result)
}

func mergeSort(arr []rune) []rune {
	if len(arr) <= 1 {
		return arr
	}

	center := len(arr) / 2
	left := arr[:center]
	right := arr[center:]

	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}

func merge(left []rune, right []rune) []rune {
	l, r := 0, 0

	result := make([]rune, 0, len(left)+len(right))

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}

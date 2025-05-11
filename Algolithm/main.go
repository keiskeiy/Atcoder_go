package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var asum, bsum int
	var input []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	field := strings.Fields(scanner.Text())
	for _, i := range field {
		num, _ := strconv.Atoi(i)
		input = append(input, num)
	}
	bubble := append([]int{}, input...)
	insert := append([]int{}, input...)
	merge := append([]int{}, input...)
	quick := append([]int{}, input...)
	heap := append([]int{}, input...)

	//普通のソート
	sort.Sort(sort.Reverse(sort.IntSlice(input)))
	fmt.Println(bubble)
	for index, i := range input {
		if index%2 == 0 {
			asum += i
		} else {
			bsum += i
		}
	}

	// バブルソート O(n)
	for i := 0; i < len(bubble); i++ {
		for j := 0; j < len(bubble)-1-i; j++ {
			if bubble[j] < bubble[j+1] {
				bubble[j], bubble[j+1] = bubble[j+1], bubble[j]
			}
		}
	}

	//挿入ソート O(n)
	for i := 1; i < len(insert); i++ {
		key := insert[i]
		j := i - 1

		for j >= 0 && insert[j] < key {
			insert[j+1] = insert[j]

			j--
		}
		insert[j+1] = key
	}

	//まーじソート
	merge = mergeSort(merge)
	//quickソート
	quick = pivotSort(quick)
	//ヒープソート
	heap = heapSort(heap)
	//挿入ソート
	fmt.Println(insert)
	//バブルソート
	fmt.Println(bubble)
	//まーじソート
	fmt.Println(merge)
	//quiclソート
	fmt.Println(quick)
	//ヒープソート
	fmt.Println(heap)
	//挿入ソート
	fmt.Println(input)
	// fmt.Println(asum - bsum)

}

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	center := len(arr) / 2
	left := arr[:center]
	right := arr[center:]

	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
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

func pivotSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]

	var high, equal, low []int

	for _, i := range arr {
		if i > pivot {
			high = append(high, i)
		} else if i == pivot {
			equal = append(equal, i)
		} else {
			low = append(low, i)
		}
	}

	high = pivotSort(high)
	low = pivotSort(low)

	high = append(high, equal...)
	high = append(high, low...)

	return high
}

func heapify(arr []int, n int, i int) {
	smallest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && arr[l] < arr[smallest] {
		smallest = l
	}

	if r < n && arr[r] < arr[smallest] {
		smallest = r
	}

	if smallest != i {
		arr[smallest], arr[i] = arr[i], arr[smallest]
		heapify(arr, n, smallest)
	}
}

func heapSort(arr []int) []int {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// ソートフェーズ: ヒープから一つずつ最大要素を取り出して配列の後ろから埋めていく
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0] // ルートと最後の要素を交換
		heapify(arr, i, 0)
	}

	return arr
}

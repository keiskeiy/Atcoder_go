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

	fmt.Println(insert)
	fmt.Println(bubble)
	fmt.Println(input)
	// fmt.Println(asum - bsum)
}

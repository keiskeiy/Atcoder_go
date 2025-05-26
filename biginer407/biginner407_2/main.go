package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	field := strings.Fields(scanner.Text())
	var input []int

	set := make(map[int]struct{})

	for _, i := range field {
		tmp, _ := strconv.Atoi(i)
		input = append(input, tmp)
	}

	for i := 1; i <= 6; i++ {
		for j := 1; j <= 6; j++ {
			if i+j >= input[0] {
				set[i+j*100] = struct{}{}
			}
			sub := i - j
			if sub < 0 {
				sub = -sub
			}
			if sub >= input[1] {
				set[i+j*100] = struct{}{}
			}
		}
	}
	result := float64(len(set))
	result = result / 36.0
	fmt.Println(result)
}

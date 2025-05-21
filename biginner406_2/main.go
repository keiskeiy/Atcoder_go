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
	line1 := InputConvert(scanner)
	line2 := InputConvert(scanner)
	result := 1
	max := 1

	for i := 0; i < line1[1]; i++ {
		max *= 10
	}

	for i := 0; i < line1[0]; i++ {
		result *= line2[i]
		if result >= max {
			result = 1
		}
	}
	fmt.Println(result)
}

func InputConvert(reader *bufio.Scanner) []int {
	reader.Scan()
	input := strings.Fields(reader.Text())
	var n []int
	for _, i := range input {
		tmp, _ := strconv.Atoi(i)
		n = append(n, tmp)
	}
	return n
}

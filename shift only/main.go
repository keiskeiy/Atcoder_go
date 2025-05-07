package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	count := 0
	flag := true

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Scan()
	Field := strings.Fields(scanner.Text())

	for flag {
		for index, num := range Field {
			intnum, _ := strconv.Atoi(num)
			result := intnum % 2
			if result != 0 {
				flag = false
				break
			} else {
				Field[index] = strconv.Itoa(intnum / 2)
			}
		}
		count += 1
	}

	fmt.Println(count - 1)
}

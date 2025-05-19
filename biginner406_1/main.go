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
	for _, i := range field {
		tmp, _ := strconv.Atoi(i)
		input = append(input, tmp)
	}
	if input[0]-input[2] > 0 {
		fmt.Println("Yes")
		return
	}
	if input[0]-input[2] == 0 && input[1]-input[3] > 0 {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

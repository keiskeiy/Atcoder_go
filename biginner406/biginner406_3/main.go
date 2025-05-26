package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Scan()
	field := strings.Fields(scanner.Text())
	max := false
	min := false
	var input []int

	for _, i := range field {
		tmp, _ := strconv.Atoi(i)
		input = append(input, tmp)
	}
}

func Check(before int, now int, after int) {

}

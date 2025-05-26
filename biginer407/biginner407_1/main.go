package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"structs"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	field := strings.Fields(scanner.Text())
	var input []float64

	structs dise {
		x int;
		y int;
	}

	for _, i := range field {
		tmp, _ := strconv.Atoi(i)
		input = append(input, float64(tmp))
	}

	result := input[0] / input[1]
	num := math.Floor(result)

	if result-num >= num+1-result {
		fmt.Println(int(num) + 1)
	} else {
		fmt.Println(int(num))
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var x, y, z int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	field := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(field[0])
	sum, _ := strconv.Atoi(field[1])

	if sum > 10000*n {
		x, y, z = -1, -1, -1
		fmt.Printf("%d %d %d", x, y, z)
		return
	}

	for i := n; i >= 0; i-- {
		remain := sum - 10000*i
		for j := n - i; j >= 0; j-- {
			middle := remain - 5000*j
			if middle == 0 && i+j == n {
				x, y, z = i, j, 0
				fmt.Printf("%d %d %d", x, y, z)
				return
			}
			for k := n - i - j; k >= 0; k-- {
				result := middle - 1000*k
				if result == 0 && i+j+k == n {
					x, y, z = i, j, k
					fmt.Printf("%d %d %d", x, y, z)
					return
				}
			}
		}
	}

	x, y, z = -1, -1, -1
	fmt.Printf("%d %d %d", x, y, z)

}

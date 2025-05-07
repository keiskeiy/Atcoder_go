package main

import(
	"fmt"
	"bufio"
	"strconv"
	"os"
)

func main() {
	counter := 0
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	a,_ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	b,_ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	c,_ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	sum,_ := strconv.Atoi(scanner.Text())

	for i := a; i >= 0; i-- {
		for j := b; j >= 0; j-- {
			remain := sum - 500 * i - 100 * j
			if remain < 0 {
				continue
			}
			if remain == 0 || (remain / 50 <= c) {
				counter += 1
				continue
			}
		}
	}
	fmt.Println(counter)
}

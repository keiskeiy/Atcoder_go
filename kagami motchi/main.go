package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	set := make(map[string]struct{})

	for i := 0; i < n; i++ {
		scanner.Scan()
		set[scanner.Text()] = struct{}{}
	}
	fmt.Println(len(set))
}

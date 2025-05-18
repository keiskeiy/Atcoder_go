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
	count, _ := strconv.Atoi(scanner.Text())
	if count == 0 {
		return
	}
	root := make(map[int][]int)

	for i := count; i > 0; i-- {
		scanner.Scan()
		field := strings.Fields(scanner.Text())
		var in []int
		for _, i := range field {
			tmp, _ := strconv.Atoi(i)
			in = append(in, tmp)
		}
		root[i] = in
	}

	//入力部分の数値読み取り
	for i := count; i > 0; i-- {

		if i == count {
			if root[i][0]%2 == (root[i][1]+root[i][2])%2 && root[i][0] >= root[i][1]+root[i][2] {
				continue
			}
			fmt.Println("No")
			return
		}

		s := root[i][0] - root[i+1][0]
		v := root[i][2] - root[i+1][2]
		h := root[i][1] - root[i+1][1]

		length := h + v
		if length < 0 {
			length = -length
		}

		if length > s {
			fmt.Println("No")
			return
		}

		if s%2 == 0 {
			if length%2 != 0 {
				fmt.Println("No")
				return
			}
			continue
		} else {
			if length%2 != 1 {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}

package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	line1 := InputConvert(scanner)
	line2 := InputConvert(scanner)

	result := big.NewInt(1)
	max := big.NewInt(1)

	for i := 0; i < line1[1]; i++ {
		tmp := big.NewInt(10)
		max = tmp.Mul(tmp, max)
	}

	for i := 0; i < line1[0]; i++ {
		multi_tmp := big.NewInt(int64(line2[i]))
		result = multi_tmp.Mul(multi_tmp, result)
		if result.Cmp(max) >= 0 {
			result = big.NewInt(1)
		}
	}
	fmt.Println(result.String())
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

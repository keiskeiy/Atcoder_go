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
	input := scanner.Text()
	input = reverse(input)
	counter := 0
	plus := 0
	length := len(input)

	for i := 0; i < length; i++ {
		num, _ := strconv.Atoi(string(input[i]))
		plus += 1
		tmp := num - counter
		if tmp < 0 {
			tmp = -tmp
		}
		if tmp >= 10 {
			tmp = tmp % 10
		}
		counter += tmp
	}
	fmt.Println(counter + plus)
}

func reverse(s string) string {
	if len(s) <= 1 {
		return s
	}

	runes := []rune(s)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}

	return string(runes)
}

func reverse2(s string) string {
	runes := []rune(s)
	length := len(runes)
	var r []rune

	for i := length - 1; i >= 0; i-- {
		r = append(r, runes[i])
	}
	return string(r)
}

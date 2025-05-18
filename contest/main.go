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
	var N, M = 36, 12

	input := make(map[int]string)

	var max = 0

	for i := N; i > 0; i-- {
		scanner.Scan()
		field := strings.Fields(scanner.Text())
		score, _ := strconv.Atoi(field[1])
		ex := score / len(field[0])
		input[ex] = field[0]

		//1文字あたりの期待値高いのを入れる
		if max < ex {
			max = ex
		}
	}

	length := len(input[max])
	output(input[max])

	if length < M {
		max2 := 0
		for ex2, value := range input {
			if len(value) <= M-length {
				if max < ex2 {
					max = ex2
				}
			}
		}
		if max2 != 0 {
			output(input[max2])
			length = M - length - len(input[max2])
		}

		for i := M - length; i > 0; i-- {
			output := [12]int{100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
			fmt.Printf("%c %d %d %d %d %d %d %d %d %d %d %d %d\n", 'a', output[0], output[1], output[2], output[3], output[4], output[5], output[6], output[7], output[8], output[9], output[10], output[11])
		}
	}

}

func output(sentence string) {
	for index, s := range sentence {
		output := [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		if index == len(sentence)-1 {
			output[0] = 100
		} else {
			output[index+1] = 100
		}
		fmt.Printf("%c %d %d %d %d %d %d %d %d %d %d %d %d\n", s, output[0], output[1], output[2], output[3], output[4], output[5], output[6], output[7], output[8], output[9], output[10], output[11])
	}
}

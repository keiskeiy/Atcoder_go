package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	length := len(input)

	pointer := 0
	//dream dreamer erase eraser
	for pointer < length-1 {
		if len(input) >= 5 && input[:5] == "dream" {
			if len(input) >= 7 && input[:7] == "dreamer" {
				if len(input) >= 8 && input[5:8] != "ase" {
					pointer += 5
					input = input[5:]
					continue
				}
				if len(input) >= 9 && input[5:9] != "aser" {
					pointer += 5
					input = input[5:]
					continue
				}
				pointer += 7
				if pointer == length {
					fmt.Println("YES")
					break
				}
				input = input[7:]
				continue
			}
			pointer += 5
			if pointer == length {
				fmt.Println("YES")
				break
			}
			input = input[5:]
			continue
		}
		if len(input) >= 5 && input[:5] == "erase" {
			if len(input) >= 6 && input[:6] == "eraser" {
				pointer += 6
				if pointer == length {
					fmt.Println("YES")
					break
				}
				input = input[6:]
				continue
			}
			pointer += 5
			if pointer == length {
				fmt.Println("YES")
				break
			}
			input = input[5:]
			continue
		}
		break
	}
	if pointer != length {
		fmt.Println("NO")
	}
}

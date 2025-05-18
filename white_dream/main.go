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
	if input == "" {
		fmt.Println("NO")
		return
	}
	pointer := len(input)
	for pointer > 0 {
		if pointer >= 7 && input[pointer-7:pointer] == "dreamer" {
			pointer -= 7
			continue
		}
		if pointer >= 6 && input[pointer-6:pointer] == "eraser" {
			pointer -= 6
			continue
		}
		if pointer >= 5 && input[pointer-5:pointer] == "dream" {
			pointer -= 5
			continue
		}
		if pointer >= 5 && input[pointer-5:pointer] == "erase" {
			pointer -= 5
			continue
		}
		break
	}
	if pointer != 0 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}

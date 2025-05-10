package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"strconv"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	text := strings.Fields(scanner.Text())
	N, _ := strconv.Atoi(text[0])
	a, _ := strconv.Atoi(text[1])
	b, _ := strconv.Atoi(text[2])

	var ans int

	for i:=0; i <=N; i++{
		tth := i / 10000
		n4 := i - (tth * 10000)
		th := n4 / 1000
		n3 := n4 - (th * 1000)
		h := n3 / 100
		n2 := n3 - (h * 100)
		t := n2 / 10
		n1 := n2 - (t * 10)
		o := n1 % 10
		if a <= tth + th + h + t + o && tth + th + h + t + o <= b{
			ans += i
		}
	}
	fmt.Println(ans)
}
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	//Scan関数は行があればtrueを返し、なければfalseを返す
	//i.e. ここではwhile文を書いている
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

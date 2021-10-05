package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//https://www.acmicpc.net/problem/1152
func main() {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	result := strings.Split(line, " ")
	var cnt = 0
	for _, v := range result {
		if v == " " || v == "" || v == "\n" {
			continue
		} else {
			cnt++
		}
	}
	fmt.Printf("%d", cnt)
}

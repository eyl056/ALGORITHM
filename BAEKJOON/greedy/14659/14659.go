package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var size int
	_, _ = fmt.Scanf("%d", &size)

	var arr = make([]int, size+1, size+1)

	scanner.Scan()
	str := strings.Split(scanner.Text(), " ")
	for i := 0; i < size; i++ {
		arr[i], _ = strconv.Atoi(str[i])
	}

	var max int
	for i := 0; i < size; i++ {
		var tmp int
		for j := i + 1; j < size; j++ {
			if arr[i] > arr[j] {
				tmp++
			} else {
				break
			}
		}
		if max < tmp {
			max = tmp
		}
	}
	fmt.Println(max)
}

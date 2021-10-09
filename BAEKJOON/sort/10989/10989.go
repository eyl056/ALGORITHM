package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var size int

	scanner.Scan()
	size, _ = strconv.Atoi(scanner.Text())

	var countArr = make([]int, 10001)

	for i := 0; i < size; i++ {
		var tmp int
		scanner.Scan()
		tmp, _ = strconv.Atoi(scanner.Text())
		countArr[tmp]++
	}

	for i := 1; i < 10001; i++ {
		for j := 0; j < countArr[i]; j++ {
			_, _ = fmt.Fprintln(writer, i)
		}
	}
}

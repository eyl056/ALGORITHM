package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getResult(arr [][]string, size *int) (result int) {
	result = 1
	for i := 0; i < *size; i++ {
		tmp := 1
		for j := 1; j < *size; j++ {
			if arr[i][j-1] == arr[i][j] {
				tmp++
			} else {
				if result < tmp {
					result = tmp
					tmp = 1
				}
			}
		}
		if result < tmp {
			result = tmp
		}
	}
	for i := 0; i < *size; i++ {
		tmp := 1
		for j := 1; j < *size-1; j++ {
			if arr[j+1][i] == arr[j][i] {
				tmp++
			} else {
				if result < tmp {
					result = tmp
					tmp = 1
				}
			}
		}
		if result < tmp {
			result = tmp
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var size int
	_, _ = fmt.Scanf("%d", &size)
	var arr = make([][]string, size, size)
	for i := 0; i < size; i++ {
		scanner.Scan()
		arr[i] = strings.Split(scanner.Text(), "")
	}
	fmt.Println(arr)
	var result = 0
loop:
	for i := 0; i < size; i++ {
		for j := 0; j < size-1; j++ {
			arr[i][j], arr[i][j+1] = arr[i][j+1], arr[i][j]
			if tmp := getResult(arr, &size); result < tmp {
				result = tmp
			}
			arr[i][j], arr[i][j+1] = arr[i][j+1], arr[i][j]

			arr[j][i], arr[j+1][i] = arr[j+1][i], arr[j][i]
			if tmp := getResult(arr, &size); result < tmp {
				result = tmp
			}
			arr[j][i], arr[j+1][i] = arr[j+1][i], arr[j][i]

			if result == size {
				break loop
			}
		}
	}
	_, _ = fmt.Fprintf(writer, "%d\n", result)
}

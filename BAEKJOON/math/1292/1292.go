package main

import "fmt"

func main() {
	var a, b, idx, result int
	_, _ = fmt.Scanf("%d %d", &a, &b)
	var arr = make([]int, 1001, 1001)
	for i := 1; i <= 1000; i++ {
		for j := 1; j <= i; j++ {
			arr[idx] = i
			if idx >= 1000 {
				break
			}
			idx++
		}
	}
	for i := a - 1; i < b; i++ {
		result += arr[i]
	}
	fmt.Println(result)
}

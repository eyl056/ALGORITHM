package main

import "fmt"

func main() {
	var a, b, result int
	_, _ = fmt.Scanf("%d", &a)
	_, _ = fmt.Scanf("%d", &b)
	var min = b
	for i := a; i <= b; i++ {
		var count = 0
		for j := 2; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count == 1 {
			result += i
			if min > i {
				min = i
			}
		}
	}
	if result == 0 {
		fmt.Println(-1)
	} else {
		fmt.Printf("%d\n%d", result, min)
	}
}

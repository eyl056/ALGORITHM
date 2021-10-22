package main

import (
	"fmt"
)

func main() {
	var money, YEN500, YEN100, YEN50, YEN10, YEN5 int = 0, 500, 100, 50, 10, 5
	_, _ = fmt.Scanf("%d", &money)
	money = 1000 - money
	var result = 0
	if money >= YEN500 {
		money -= YEN500
		result++
	}
	if money >= YEN100 {
		tmp := money / YEN100
		money -= tmp * YEN100
		result += tmp
	}
	if money >= YEN50 {
		tmp := money / YEN50
		money -= tmp * YEN50
		result += tmp
	}
	if money >= YEN10 {
		tmp := money / YEN10
		money -= tmp * YEN10
		result += tmp
	}
	if money >= YEN5 {
		tmp := money / YEN5
		money -= tmp * YEN5
		result += tmp
	}
	if money < YEN5 {
		result += money
	}
	fmt.Println(result)
}

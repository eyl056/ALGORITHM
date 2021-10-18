package main

import (
	"fmt"
)

func getResult(result int, idx int, operator []int, arr []int, max *int, min *int) (int, int) {
	if idx == len(arr) {
		if *max < result || *max == 2100000000 {
			*max = result
		}
		if *min > result || *min == -2100000000 {
			*min = result
		}
		return *max, *min
	}
	for i := 0; i < 4; i++ {
		if operator[i] > 0 {
			operator[i]--
			switch i {
			case 0:
				getResult(result+arr[idx], idx+1, operator, arr, max, min)
				break
			case 1:
				getResult(result-arr[idx], idx+1, operator, arr, max, min)
				break
			case 2:
				getResult(result*arr[idx], idx+1, operator, arr, max, min)
				break
			case 3:
				getResult(result/arr[idx], idx+1, operator, arr, max, min)
				break
			}
			operator[i]++
		}
	}
	return *max, *min
}

func main() {
	var size, max, min int = 0, 2100000000, -2100000000
	_, _ = fmt.Scanf("%d", &size)
	var arr = make([]int, size, size)
	for i := 0; i < size; i++ {
		_, _ = fmt.Scanf("%d", &arr[i])
	}
	var operator = make([]int, 4, 4)
	for i := 0; i < 4; i++ {
		_, _ = fmt.Scanf("%d", &operator[i])
	}
	maxVal, minVal := getResult(arr[0], 1, operator, arr, &max, &min)
	fmt.Println(maxVal)
	fmt.Println(minVal)
}

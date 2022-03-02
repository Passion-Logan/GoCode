package main

import (
	"fmt"
	"math"
)

/*
用 Golang 实现，求素数个数。
判断 101-200 之间有多少个素数，并输出所有素数。
*/
func main() {
	var allNum []int
	var j, k = 0, 0

	for i := 101; i <= 200; i++ {
		k = int(math.Sqrt(float64(i)))
		for j = 2; j <= k; j++ {
			if i%j == 0 {
				break
			}
		}
		if j == k+1 {
			allNum = append(allNum, i)
		}
	}

	fmt.Printf("101-200之间有%d个素数，分别是%v", len(allNum), allNum)
}

package main

import (
	"fmt"
)

/*
用 Golang 实现，打印出所有的水仙花数。
打印出所有的 “水仙花数”，所谓 “水仙花数” 是指一个三位数，其各位数字立方和等于该数本身。例如：153 是一个 “水仙花数”，因为 153=1 的三次方＋5 的三次方＋3 的三次方。
*/
func main() {
	var allNum []int
	a, b, c := 0, 0, 0
	for i := 100; i < 1000; i++ {
		a = i / 100
		b = i / 10 % 10
		c = i % 10

		if (a*a*a + b*b*b + c*c*c) == i {
			allNum = append(allNum, i)
		}
	}

	fmt.Printf("所有的水仙花数：%v", allNum)

}

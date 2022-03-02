package main

import "fmt"

/*
用 Golang 实现，数字从小到大排序。
输入三个 整数 x，y，z，请把这三个数由小到大输出。
*/
func main() {
	var a, b, c = 0, 0, 0
	fmt.Printf("请输入三个数字\n")
	fmt.Scanf("%d%d%d", &a, &b, &c)

	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}
	fmt.Printf("%d < %d < %d\n", a, b, c)
}

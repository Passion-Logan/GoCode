package main

import "fmt"

/*
Golang兔子问题
古典问题：有一对兔子，从出生后第 3 个月起每个月都生一对兔子，小兔子长到第三个月后每个月又生一对兔子，假如兔子都不死，问每个月的兔子总数为多少？
输入一个月份，输入当月兔子的数量

1，1，2，3，5，8，13，21…
*/
func main() {

	var a, b = 1, 1
	fmt.Print("请输入月份：\n")
	var month = 1
	var d = 1
	fmt.Scanf("%d", &month)
	if month < 3 {
		fmt.Printf("%d月兔子有%d只", month, d)
	} else {
		for i := 3; i <= month; i++ {
			d = a + b
			a = b
			b = d
		}
		fmt.Printf("%d月兔子有%d只", month, d)
	}

}

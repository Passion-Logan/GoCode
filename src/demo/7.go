package main

import "fmt"

/*
Go语言打印乘法口诀
输出 9*9 乘法口诀表。
 */
func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %-3d  ", j, i, i * j)
		}
		fmt.Print("\n")
	}

}

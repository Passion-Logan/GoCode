package main

import (
	"fmt"
)

func fibonacci1(n int) (ind int, res int) {
	if n <= 1 {
		res = 1
	} else {
		_, r1 := fibonacci1(n - 1)
		_, r2 := fibonacci1(n - 2)
		res = r1 + r2
	}
	ind = n
	return
}

func sout(a int) {
	fmt.Println(a)
	if a < 10 {
		sout(a + 1)
	}
}

func factorial(i int) int {
	if i <= 1 {
		return 1
	}

	return i * factorial(i-1)
}

func main() {
	fmt.Println(factorial(4))
}

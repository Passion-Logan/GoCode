package main

import (
	"fmt"
	"sort"
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

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}
}

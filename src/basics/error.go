package main

import (
	"errors"
	"fmt"
)

func 只定义error() {
	var errNotFound error = errors.New("Not found error")
	fmt.Printf("error: %v", errNotFound)
}

func main() {
	只定义error()
}

package main

import (
	"errors"
	"fmt"
)

func 只定义error() {
	var errNotFound error = errors.New("Not found error")
	fmt.Printf("error: %v", errNotFound)
}

/**
终止程序
*/
func 运行时异常和panic() {
	fmt.Println("启动程序")
	panic("发生严重错误：停止程序！")
	fmt.Println("结束程序")
}

func badCall() {
	panic("坏结局")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall()
	fmt.Printf("打错电话后\r\n") // <-- wordt niet bereikt
}

/**
（recover）内建函数被用于从 panic 或 错误场景中恢复：
让程序可以从 panicking 重新获得控制权，停止终止过程进而恢复正常执行。

recover 只能在 defer 修饰的函数中使用：用于取得 panic 调用中传递过来的错误值，
如果是正常执行，调用 recover 会返回 nil，且没有其它效果。

总结：panic 会导致栈被展开直到 defer 修饰的 recover() 被调用或者程序中止。
*/
func 从panic中恢复() {
	fmt.Printf("通话测试\r\n")
	test()
	fmt.Printf("测试完成\r\n")
}

func main() {
	从panic中恢复()
}

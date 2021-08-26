package main

import (
	"fmt"
	"time"
)

func longWait() {
	fmt.Println("开始 longWait()")
	time.Sleep(5 * 1e9) // sleep for 5 seconds
	fmt.Println("结束 longWait()")
}

func shortWait() {
	fmt.Println("开始 shortWait()")
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("结束 shortWait()")
}

func go的并发并行和协程() {
	fmt.Println("在main（）")
	go longWait()
	go shortWait()
	fmt.Println("即将在 main() 中睡觉")
	// sleep works with a Duration in nanoseconds (ns) !
	time.Sleep(10 * 1e9)
	fmt.Println("在 main() 结束时")
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}

func 通道() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	// 等待了 1 秒让两个协程完成，如果不这样，sendData() 就没有机会输出。
	time.Sleep(1e9)
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func 通道阻塞() {
	ch1 := make(chan int)
	//pump() 函数为通道提供数值，也被叫做生产者。
	go pump(ch1)
	fmt.Println(<-ch1)
}

func main() {
	通道阻塞()
}

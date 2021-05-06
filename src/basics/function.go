package main

import "fmt"

/*
返回值一般用return返回
没有参数的函数通常被称为 niladic 函数（niladic function），就像 main.main()。
它的形参一般是有名字的，不过我们也可以定义没有形参名的函数，只有相应的形参类型，就像这样：func f(int, int, float64)。

按值传递（call by value） 按引用传递（call by reference）
Go 默认使用按值传递来传递参数，也就是传递参数的副本。函数接收参数副本之后，在使用变量的过程中可能对副本的值进行更改，但不会影响到原来的变量，比如 Function(arg1)。
如需改变原来的值可以传递指针
几乎在任何情况下，传递指针（一个32位或者64位的值）的消耗都比传递副本来得少。
在函数调用时，像切片（slice）、字典（map）、接口（interface）、通道（channel）这样的引用类型都是默认使用引用传递（即使没有显式的指出指针）。

返回值写在参数类型的后面比如:
func test(a int) int {}
返回值有命名返回和非命名返回
func getX2AndX3(input int) (int, int) {
    return 2 * input, 3 * input
}

func getX2AndX3_2(input int) (x2 int, x3 int) {
    x2 = 2 * input
    x3 = 3 * input
    // return x2, x3
    return
}

空白符
空白符使用_表示
不可以使用 函数结束过后会自动被回收
*/
func 函数参数与返回值() {

}

/*
如果函数的最后一个参数是采用 ...type 的形式，那么这个函数就可以处理一个变长的参数，这个长度可以为 0，这样的函数称为变参函数。
func myFunc(a, b, arg ...int) {}

变长参数类似于一个切片 可以使用range函数遍历

如果变长参数的参数类型都不一样，有两种解决办法：
1 使用结构：
定义一个结构类型，假设它叫 Options，用以存储所有可能的参数：

type Options struct {
	par1 type1,
	par2 type2,
	...
}

2 使用空接口：
如果一个变长参数的类型没有被指定，则可以使用默认的空接口 interface{}，这样就可以接受任何类型的参数。
该方案不仅可以用于长度未知的参数，还可以用于任何不确定类型的参数。一般而言我们会使用一个 for-range 循环以及 switch 结构对每个参数的类型进行判断：

func typecheck(..,..,values … interface{}) {
	for _, value := range values {
		switch v := value.(type) {
			case int: …
			case float: …
			case string: …
			case bool: …
			default: …
		}
	}
}
*/
func 传递变长参数() {

}

/*
关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数（为什么要在返回之后才执行这些语句？
因为 return 语句同样可以包含一些操作，而不是单纯地返回某个值）。

关键字 defer 的用法类似于面向对象编程语言 Java 和 C# 的 finally 语句块，它一般用于释放某些已分配的资源。
当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）

关键字 defer 允许我们进行一些函数执行完成后的收尾工作，例如：
1 关闭文件流
// open a file
defer file.Close()
2 解锁一个加锁的资源
mu.Lock()
defer mu.Unlock()
3 打印最终报告
printHeader()
defer printFooter()
4 关闭数据库链接
// open a database connection
defer disconnectFromDB()
*/
func defer和追踪() {

}

/*
名称	说明
close	用于管道通信
len、cap	len 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）；cap 是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）
new、make	new 和 make 均是用于分配内存：new 用于值类型和用户定义的类型，如自定义结构，make 用于内置引用类型（切片、map 和管道）。它们的用法就像是函数，但是将类型作为参数：new(type)、make(type)。new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针（详见第 10.1 节）。它也可以被用于基本类型：v := new(int)。make(T) 返回类型 T 的初始化之后的值，因此它比 new 进行更多的工作（详见第 7.2.3/4 节、第 8.1.1 节和第 14.2.1 节）new() 是一个函数，不要忘记它的括号
copy、append	用于复制和连接切片
panic、recover	两者均用于错误处理机制
print、println	底层打印函数（详见第 4.2 节），在部署环境中建议使用 fmt 包
complex、real imag	用于创建和操作复数（详见第 4.5.2.2 节）
*/
func 内置函数() {

}

/*
最经典的例子便是计算斐波那契数列，即前两个数为1，从第三个数开始每个数均为前两个数之和。

*/
func 递归函数() {
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func main() {
	递归函数()
}

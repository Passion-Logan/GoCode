package main

import (
	"fmt"
	"reflect"
	"runtime"
)

type struct1 struct {
	i1  int
	f1  float32
	str string
}

/*
结构体定义的一般方式如下：
type identifier struct {
    field1 type1
    field2 type2
    ...
}

type T struct {a, b int} 也是合法的语法，它更适用于简单的结构体。
结构体里的字段都有 名字，像 field1、field2 等，如果字段在代码中从来也不会被用到，那么可以命名它为 _。

使用 new 函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针：var t *T = new(T)，
如果需要可以把这条语句放在不同的行（比如定义是包范围的，但是分配却没有必要在开始就做）。
t := new(T)，变量 t 是一个指向 T的指针，此时结构体字段的值是它们所属类型的零值。

*/
func 结构体定义() {
	//ms := &struct1{10, 15.5, "Chris"}
	// 此时ms的类型是 *struct1
	// 或者：ms := struct1{10, 15.5, "Chris"}
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"

	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	fmt.Println(ms)
}

type File struct {
	fd   int    // 文件描述符
	name string // 文件名
}

/*
工厂方法创建结构体
*/
func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}

	return &File{fd, name}
}

/*
带标签的结构体
结构体中的字段除了有名字和类型外，还可以有一个可选的标签（tag）：它是一个附属于字段的字符串，可以是文档或其他的重要标记。
标签的内容不可以在一般的编程中使用，只有包 reflect 能获取它
*/
type TagType struct { // tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}

func 带标签的结构体() {
	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b      int
	c      float32
	int    // anonymous field
	innerS //anonymous field
}

/*
结构体可以包含一个或多个 匿名（或内嵌）字段，即这些字段没有显式的名字，只有字段的类型是必须的，此时类型就是字段的名字。
匿名字段本身可以是一个结构体类型，即 结构体可以包含内嵌结构体。

可以粗略地将这个和面向对象语言中的继承概念相比较，随后将会看到它被用来模拟类似继承的行为。Go 语言中的继承是通过内嵌或组合来实现的，
所以可以说，在 Go 语言中，相比较于继承，组合更受青睐。
*/
func 匿名字段() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:", outer2)
}

type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by float32
}

/*
同样地结构体也是一种数据类型，所以它也可以作为一个匿名字段来使用，
如同上面例子中那样。外层结构体通过 outer.in1 直接进入内层结构体的字段，
内嵌结构体甚至可以来自其他包。内层结构体被简单的插入或者内嵌进外层结构体。
这个简单的“继承”机制提供了一种方式，使得可以从另外一个或一些类型继承部分或全部实现。
*/
func 内嵌结构体() {
	b := B{A{1, 2}, 3.0, 4.0}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
}

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

/*
当定义了一个有很多方法的类型时，十之八九你会使用 String() 方法来定制类型的字符串形式的输出，
换句话说：一种可阅读性和打印性的输出。如果类型定义了 String() 方法，它会被用在 fmt.Printf() 中生成默认的输出：
等同于使用格式化描述符 %v 产生的输出。还有 fmt.Print() 和 fmt.Println() 也会自动使用 String() 方法。
*/
func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}

func 内嵌方法和继承() {
	c := new(Customer)
	c.Name = "Barak Obama"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"
	// shorter
	c = &Customer{"Barak Obama", &Log{"1 - Yes we can!"}}
	// fmt.Println(c) &{Barak Obama 1 - Yes we can!}
	c.Log().Add("2 - After me the world will be a better place!")
	//fmt.Println(c.log)
	fmt.Println(c.Log())
}

/*
Go 开发者不需要写代码来释放程序中不再使用的变量和结构占用的内存，在 Go 运行时中有一个独立的进程，
即垃圾收集器（GC），会处理这些事情，它搜索不再使用的变量然后释放它们的内存。
可以通过 runtime 包访问 GC 进程。

通过调用 runtime.GC() 函数可以显式的触发 GC，但这只在某些罕见的场景下才有用，比如当内存资源不足时调用 runtime.GC()，
它会在此函数执行的点上立即释放一大片内存，此时程序可能会有短时的性能下降（因为 GC 进程在执行）。

如果需要在一个对象 obj 被从内存移除前执行一些特殊操作，比如写到日志文件中，可以通过如下方式调用函数来实现：

runtime.SetFinalizer(obj, func(obj *typeObj))
func(obj *typeObj) 需要一个 typeObj 类型的指针参数 obj，特殊操作会在它上面执行。func 也可以是一个匿名函数。

在对象被 GC 进程选中并从内存中移除以前，SetFinalizer 都不会执行，即使程序正常结束或者发生错误。
*/
func 垃圾回收和SetFinalizer() {
	// fmt.Printf("%d\n", runtime.MemStats.Alloc/1024)
	// 此处代码在 Go 1.5.1下不再有效，更正为
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
}

func main() {
	垃圾回收和SetFinalizer()
}

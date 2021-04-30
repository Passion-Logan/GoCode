package main

import (
	"container/list"
	"fmt"
	"sort"
	"unsafe"
)

/*
-- 数组的声明语法：var 数组变量名 [元素数量]Type
数组变量名：数组声明及使用时的变量名。
元素数量：数组的元素数量，可以是一个表达式，但最终通过编译期计算的结果必须是整型数值，元素数量不能含有到运行时才能确认大小的数值。
Type：可以是任意基本类型，包括数组本身，类型为数组本身时，可以实现多维数组。
*/
func 数组() {
	var a [3]int             // 定义三个整数的数组
	fmt.Println(a[0])        // 打印第一个元素
	fmt.Println(a[len(a)-1]) // 打印最后一个元素
	// 打印索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	// 仅打印元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// 默认情况下，数组的每个元素都会被初始化为元素类型对应的零值，对于数字类型来说就是 0，同时也可以使用数组字面值语法，用一组值来初始化数组：
	// var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // "0"
	// 在数组的定义中，如果在数组长度的位置出现“...”省略号，则表示数组的长度是根据初始化值的个数来计算，因此，上面数组 q 的定义可以简化为：
	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"
	// 比较两个数组是否相等
	// 如果两个数组类型相同（包括数组的长度，数组中元素的类型）的情况下，我们可以直接通过较运算符（==和!=）来判断两个数组是否相等，只有当两个数组的所有元素都是相等的时候数组才是相等的，不能比较两个类型不同的数组，否则程序将无法完成编译。
	d := [2]int{1, 2}
	e := [...]int{1, 2}
	f := [2]int{1, 3}
	fmt.Println(d == e, d == f, e == f) // "true false false"
	// g := [3]int{1, 2}
	// fmt.Println(d == g) // 编译错误：无法比较 [2]int == [3]int

	// 遍历数组——访问每一个数组元素
	var team [3]string
	team[0] = "hammer"
	team[1] = "soldier"
	team[2] = "mum"
	for k, v := range team {
		fmt.Println(k, v)
	}
}

/*
-- 声明多维数组的语法：var array_name [size1][size2]...[sizen] array_type
	其中，array_name 为数组的名字，array_type 为数组的类型，size1、size2 等等为数组每一维度的长度。
*/
func 多维数组() {
	// 声明一个二维整型数组，两个维度的长度分别是 4 和 2
	var array [4][2]int
	// 使用数组字面量来声明并初始化一个二维整型数组
	array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	// 声明并初始化数组中索引为 1 和 3 的元素
	array = [4][2]int{1: {20, 21}, 3: {40, 41}}
	// 声明并初始化数组中指定的元素
	array = [4][2]int{1: {0: 20}, 3: {1: 41}}

	// 为了访问单个元素，需要反复组合使用[ ]方括号，如下所示
	fmt.Println(array)
	fmt.Println(array[1][0])
}

/*
-- 切片
切片（slice）是对数组的一个连续片段的引用，所以切片是一个引用类型（因此更类似于 C/C++ 中的数组类型，或者 Python 中的 list 类型），
这个片段可以是整个数组，也可以是由起始和终止索引标识的一些项的子集，需要注意的是，终止索引标识的项不包括在切片内。

Go语言中切片的内部结构包含地址、大小和容量，切片一般用于快速地操作一块数据集合，如果将数据集合比作切糕的话，
切片就是你要的“那一块”，切的过程包含从哪里开始（切片的起始位置）及切多大（切片的大小），容量可以理解为装切片的口袋大小。

从数组或切片生成新的切片
切片默认指向一段连续内存区域，可以是数组，也可以是切片本身。
从连续内存区域生成切片是常见的操作，格式如下：
slice [开始位置 : 结束位置]
*/
func 切片() {
	// 从数组生成切片
	var a = [3]int{1, 2, 3}
	// 其中 a 是一个拥有 3 个整型元素的数组，被初始化为数值 1 到 3，使用 a[1:2] 可以生成一个新的切片
	/*
		从数组或切片生成新的切片拥有如下特性：
		取出的元素数量为：结束位置 - 开始位置；
		取出元素不包含结束位置对应的索引，切片最后一个元素使用 slice[len(slice)] 获取；
		当缺省开始位置时，表示从连续区域开头到结束位置；
		当缺省结束位置时，表示从开始位置到整个连续区域末尾；
		两者同时缺省时，与切片本身等效；
		两者同时为 0 时，等效于空切片，一般用于切片复位。
	*/
	fmt.Println(a, a[1:2])

	// 从指定范围中生成切片
	var highRiseBuilding [30]int
	for i := 0; i < 30; i++ {
		highRiseBuilding[i] = i + 1
	}
	// 区间
	fmt.Println(highRiseBuilding[10:15])
	// 中间到尾部的所有元素
	fmt.Println(highRiseBuilding[20:])
	// 开头到中间指定位置的所有元素
	fmt.Println(highRiseBuilding[:2])
	// 表示原有的切片
	fmt.Println(a[:])
	// 重置切片，清空拥有的元素
	fmt.Println(a[0:0])

	// 直接声明新的切片:var name []Type
	// var strList []string

	// 使用 make() 函数构造切片:make( []Type, size, cap )
	// 其中 Type 是指切片的元素类型，size 指的是为这个类型分配多少个元素，cap 为预分配的元素数量，这个值设定后不影响 size，只是能提前分配空间，降低多次分配空间造成的性能问题。
	h := make([]int, 2)
	i := make([]int, 2, 10)
	fmt.Println(h, i)
	fmt.Println(len(h), len(i))

	// Go语言的内建函数 append() 可以为切片动态添加元素
	var j []int
	j = append(j, 1)                    // 追加1个元素
	j = append(j, 1, 2, 3)              // 追加多个元素, 手写解包方式
	j = append(j, []int{1, 2, 3}...)    // 追加一个切片, 切片需要解包
	j = append([]int{0}, j...)          // 在开头添加1个元素
	j = append([]int{-3, -2, -1}, j...) // 在开头添加1个切片
	// 因为 append 函数返回新切片的特性，所以切片也支持链式操作，我们可以将多个 append 操作组合起来，实现在切片中间插入元素：
	// var a []int
	// a = append(a[:i], append([]int{x}, a[i:]...)...) // 在第i个位置插入x
	// a = append(a[:i], append([]int{1,2,3}, a[i:]...)...) // 在第i个位置插入切片
	// 每个添加操作中的第二个 append 调用都会创建一个临时切片，并将 a[i:] 的内容复制到新创建的切片中，然后将临时创建的切片再追加到 a[:i] 中。
	fmt.Println(j)

	// 切片复制（切片拷贝）
	// Go语言的内置函数 copy() 可以将一个数组切片复制到另一个数组切片中，如果加入的两个数组切片不一样大，就会按照其中较小的那个数组切片的元素个数进行复制。
	// copy() 函数的使用格式如下： copy( destSlice, srcSlice []T) int
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中 [1 2 3 4 5] [1 2 3]
	//copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置 [5 4 3 4 5] [5 4 3]
	fmt.Println(slice1, slice2)
	// 虽然通过循环复制切片元素更直接，不过内置的 copy() 函数使用起来更加方便，copy() 函数的第一个参数是要复制的目标 slice，第二个参数是源 slice，两个 slice 可以共享同一个底层数组，甚至有重叠也没有问题。
	/*
		切片引用复制导致的数据改变
		//切片()
			// 设置元素数量为1000
			const elementCount = 1000
			// 预分配足够多的元素切片
			srcData := make([]int, elementCount)
			// 将切片赋值
			for i := 0; i < elementCount; i++ {
				srcData[i] = i
			}
			// 引用切片数据
			refData := srcData
			// 预分配足够多的元素切片
			copyData := make([]int, elementCount)
			// 将数据复制到新的切片空间中
			copy(copyData, srcData)
			// 修改原始数据的第一个元素
			srcData[0] = 999
			// 打印引用切片的第一个元素
			fmt.Println(refData[0])
			// 打印复制切片的第一个和最后一个元素
			fmt.Println(copyData[0], copyData[elementCount-1])
			// 复制原始数据从4到6(不包含)
			copy(copyData, srcData[4:6])
			for i := 0; i < 5; i++ {
				fmt.Printf("%d ", copyData[i])
			}
	*/

	// 从切片中删除元素
	// Go语言并没有对删除切片元素提供专用的语法或者接口，需要使用切片本身的特性来删除元素，根据要删除元素的位置有三种情况，分别是从开头位置删除、从中间位置删除和从尾部删除，其中删除切片尾部的元素速度最快。
	// 从开头位置删除
	// copy1 := []int{1, 2, 3, 4, 5}
	// copy1 = copy1[1:] // 删除开头1个元素
	// copy1 = copy1[N:] // 删除开头N个元素
	// 也可以不移动数据指针，但是将后面的数据向开头移动，可以用 append 原地完成（所谓原地完成是指在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化）：
	// copy1 append(copy1[:0], copy1[1:]...) // 删除开头1个元素
	// copy1 = append(copy1[:0], copy1[N:]...) // 删除开头N个元素
	// 还可以用 copy() 函数来删除开头的元素：
	// copy1 = copy1[:copy(copy1, copy1[1:])] // 删除开头1个元素
	// copy1 = copy1[:copy(copy1, copy1[N:])] // 删除开头N个元素

	// 从中间位置删除
	// copy1 = append(copy1[:i], copy1[i+1:]...) // 删除中间1个元素
	// copy1 = append(copy1[:i], copy1[i+N:]...) // 删除中间N个元素
	// copy1 = copy1[:i+copy(a[i:], copy1[i+1:])] // 删除中间1个元素
	// copy1 = copy1[:i+copy(a[i:], copy1[i+N:])] // 删除中间N个元素
	/*
		test := []int{1, 2, 3, 4, 5}
		test = append(test[:2], test[3:]...)
		fmt.Print(test)
	*/
	// 从尾部删除
	// seq := []string{"a", "b", "c", "d", "e"}
	// 查看删除位置之前的元素和之后的元素
	// fmt.Println(seq[:len(seq) - 1]) // 删除尾部1个元素
	// fmt.Println(seq[:len(seq) - N]) // 删除尾部N个元素

	// 循环迭代切片
	// 创建一个整型切片，并赋值
	slice := []int{10, 20, 30, 40}
	// 迭代每一个元素，并显示其值
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}
	// 需要强调的是，range 返回的是每个元素的副本，而不是直接返回对该元素的引用
	// 关键字 range 总是会从切片头部开始迭代
}

/*
声明一个多维数组的语法格式如下：var sliceName [][]...[]sliceType
其中，sliceName 为切片的名字，sliceType为切片的类型，每个[]代表着一个维度，切片有几个维度就需要几个[]。
*/
func 多维切片() {
	// 声明一个二维切片
	// var slice [][]int
	// 为二维切片赋值
	// slice = [][]int{{10}, {100, 200}}
	// 声明一个二维整型切片并赋值
	// slice := [][]int{{10}, {100, 200}}
	// 追加值 slice[0] = append(slice[0], 20)
	// [[10 20] [100 200]]
}

/*
Go语言中 map 是一种特殊的数据结构，一种元素对（pair）的无序集合
声明语法：var mapname map[keytype]valuetype
其中：
mapname 为 map 的变量名。
keytype 为键类型。
valuetype 是键对应的值类型。
提示：[keytype] 和 valuetype 之间允许有空格。
在声明的时候不需要知道 map 的长度，因为 map 是可以动态增长的，未初始化的 map 的值是 nil，使用函数 len() 可以获取 map 中 pair 的数目。
*/
func 字典() {
	var mapLit map[string]int
	//var mapCreated map[string]float32
	var mapAssigned map[string]int
	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3
	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])

	// map 容量 make(map[keytype]valuetype, cap)
	// map2 := make(map[string]float, 100)
	//noteFrequency := map[string]float32 {
	//	"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
	//	"G0": 24.50, "A0": 27.50, "B0": 30.87, "A4": 440}

	// 用切片作为 map 的值
	// mp1 := make(map[int][]int)
	// mp2 := make(map[int]*[]int)

	// 遍历map
	// map 的遍历过程使用 for range 循环完成，代码如下：
	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	for k, v := range scene {
		fmt.Println(k, v)
	}
	// 如果需要特定顺序的遍历结果，正确的做法是先排序，代码如下：
	// 声明一个切片保存map数据
	var sceneList []string
	// 将map数据遍历复制到切片中
	for k := range scene {
		sceneList = append(sceneList, k)
	}
	// 对切片进行排序
	sort.Strings(sceneList)
	// 输出
	fmt.Println(sceneList)
	// [brazil china route]

	// map元素的删除和清空
	// 使用 delete() 内建函数从 map 中删除一组键值对，delete() 函数的格式如下：delete(map, 键)
	// 清空 map 中的所有元素
	// Go语言中并没有为 map 提供任何清空所有元素的函数、方法，清空 map 的唯一办法就是重新 make 一个新的 map，不用担心垃圾回收的效率，Go语言中的并行垃圾回收效率比写一个清空函数要高效的多。

	// map的并发
	/*
		var scene sync.Map
		// 将键值对保存到sync.Map
		scene.Store("greece", 97)
		scene.Store("london", 100)
		scene.Store("egypt", 200)
		// 从sync.Map中根据键取值
		fmt.Println(scene.Load("london"))
		// 根据键删除对应的键值对
		scene.Delete("london")
		// 遍历所有sync.Map中的键值对
		scene.Range(func(k, v interface{}) bool {
			fmt.Println("iterate:", k, v)
			return true
		})
	*/
}

/*
在Go语言中，列表使用 container/list 包来实现，内部的实现原理是双链表，列表能够高效地进行任意位置的元素插入和删除操作。
list 的初始化有两种方法：分别是使用 New() 函数和 var 关键字声明，两种方法的初始化效果都是一致的。
变量名 := list.New()
var 变量名 list.List
双链表支持从队列前方或后方插入元素，分别对应的方法是 PushFront 和 PushBack。
*/
func 列表() {
	// 列表()
	l := list.New()
	// 尾部添加
	l.PushBack("canon")
	// 头部添加
	l.PushFront(67)
	// 尾部添加后保存元素句柄
	element := l.PushBack("fist")
	// 在fist之后添加high
	l.InsertAfter("high", element)
	// 在fist之前添加noon
	l.InsertBefore("noon", element)
	// 使用
	l.Remove(element)

	// 遍历列表——访问列表的每一个元素
	// 遍历双链表需要配合 Front() 函数获取头元素，遍历时只要元素不为空就可以继续进行，每一次遍历都会调用元素的 Next() 函数，代码如下所示。
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

/*
在Go语言中，布尔类型的零值（初始值）为 false，数值类型的零值为 0，字符串类型的零值为空字符串""，而指针、切片、映射、通道、函数和接口的零值则是 nil。
nil 是Go语言中一个预定义好的标识符，有过其他编程语言开发经验的开发者也许会把 nil 看作其他语言中的 null（NULL），
其实这并不是完全正确的，因为Go语言中的 nil 和其他语言中的 null 有很多不同点。

nil 标识符是不能比较的
nil 不是关键字或保留字
nil 没有默认类型
不同类型 nil 的指针是一样的
不同类型的 nil 是不能比较的
两个相同类型的 nil 值也可能无法比较
nil 是 map、slice、pointer、channel、func、interface 的零值
不同类型的 nil 值占用的内存大小可能是不一样的
(具体的大小取决于编译器和架构,对应 32 位的架构的，打印的大小将减半)
*/
func 空值() {
	var p *struct{}
	fmt.Println(unsafe.Sizeof(p)) // 8
	var s []int
	fmt.Println(unsafe.Sizeof(s)) // 24
	var m map[int]bool
	fmt.Println(unsafe.Sizeof(m)) // 8
	var c chan string
	fmt.Println(unsafe.Sizeof(c)) // 8
	var f func()
	fmt.Println(unsafe.Sizeof(f)) // 8
	var i interface{}
	fmt.Println(unsafe.Sizeof(i)) // 16
}

func main() {
	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}

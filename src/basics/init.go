package main

import (
	"fmt"
	"math"
)

/**
-- 变量的声明
Go语言的基本类型有：
bool
string
int、int8、int16、int32、int64
uint、uint8、uint16、uint32、uint64、uintptr
byte // uint8 的别名
rune // int32 的别名 代表一个 Unicode 码
float32、float64
complex64、complex128
*/
func Statement() {
	// 标准格式
	// 声明变量 var name type
	var a string = "123"

	// 简短格式  名字 := 表达式
	// i, j := 0, 1

	// 批量格式
	/*
		var (
			a int
			b string
			c []float32
			d func() bool
			e struct {
				x int
			}
		)*/

	fmt.Print(a)
}

/**
-- 变量的初始化
整型和浮点型变量的默认值为 0 和 0.0。
字符串变量的默认值为空字符串。
布尔型变量默认为 bool。
切片、函数、指针变量的默认为 nil。
*/
func Init() {
	// 标准格式 var 变量名 类型 = 表达式
	var hp int = 100
	var test int = 200

	// 	编译器推导格式
	// var hp = 100

	// 短变量声明并初始化  重复声明会报错
	// hp := 100

	// 简单赋值
	hp = 300

	// 多重赋值
	// hp, test = test, hp
	hp, test = test, hp

	// fmt.Println(hp, test)
}

/**
-- 匿名变量
匿名变量就是一个下划线"_",它本身就是一个特殊的标识符，被称为空白标识符。
它可以像其他标识符那样用于变量的声明或赋值（任何类型都可以赋值给它），但任何赋给这个标识符的值都将被抛弃，
因此这些值不能在后续的代码中使用，也不可以使用这个标识符作为变量对其它变量进行赋值或运算。
使用匿名变量时，只需要在变量声明的地方使用下画线替换即可。

匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。
*/
func Anonymous() (int, int) {
	/*
		a, _ := Anonymous()
		_, b := Anonymous()

		fmt.Println(a, b)
	*/
	return 100, 200
}

// 全局变量
// 全局变量声明必须以 var 关键字开头，如果想要在外部包中使用全局变量的首字母必须大写。
var c int

/**
-- 变量的作用域
函数内定义的变量称为局部变量
函数外定义的变量称为全局变量
函数定义中的变量称为形式参数
*/
func Scope(a, b int) int {
	// a,b是形式参数
	// 局部变量
	result := a + b

	return result
}

/**
-- 整数类型
数值类型分为以下几种：整数、浮点数、复数，其中每一种都包含了不同大小的数值类型，
例如有符号整数包含 int8、int16、int32、int64 等，
每种数值类型都决定了对应的大小范围和是否支持正负符号。
此外还有两种整数类型 int 和 uint，它们分别对应特定 CPU 平台的字长（机器字大小），
其中 int 表示有符号整数，应用最为广泛，uint 表示无符号整数。
实际开发中由于编译器和计算机硬件的不同，int 和 uint 所能表示的整数大小会在 32bit 或 64bit 之间变化。
大多数情况下，我们只需要 int 一种整型即可，它可以用于循环计数器（for 循环中控制循环次数的变量）、数组和切片的索引，以及任何通用目的的整型运算符，通常 int 类型的处理速度也是最快的。

用来表示 Unicode 字符的 rune 类型和 int32 类型是等价的，通常用于表示一个 Unicode 码点。
这两个名称可以互换使用。同样，byte 和 uint8 也是等价类型，byte 类型一般用于强调数值是一个原始的数据而不是一个小的整数。

最后，还有一种无符号的整数类型 uintptr，它没有指定具体的 bit 大小但是足以容纳指针。
uintptr 类型只有在底层编程时才需要，特别是Go语言和C语言函数库或操作系统接口相交互的地方。

尽管在某些特定的运行环境下 int、uint 和 uintptr 的大小可能相等，但是它们依然是不同的类型，比如 int 和 int32，
虽然 int 类型的大小也可能是 32 bit，但是在需要把 int 类型当做 int32 类型使用的时候必须显示的对类型进行转换，反之亦然。

Go语言中有符号整数采用 2 的补码形式表示，也就是最高 bit 位用来表示符号位，一个 n-bit 的有符号数的取值范围是从 -2(n-1) 到 2(n-1)-1。
无符号整数的所有 bit 位都用于表示非负数，取值范围是 0 到 2n-1。例如，int8 类型整数的取值范围是从 -128 到 127，而 uint8 类型整数的取值范围是从 0 到 255。
*/
func Integer() {
	//哪些情况下使用 int 和 uint
	//程序逻辑对整型范围没有特殊需求。例如，对象的长度使用内建 len() 函数返回，这个长度可以根据不同平台的字节长度进行变化。实际使用中，切片或 map 的元素数量等都可以用 int 来表示。
	//反之，在二进制传输、读写文件的结构描述时，为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使用 int 和 uint。
}

/**
-- 小数类型
常量 math.MaxFloat32 表示 float32 能取到的最大数值，大约是 3.4e38；
常量 math.MaxFloat64 表示 float64 能取到的最大数值，大约是 1.8e308；
float32 和 float64 能表示的最小值分别为 1.4e-45 和 4.9e-324。
*/
func Decimal() {
	// 一个 float32 类型的浮点数可以提供大约 6 个十进制数的精度，
	// 而 float64 则可以提供约 15 个十进制数的精度，
	// 通常应该优先使用 float64 类型，因为 float32 类型的累计计算误差很容易扩散，并且 float32 能精确表示的正整数并不是很大。
	// var f float32 = 16777216 // 1 << 24
	// fmt.Println(f == f + 1)

	// 浮点数在声明的时候可以只写整数部分或者小数部分，像下面这样：
	// const e = .71828 // 0.71828
	// const f = 1.     // 1

	// 很小或很大的数最好用科学计数法书写，通过 e 或 E 来指定指数部分：
	const Avogadro = 6.02214129e23 // 阿伏伽德罗常数
	const Planck = 6.62606957e-34  // 普朗克常数

	// 用 Printf 函数打印浮点数时可以使用“%f”来控制保留几位小数
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)
}

/**
-- 复数类型
在计算机中，复数是由两个浮点数表示的，其中一个表示实部（real），一个表示虚部（imag）。
Go语言中复数的类型有两种，分别是 complex128（64 位实数和虚数）和 complex64（32 位实数和虚数），其中 complex128 为复数的默认类型。
复数的值由三部分组成 RE + IMi，其中 RE 是实数部分，IM 是虚数部分，RE 和 IM 均为 float 类型，而最后的 i 是虚数单位
*/
func Complex() {

	// 语法:其中 name 为复数的变量名，complex128 为复数的类型，“=”后面的 complex 为Go语言的内置函数用于为复数赋值，
	// x、y 分别表示构成该复数的两个 float64 类型的数值，x 为实部，y 为虚部。
	// var name complex128 = complex(x, y)
	// name := complex(x, y)

	// 可以通过Go语言的内置函数real(z) 来获得该复数的实部，也就是 x；通过imag(z) 获得该复数的虚部，也就是 y。

	// var x complex128 = complex(1, 2) // 1+2i
	// var y complex128 = complex(3, 4) // 3+4i
	// fmt.Println(x*y)                 // "(-5+10i)"
	// fmt.Println(real(x*y))           // "-5"
	// fmt.Println(imag(x*y))           // "10"
}

/**
-- 布尔类型
一个布尔类型的值只有两种：true 或 false。if 和 for 语句的条件部分都是布尔类型的值，并且==和<等比较操作也会产生布尔型的值。
*/
func Boolean() {
	// 默认是false
	var temp bool
	fmt.Println(temp)
}

/**
-- 字符串类型
一个字符串是一个不可改变的字节序列，字符串可以包含任意的数据，但是通常是用来包含可读的文本，
字符串是 UTF-8 字符的一个序列（当字符为 ASCII 码表上的字符时则占用 1 个字节，其它字符根据需要占用 2-4 个字节）。
*/
func String() {
	hello := "hello"

	//定义字符串
	/**
	\n：换行符
	\r：回车符
	\t：tab 键
	\u 或 \U：Unicode 字符
	\\：反斜杠自身
	*/
	// fmt.Print("hello\nworld")

	// 字符串长度
	fmt.Println(len(hello))
	// 字符串的拼接 +
	hello += " world"
	fmt.Println(hello)

	// 多行字符串
	more := `
			第一行
			第二行
			第三行
			`

	fmt.Println(more)
}

/**
-- 字符类型
一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
另一种是 rune 类型，代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型。rune 类型等价于 int32 类型。
byte 类型是 uint8 的别名，对于只占用 1 个字节的传统 ASCII 编码的字符来说，完全没有问题，例如 var ch byte = 'A'，字符使用单引号括起来。

Go语言同样支持 Unicode（UTF-8），因此字符同样称为 Unicode 代码点或者 runes，并在内存中使用 int 来表示。在文档中，一般使用格式 U+hhhh 来表示，其中 h 表示一个 16 进制数。
*/
func Character() {
	// Unicode 包中内置了一些用于测试字符的函数，这些函数的返回值都是一个布尔值，如下所示（其中 ch 代表字符）：
	/**
	判断是否为字母：unicode.IsLetter(ch)
	判断是否为数字：unicode.IsDigit(ch)
	判断是否为空白符号：unicode.IsSpace(ch)
	*/
}

/**
-- 数据类型装换
Go语言不存在隐式类型转换，因此所有的类型转换都必须显式的声明
valueOfTypeB = typeB(valueOfTypeA)
*/
func TypeTo() {
	// 输出各数值范围
	fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16 range:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32 range:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64 range:", math.MinInt64, math.MaxInt64)
	// 初始化一个32位整型值
	var a int32 = 1047483647
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int32: 0x%x %d\n", a, a)
	// 将a变量数值转换为十六进制, 发生数值截断
	b := int16(a)
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int16: 0x%x %d\n", b, b)
	// 将常量保存为float32类型
	var c float32 = math.Pi
	// 转换为int类型, 浮点发生精度丢失
	fmt.Println(int(c))
}

/**
-- 指针类型
类型指针，允许对这个指针类型的数据进行修改，传递数据可以直接使用指针，而无须拷贝数据，类型指针不能进行偏移和运算。
切片，由指向起始元素的原始指针、元素数量和容量组成。

Go语言中使用在变量名前面添加&操作符（前缀）来获取变量的内存地址（取地址操作）
ptr := &v    // v 的类型为 T
其中 v 代表被取地址的变量，变量 v 的地址使用变量 ptr 进行接收，ptr 的类型为*T，称做 T 的指针类型，*代表指针

*/
func Pointer() {
	/*var cat int = 1
	var str string = "banana"
	fmt.Printf("%p %p", &cat, &str)*/

	// 从指针获取指针指向的值
	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"
	// 对字符串取地址, ptr类型为*string
	ptr := &house
	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)
	// 打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr)
	// 对指针进行取值操作
	value := *ptr
	// 取值后的类型
	fmt.Printf("value type: %T\n", value)
	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)

	// 使用指针修改值
	// 准备两个变量, 赋值1和2
	x, y := 1, 2
	// 交换变量值
	swap(&x, &y)
	// 输出变量值
	fmt.Println(x, y)

	// 创建指针的另一种方法——new() 函数 new(类型)
	// new() 函数可以创建一个对应类型的指针，创建过程会分配内存，被创建的指针指向默认值。
	str := new(string)
	*str = "Go语言教程"
	fmt.Println(*str)
}

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func dummy(b int) int {
	// 声明一个变量c并赋值
	var c int
	c = b
	return c
}

// 空函数, 什么也不做
func void() {
	// go run -gcflags "-m -l" ***.go
}

/**
-- 生命周期
全局变量：它的生命周期和整个程序的运行周期是一致的；
局部变量：它的生命周期则是动态的，从创建这个变量的声明语句开始，到这个变量不再被引用为止；
形式参数和函数返回值：它们都属于局部变量，在函数被调用的时候创建，函数调用结束后被销毁。
*/
func Lifecycle() {
}

/**
-- 常量
常量使用关键字 const 定义，用于存储不会改变的数据，常量是在编译时被创建的，即使定义在函数内部也是如此，
并且只能是布尔型、数字型（整数型、浮点型和复数）和字符串型。
由于编译时的限制，定义常量的表达式必须为能被编译器求值的常量表达式。
*/
func Constant() {
	// 显式类型定义 const b string = "abc"
	// 隐式类型定义： const b = "abc"
	// 常量的值必须是能够在编译时就能够确定的，可以在其赋值表达式中涉及计算过程，但是所有用于计算的值必须在编译期间就能获得。
	// 正确的做法：const c1 = 2/3
	// 错误的做法：const c2 = getNumber() // 引发构建错误: getNumber() 用做值
	/**
	const (
	    e  = 2.7182818
	    pi = 3.1415926
	)
	*/

	// iota 常量生成器
	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	// 首先定义一个 Weekday 命名类型，然后为一周的每天定义了一个常量，从周日 0 开始。在其它编程语言中，这种类型一般被称为枚举类型。
	// 周日将对应 0，周一为 1，以此类推。
}

/**
-- 类型别名使用type关键字
Go 1.9 版本之前定义内建类型的代码是这样写的
type byte uint8
type rune int32
而在 Go 1.9 版本之后变为：
type byte = uint8
type rune = int32
*/
func Typealias() {
	// 区分类型别名与类型定义
	/**
	定义类型别名的写法为：
	type TypeAlias = Type
	类型别名规定：TypeAlias 只是 Type 的别名，本质上 TypeAlias 与 Type 是同一个类型，就像一个孩子小时候有小名、乳名，上学后用学名，英语老师又会给他起英文名，但这些名字都指的是他本人。
	*/
	// 将NewInt定义为int类型
	type NewInt int
	// 将int取一个别名叫IntAlias
	type IntAlias = int

	// 将a声明为NewInt类型
	var a NewInt
	// 查看a的类型名
	fmt.Printf("a type: %T\n", a)
	// 将a2声明为IntAlias类型
	var a2 IntAlias
	// 查看a2的类型名
	fmt.Printf("a2 type: %T\n", a2)

	// 非本地类型不能定义方法
	/**
	// 定义time.Duration的别名为MyDuration
	type MyDuration = time.Duration
	// 为MyDuration添加一个函数
	func (m MyDuration) EasySet(a string) {
	}
	编译上面代码报错，信息如下：
	cannot define new methods on non-local type time.Duration
	解决这个问题有下面两种方法：
	1 将第 8 行修改为 type MyDuration time.Duration，也就是将 MyDuration 从别名改为类型；
	2 将 MyDuration 的别名定义放在 time 包中。
	*/
}

func main() {
	Typealias()
}

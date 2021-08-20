package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func 读取用户的输入() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
	// 输出结果: From the string we read: 56.12 5212 Go
}

func 缓冲读取() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}

func 文件读取() {
	path, _ := filepath.Abs("")
	path += "\\src\\test.txt"

	fmt.Println(path)

	inputFile, inputError := os.Open(path)
	if inputError != nil {
		fmt.Printf("文件读取失败!")
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		// 逐行读取
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("读取到的内容：%s", inputString)
		if readerError == io.EOF {
			return
		}
	}
}

func 内容复制() {
	inputFile := "./src/test.txt"
	outputFile := "./src/copy.txt"

	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "读取失败：%s", err)
	}
	fmt.Printf("%s", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644)
	if err != nil {
		panic(err.Error())
	}
}

/**
如果数据是按列排列并用空格分隔的，你可以使用 fmt 包提供的以 FScan 开头的一系列函数来读取他们。
将 3 列的数据分别读入变量 v1、v2 和 v3 内，然后分别把他们添加到切片的尾部。
*/
func 按列读取() {
	file, err := os.Open("./src/test2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}

	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}

func 读取压缩文件() {
	fName := "./src/test.zip"
	var r *bufio.Reader
	fi, err := os.Open(fName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, 打开 %s: 错误: %s\n", os.Args[0], fName,
			err)
		os.Exit(1)
	}

	defer fi.Close()

	fz, err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("读完文件")
			os.Exit(0)
		}
		fmt.Println(line)
	}

}

func main() {
	读取压缩文件()
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	t := 1
	fmt.Println('a' + t)
	//s := ""
	//fmt.Scan(&s)
	//fmt.Scanln(&s)
	//fmt.Scanf("%s", &s)
	//
	// 输入带空格的字符串
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(input)

}

// Test go的循环输入
func Test() {
	// 记得引入头文件  "io"
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err == io.EOF {
			break
		}
	}
}

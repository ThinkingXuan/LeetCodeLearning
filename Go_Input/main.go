package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	inputInt()     // 输入整形
	inputString()  // 输入字符串
	inputDynamic() // 整数和按行输入

	inputLoop() // 循环输入

	// 读入大量的整数，go的fmt.Scan()消耗很大。
	//使用自定义ReadInt输出
}

// 整形输入
func inputInt() {
	// 输入格式 1 2
	// 输入格式:   1
	//            2
	// 都可以
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(a, b)

	// fmt.Scanln() 和fmt.Scan()区别
	//  比如 fmt.Scan(&x, &y, &z)    x y在一行输入， z换行输入。。是可行的。
	// fmt.Scanln(&x, &y, &z)       x y在一行输入， z换行输入。 是不可行的。遇到换行就会结束输入。
}

// 字符串输入
func inputString() {
	// 不带空格string
	s := ""
	fmt.Scan(&s)
	fmt.Println(s)

	// 带空格string, 也就是按行读取。
	inputRead := bufio.NewReader(os.Stdin)
	for i := 0; i < 5; i++ {
		// 带空格string, 也就是按行读取。注意line是[]byte类型.
		line, _, _ := inputRead.ReadLine()
		fmt.Println(string(line))
	}
	// 读带空格string,  但是末尾会多一个换行符

	line, _ := inputRead.ReadString('\n')
	fmt.Println(line)
}

// 混合输入，fmt.Scan和ReadLine()混合  有大坑。
func inputDynamic() {

	inputRead := bufio.NewReader(os.Stdin)

	// 方式 1
	// 如果用scan输入。  后面会多一个换行符。所以要 inputRead.ReadLine() 下将换行符读取。。 然后才能输入真正内容。
	var a, b int
	fmt.Scan(&a, &b)

	inputRead.ReadLine()

	line, _, _ := inputRead.ReadLine()
	fmt.Println(string(line))

	// 方式 2
	fmt.Scanln(&a, &b)
	line, _, _ = inputRead.ReadLine()
	fmt.Println(string(line))

}

func inputLoop() {
	// 记得引入头文件  "io"
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err == io.EOF {
			break
		}
	}
}

// 高效率输入
var in = bufio.NewReader(os.Stdin)

func readInt(out *int) error {
	var ans, sign = 0, 1
	var readed = false
	c, err := in.ReadByte()
	for ; err == nil && (c < '0' || '9' < c); c, err = in.ReadByte() {
		if c == '-' {
			sign = -sign
		}
	}
	for ; err == nil && '0' <= c && c <= '9'; c, err = in.ReadByte() {
		ans = ans<<3 + ans<<1 + int(c-'0')
		readed = true
	}
	if readed {
		*out = ans * sign
		return nil
	}
	return err
}

func readInt64(out *int64) error {
	var ans, sign int64 = 0, 1
	var readed = false
	c, err := in.ReadByte()
	for ; err == nil && (c < '0' || '9' < c); c, err = in.ReadByte() {
		if c == '-' {
			sign = -sign
		}
	}
	for ; err == nil && '0' <= c && c <= '9'; c, err = in.ReadByte() {
		ans = ans<<3 + ans<<1 + int64(c-'0')
		readed = true
	}
	if readed {
		*out = ans * sign
		return nil
	}
	return err
}

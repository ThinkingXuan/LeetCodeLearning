package main

import (
	"fmt"
	"strings"
)

//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。



func main() {

	fmt.Println(replaceSpace("We are happy."))
}

func replaceSpace(s string) string {
	return strings.Replace(s," ", "%20",-1)
}
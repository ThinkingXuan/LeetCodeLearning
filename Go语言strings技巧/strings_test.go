package Go语言strings技巧

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringsMap(t *testing.T) {
	// Strings.Map语言对按照规则对字符串字符进行处理
	s := "Hello World!"
	s2 := strings.Map(func(r rune) rune {
		if r == ' ' {
			// 返回小于0得书就是删除
			return -1
		}
		return r
	}, s)
	fmt.Println(s2)
}

func TestScan(t *testing.T) {

}

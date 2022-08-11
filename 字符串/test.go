package main

import (
	"fmt"
	"strings"
)

type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("print：%v", p)
}

func main() {
	s := "abcdefg"
	fmt.Println(strings.Count(s, ""))

	fmt.Println(fmt.Sprintf("%b", 10))

	p := &People{}
	p.String()
}

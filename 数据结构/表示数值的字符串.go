package main

//剑指 Offer 20.
//请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。

type State int
type CharType int

const (
	stateInitial State = iota
	stateIntSign
	stateInteger
	statePoint
	statePointWithoutInt
	stateFraction
	stateExp
	stateExpSign
	stateExpNumber
	stateEND
)

const (
	charNumber CharType = iota
	charExp
	charPoint
	charSign
	charSpace
	charIllegal
)

func toCharType(ch byte) CharType {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return charNumber
	case 'e', 'E':
		return charExp
	case '.':
		return charPoint
	case '+', '-':
		return charSign
	case ' ':
		return charSpace
	default:
		return charIllegal
	}
}

func main() {

}

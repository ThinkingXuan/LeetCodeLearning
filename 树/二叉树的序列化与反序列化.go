package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Codec struct {
}

func Constructor() Codec {
	code := Codec{}
	return code
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "X"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

func buildTree(data *[]string) *TreeNode {
	rootVal := (*data)[0]
	*data = (*data)[1:]
	if rootVal == "X" {
		return nil
	}

	Val, _ := strconv.Atoi(rootVal)

	root := &TreeNode{Val: Val}
	root.Left = buildTree(data)
	root.Right = buildTree(data)
	return root

}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	list := strings.Split(data, ",")
	return buildTree(&list)
}

func main() {
	code := Constructor()
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}

	fmt.Printf(code.serialize(tree))
}

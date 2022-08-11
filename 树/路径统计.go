package main

import (
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {

	var res []string

	var helper func(root *TreeNode, s string)
	helper = func(root *TreeNode, s string) {
		if root == nil {
			res = append(res, s[1:])
		}
		newS := s + "->" + strconv.Itoa(root.Val)
		helper(root.Left, newS)
		helper(root.Right, newS)
	}
	helper(root, "")
	return res
}

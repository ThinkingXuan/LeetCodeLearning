package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1==nil && root2== nil {
		return nil
	}

	root := &TreeNode{}

	if root1!=nil && root2!=nil {
		root.Val = root1.Val+root2.Val
	}
	if root1!=nil && root2==nil {
		root.Val = root1.Val
	}
	if root1==nil && root2!=nil {
		root.Val = root2.Val
	}

	leftNode := mergeTrees(root1.Left, root2.Left)
	rightNode :=  mergeTrees(root1.Right, root2.Right)
	if leftNode!=nil {
		root.Left = leftNode
	}
	if rightNode!=nil {
		root.Right = rightNode
	}
	return root
}

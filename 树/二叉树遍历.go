package main

// 二叉树的非递归遍历的通用方式

func main() {

}

func preOrder(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int

	if root != nil {
		stack = append(stack, root)
	} else {
		return []int{}
	}

	for len(stack) > 0 {
		node := stack[len(stack)-1]

		if node != nil {
			stack = stack[:len(stack)-1] // 将该节点弹出，避免重复操作。
			//先序遍历  中左右，层数遍历是道过来入栈
			if node.Right != nil { // 右
				stack = append(stack, node.Right)
			}
			if node.Left != nil { // 左
				stack = append(stack, node.Left)
			}
			stack = append(stack, node) // 中
			stack = append(stack, nil)

		} else {
			stack = stack[:len(stack)-1] // 将空节点弹出

			node = stack[len(stack)-1] // 重新去栈中元素
			stack = stack[:len(stack)-1]
			res = append(res, node.Val) // 加入结果集
		}
	}
	return res
}

func inOrder(root *TreeNode) {
	// 中序遍历把入栈顺序调整下
	// 左中右  =>        右中左
}

func postOrder(root *TreeNode) {
	// 后序遍历把入栈顺序调整下
	// 左右中 =>             中左右

}

// 二叉树的层次遍历（非递归）
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root)
	} else {
		return [][]int{}
	}
	for len(queue) > 0 {
		var levelRes []int
		length := len(queue)

		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			levelRes = append(levelRes, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, levelRes)
	}

	return res
}

// 二叉树的层次遍历（递归）
func levelOrder2(root *TreeNode) [][]int {
	var res [][]int
	var helper func(root *TreeNode, level int)

	helper = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(res) == level {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)

		helper(root.Left, level+1)
		helper(root.Right, level+1)
	}
	helper(root, 0)

	return res
}

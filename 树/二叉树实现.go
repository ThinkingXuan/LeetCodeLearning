package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Tree 定义二叉树
type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

//  traverse 二叉树中序遍历
func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Println(t.Value, " ")
	traverse(t.Right)
}

// create 创建二叉树
func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		tmp := rand.Intn(n * 2)
		t = insert(t, tmp)
	}

	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{v, nil, nil}
	}
	if v == t.Value {
		return t
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

// 由数组创建二叉树，记住很重要。(递归写法)
func createTree(nums []int, i int) *TreeNode {
	if i >= len(nums) || nums[i] == -1 {
		return nil
	}

	root := &TreeNode{Val: nums[i]}
	root.Left = createTree(nums, i*2+1)
	root.Left = createTree(nums, i*2+2)
	return root
}

// 二叉树搜索数的插入
func insertBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		node := &TreeNode{Val: val}
		return node
	}
	if root.Val > val {
		root.Left = insertBST(root.Left, val)
	} else if root.Val < val {
		root.Right = insertBST(root.Right, val)
	}
	return root
}

// 二叉搜索树的删除
func deleteBST(root *TreeNode, key int) *TreeNode {
	// 1. 没有找到要删除的节点，遍历到空节点直接返回了
	if root == nil {
		return root
	}
	if root.Val > key {
		root.Left = deleteBST(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteBST(root.Right, key)
	} else {
		if root.Left == nil && root.Right == nil { // 2. 左右孩子都为空（叶子节点），直接删除节点， 返回NULL为根节点
			return nil
		} else if root.Left != nil && root.Right == nil { //3. 删除节点的左孩子为空，右孩子不为空，删除节点，右孩子补位，返回右孩子为根节点
			return root.Left
		} else if root.Left == nil && root.Right != nil { //4. 删除节点的右孩子为空，左孩子不为空，删除节点，左孩子补位，返回左孩子为根节点
			return root.Right
		} else { //5.左右孩子节点都不为空，则将删除节点的左子树头结点（左孩子）放到删除节点的右子树的最左面节点的左孩子上，返回删除节点右孩子为新的根节点。
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = root.Left
			return root.Right
		}
	}

	return root
}

// 普通二叉树的删除
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key {
		if root.Right == nil {
			return root.Left
		}
		cur := root.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		root.Val, cur.Val = cur.Val, root.Val
	}
	root.Left = deleteNode(root.Left, key)
	root.Right = deleteNode(root.Right, key)
	return root
}

func main() {
	// 切片的插入  方式1： 通过一个临时的切片进行复制
	var a = make([]int, 5)
	i := 10
	a = append(a[:i], append([]int{1, 2, 3}, a[i:]...)...)

	//  切片的插入  方式2： 使用copy
	a = append(a, 0)     // 切片拓展一个空间
	copy(a[i+1:], a[i:]) // a[i:] 向后启动一个位置
	a[i] = 10

	//  切片的插入  方式3： 使用copy和append组合
	x := []int{1, 2, 3}
	a = append(a, x...)       // 为x切片扩展足够的空间
	copy(a[i+len(x):], a[i:]) // a[i:]向后移动len(x)个位置
	copy(a[i:], x)            // 复制新添加的切片
}

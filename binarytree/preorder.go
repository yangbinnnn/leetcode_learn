package main

import "fmt"

// Given a binary tree, return the preorder traversal of its nodes' values.
//
// For example:
// Given binary tree [1,null,2,3],
//    1
//     \
//      2
//     /
//    3
// return [1,2,3].
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// root -> left -> right
func preorderTraversal(root *TreeNode) []int {
	return recursive(root)
}

func recursive(root *TreeNode) []int {
	vals := []int{}
	if root == nil {
		return vals
	}
	// root fisrt
	vals = append(vals, root.Val)
	// left secord
	if root.Left != nil {
		vals = append(vals, recursive(root.Left)...)
	}
	// right third
	if root.Right != nil {
		vals = append(vals, recursive(root.Right)...)
	}
	return vals
}

func iteratively(root *TreeNode) []int {
	// 使用stack 可以实现遍历, stack 大小小于等于树深度
}

func main() {
	fmt.Println("vim-go")
}

package treeTag

/**
102. 二叉树的层次遍历
 */

import (
	"github.com/chentaihan/leetcode/common"
	"fmt"
)

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := QueueEx{}
	ret := [][]int{}
	queue.Push(root, 0)
	for !queue.Empty() {
		node, level := queue.Pop()
		if node.Left != nil {
			queue.Push(node.Left, level+1)
		}
		if node.Right != nil {
			queue.Push(node.Right, level+1)
		}
		if level >= len(ret) {
			ret = append(ret, []int{node.Val})
		}else{
			ret[level] = append(ret[level], node.Val)
		}
	}
	return ret
}

func TestlevelOrder() {
	tests := []struct {
		nums   []int
		result [][]int
	}{
		{
			[]int{},
			[][]int{{}},
		},
		{
			[]int{7},
			[][]int{{7}},
		},
		{
			[]int{7, 3},
			[][]int{{7}, {3}},
		},
		{
			[]int{7, 15},
			[][]int{{7}, {15}},
		},
		{
			[]int{7, 3, 15},
			[][]int{{7}, {3, 15}},
		},
		{
			[]int{7, 3, 15, 9},
			[][]int{{7}, {3, 15}, {9}},
		},
		{
			[]int{7, 3, 15, 20},
			[][]int{{7}, {3, 15}, {20}},
		},
		{
			[]int{7, 3, 15, 9, 20},
			[][]int{{7}, {3, 15}, {9, 20}},
		},
		{
			[]int{7, 3, 15, 9, 20, 1},
			[][]int{{7}, {3, 15}, {1, 9, 20}},
		},
		{
			[]int{7, 3, 15, 9, 20, 1, 4},
			[][]int{{7}, {3, 15}, {1, 4, 9, 20}},
		},
		{
			[]int{7, 3, 15, 9, 20, 1, 4},
			[][]int{{7}, {3, 15}, {1, 4, 9, 20}},
		},
		{
			[]int{7, 3, 15, 9, 20, 1, 4,30},
			[][]int{{7}, {3, 15}, {1, 4, 9, 20},{30}},
		},
	}
	for _, test := range tests {
		root := CreateTree(test.nums)
		ret := levelOrder(root)
		for index, list := range ret {
			fmt.Println(common.IntEqual(test.result[index], list))
		}
	}
}

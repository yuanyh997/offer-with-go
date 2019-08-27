package Offer
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func treeDepth(root *TreeNode) int {
	if root==nil{
		return 0
	}
	return max(treeDepth(root.Left),treeDepth(root.Right))+1
}
func max(a,b int)int{
	if a>b{
		return a
	}else{
		return b
	}
}
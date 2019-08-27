package Offer
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirror(root *TreeNode)  {
	if root==nil{
		return
	}
	root.Left,root.Right = root.Right,root.Left
	mirror(root.Left)
	mirror(root.Right)

}

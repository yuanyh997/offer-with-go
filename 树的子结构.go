package Offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	if pRoot1 == nil || pRoot2 == nil{
		return false
	}else if isSame(pRoot1,pRoot2){
		return true
	}
	return hasSubtree(pRoot1.Left, pRoot2)  ||hasSubtree(pRoot1.Right, pRoot2)
}
func isSame(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	if pRoot2==nil{
		return true
	}else if pRoot1==nil || pRoot1.Val != pRoot2.Val{
		return false
	}
	return isSame(pRoot1.Left, pRoot2.Left) &&isSame(pRoot1.Right, pRoot2.Right)
}
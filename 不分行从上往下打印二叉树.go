package Offer
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func printFromTopToBottom(root *TreeNode) []int {
	if root ==nil{
		return nil
	}
	a := make([]*TreeNode,0)
	a = append(a, root)
	ans := make([]int,0)
	for len(a)!=0{
		ans = append(ans,a[0].Val)
		if a[0].Left != nil{
			a = append(a,a[0].Left)
		}
		if a[0].Right != nil{
			a = append(a,a[0].Right)
		}
		a = a[1:]
	}
	return ans
}
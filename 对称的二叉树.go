package Offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return root==nil || dfs(root.Right,root.Left)
}

func dfs(p,q *TreeNode) bool{
	if(p==nil ||q==nil){
		return p==nil && q==nil
	}
	return p.Val==q.Val && dfs(p.Right,q.Left) && dfs(q.Right,p.Left)

}

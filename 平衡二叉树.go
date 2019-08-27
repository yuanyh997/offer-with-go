package Offer
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var is bool

func isBalanced(root *TreeNode) bool {
	if root==nil{
		return true
	}
	dfs(root)
	return is==false
}
// 当非平衡时，return -1; 平衡时，return high;
// 首先判断左子树平衡与否，再判断右子树平衡与否，在判断整棵树平衡与否；
func dfs(root *TreeNode)int{
	if root==nil{
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	if abs(left-right)>1{
		is = true
	}
	return max(left,right)+1
}

func max(a,b int)int{
	if a>b{
		return a
	}else{
		return b
	}
}

func abs(a int)int{
	if a<0{
		return -a
	}else{
		return a
	}
}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	height := dfs(root)
	if height>=0{
		return true
	}else{
		return false
	}
}

func dfs(root *TreeNode)int{
	if root==nil{
		return 0
	}
	left := dfs(root.Left)
	if left <0{
		return -1
	}
	right := dfs(root.Right)
	if right <0{
		return -1
	}
	if right-left>1 || left-right>1{
		return -1
	}
	return max(left,right)+1
}

func max(a,b int)int{
	if a>b{
		return a
	}else{
		return b
	}
}
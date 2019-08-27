package Offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder)==0||len(inorder)==0{
		return nil
	}
	return dfs(preorder,0,len(preorder)-1,inorder,0,len(inorder)-1)
}
//递归遍历前序数组和中序数组 前序的第一个节点是当前数的根节点，在中序数组中根节点左边的为左子树，右边为右子树
func dfs(preorder []int,preStart int,preEnd int,inorder []int,inStart int,inEnd int)*TreeNode{
	if preStart >preEnd ||inStart >inEnd {//递归出口
		return nil
	}
	root := new(TreeNode)//找到当前子树的根节点
	root.Val = preorder[preStart]
	var rootIndex int//中序序列中根节点位置
	for i:=0;i<len(inorder);i++{
		if inorder[i]==root.Val{
			rootIndex = i
			break
		}
	}

	leftCount := rootIndex - inStart//左子树元素个数
	root.Left = dfs(preorder,preStart+1,preStart+leftCount,inorder,inStart,rootIndex-1)//左子树
	root.Right = dfs(preorder,preStart+leftCount+1,preEnd,inorder,rootIndex+1,inEnd)//右子树
	return root
}
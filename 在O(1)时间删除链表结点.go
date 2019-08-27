package Offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode)  {
	//var pnode *ListNode
	pnode := node.Next
	node.Val = pnode.Val
	node.Next = pnode.Next
}
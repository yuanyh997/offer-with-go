package Offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	node := head
	a := make([]int,0)
	for node!=nil{
		a = append(a,node.Val)
		node = node.Next

	}
	node = head
	i:=len(a)-1
	for node!=nil{
		node.Val = a[i]
		i--
		node = node.Next
	}
	return head

}
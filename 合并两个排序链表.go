package Offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	cur := new(ListNode)
	head := cur
	for l1!=nil &&l2!=nil{
		//fmt.Println(l1.Val,l2.Val)
		if l1.Val<l2.Val{
			cur.Next = l1
			l1 = l1.Next
		}else{
			cur.Next = l2
			l2 = l2.Next
		}
		//fmt.Println(cur.Val)
		cur = cur.Next
	}
	if l1!=nil{
		cur.Next = l1
	}
	if l2!=nil{
		cur.Next = l2
	}
	return head.Next
}
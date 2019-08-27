package Offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func findKthToTail(pListHead *ListNode, k int) *ListNode {
	if pListHead == nil {
		return nil
	}
	l := 0
	head := pListHead
	for head != nil {
		l++
		head = head.Next
	}
	//fmt.Println(l)
	node := pListHead
	l = l - k
	if l < 0 {
		return nil
	}
	for l != 0 {
		node = node.Next
		l--
	}
	return node
}

package Offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func printListReversingly(head *ListNode) []int {
	ans := make([]int,0)
	for head != nil{
		ans = append(ans,head.Val)
		head = head.Next
	}
	res := make([]int,0)
	for i:=len(ans)-1;i>=0;i--{
		res = append(res,ans[i])
	}
	return res
}

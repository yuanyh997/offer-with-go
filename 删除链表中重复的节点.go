package Offer


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplication(head *ListNode) *ListNode {
	m := make(map[int]int)
	node := head
	for node!=nil{
		if _,ok:=m[node.Val];!ok{
			m[node.Val] = 1
			//res = append(res,node.Val)
		}else{
			m[node.Val] = m[node.Val]+1
		}
		node = node.Next
	}
	//fmt.Println(m)
	h :=new(ListNode)
	res := h
	node = head
	for node!=nil{
		if m[node.Val]==1{
			//fmt.Println(node.Val)
			nx :=new(ListNode)
			nx.Val = node.Val
			h.Next = nx
			h = h.Next
		}
		node = node.Next
	}
	return res.Next
}
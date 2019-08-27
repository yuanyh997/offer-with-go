package Offer

type MinStack struct {
	elem []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	var obj MinStack
	return obj
}


func (this *MinStack) Push(x int)  {
	this.elem = append(this.elem,x)
}


func (this *MinStack) Pop()  {
	this.elem = this.elem[:len(this.elem)-1]
}


func (this *MinStack) Top() int {
	return this.elem[len(this.elem)-1]
}


func (this *MinStack) GetMin() int {
	min := this.elem[0]
	for i:=0;i<len(this.elem);i++{
		if this.elem[i]<min{
			min = this.elem[i]
		}
	}
	return min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
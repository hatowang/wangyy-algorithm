package implement_stack_using_queues

import "container/list"

type MyStack struct {
	q *list.List
	top int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
return MyStack{
	q: new(list.List),
}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.q.PushBack(x)
	this.top = x
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	return this.q.Remove(this.q.Back()).(int)
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.q.Back().Value.(int)
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.q.Len() == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
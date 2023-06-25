// Day 14

type MinStack struct {
    items []int
    minStk []int
}


func Constructor() MinStack {
    stk := MinStack{}
    return stk
}


func (this *MinStack) Push(val int)  {
    this.items = append(this.items, val)
    
    if len(this.minStk) == 0 {
        this.minStk = append(this.minStk, val)
        return
    }
    popMin := this.minStk[len(this.minStk) - 1]
    if val <= popMin {
        this.minStk = append(this.minStk, val)
    }
}


func (this *MinStack) Pop()  {
    popV := this.items[len(this.items) - 1]
    this.items = this.items[:len(this.items) - 1]
    popMin := this.minStk[len(this.minStk) - 1]
    if popV == popMin {
        this.minStk = this.minStk[:len(this.minStk) - 1]
    }

}


func (this *MinStack) Top() int {
    popV := this.items[len(this.items) - 1]
    return popV
}


func (this *MinStack) GetMin() int {
    return this.minStk[len(this.minStk) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
package offer

type CQueue struct {
	stIn []int
	stOut []int
}


func Constructor() CQueue {
	return CQueue{
			stIn: make([]int, 0),
			stOut: make([]int, 0),
	}
}


func (this *CQueue) AppendTail(value int)  {
	this.stIn = append(this.stIn, value)
}


func (this *CQueue) DeleteHead() int {
	 if len(this.stOut) == 0 {
			if len(this.stIn) == 0 {
					return -1
			}
			for len(this.stIn) != 0 {
					this.stOut = append(this.stOut, this.stIn[len(this.stIn)-1])
					this.stIn = this.stIn[:len(this.stIn)-1]
			}
	}

	value := this.stOut[len(this.stOut)-1]
	this.stOut = this.stOut[:len(this.stOut)-1]
	return value
}


/**
* Your CQueue object will be instantiated and called as such:
* obj := Constructor();
* obj.AppendTail(value);
* param_2 := obj.DeleteHead();
*/
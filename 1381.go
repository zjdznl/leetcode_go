package leetcode

type CustomStack struct {
	values    []int // a ring list
	incrCache []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		values:    make([]int, 0, maxSize),
		incrCache: make([]int, maxSize),
	}
}

func (this *CustomStack) Push(x int) {
	if len(this.values) == cap(this.values) {
		return
	}
	this.values = append(this.values, x)
}

func (this *CustomStack) Pop() int {
	if len(this.values) == 0 {
		return -1
	}

	r := this.values[len(this.values)-1] + this.incrCache[len(this.values)-1]
	if len(this.values) > 1 {
		this.incrCache[len(this.values)-2] += this.incrCache[len(this.values)-1]
	}
	this.incrCache[len(this.values)-1] = 0
	this.values = this.values[:len(this.values)-1]

	return r

}

func (this *CustomStack) Increment(k int, val int) {
	if k > len(this.values) {
		k = len(this.values)
	}
	if k > 0 {
		this.incrCache[k-1] += val
	}
}

func (this *CustomStack) Size() int {
	return len(this.values)
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */

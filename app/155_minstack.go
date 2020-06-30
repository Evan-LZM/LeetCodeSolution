package app

type MinStack struct {
	arr []Elem
}

type Elem struct {
	value    int
	minsoFar int
}

/** initialize your data structure here. */
func ConstructorStack() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	e := Elem{value: x}
	if len(this.arr) == 0 || x < this.arr[len(this.arr)-1].minsoFar {
		e.minsoFar = x
	} else {
		e.minsoFar = this.arr[len(this.arr)-1].minsoFar
	}
	this.arr = append(this.arr, e)
}

func (this *MinStack) Pop() {
	this.arr = this.arr[:len(this.arr)-1]
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1].value
}

func (this *MinStack) GetMin() int {
	return this.arr[len(this.arr)-1].minsoFar
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

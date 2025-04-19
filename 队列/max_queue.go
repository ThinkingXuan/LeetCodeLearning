package main

// MaxQueue 最大值的队列
type MaxQueue struct {
	queue []int
	deque []int
}

func NewMaxQueue() MaxQueue {
	return MaxQueue{queue: []int{}, deque: []int{}}
}

func (this *MaxQueue) maxValue() int {
	if len(this.deque) == 0 {
		return -1
	}
	return this.deque[0]
}

func (this *MaxQueue) pop() int {
	if len(this.queue) == 0 {
		return -1
	}

	if this.deque[0] == this.queue[0] {
		this.deque = this.deque[1:]
	}
	v := this.queue[0]
	this.queue = this.queue[1:]
	return v
}

func (this *MaxQueue) push(val int) {
	this.queue = append(this.queue, val)
	for len(this.deque) > 0 {
		if this.deque[len(this.deque)-1] < val {
			this.deque = this.deque[:len(this.deque)-1]
		} else {
			break
		}
	}
	this.deque = append(this.deque, val)
}

// 另外的方法
// 优先级队列： 参考heap的实现

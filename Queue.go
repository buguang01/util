package util

import (
	"errors"
)

const (
	defaultQueueSize = 10
)

type Queue struct {
	front        int
	rear         int
	currentCount int
	queueSize    int
	elements     []interface{}
}

/**
  指定大小的初始化
*/
func NewQueueBySize(size int) *Queue {
	return &Queue{0, size - 1, 0, size, make([]interface{}, size)}
}

/**
  按默认大小进行初始化
*/
func NewQueue() *Queue {
	return NewQueueBySize(defaultQueueSize)
}

/**
  向下一个位置做探测
*/
func (queue *Queue) ProbeNext(i int) int {
	return (i + 1) % queue.queueSize
}

func (queue *Queue) Probe(i int) int {
	return i % queue.queueSize
}

/**
  清空队列
*/
func (queue *Queue) ClearQueue() {
	queue.front = 0
	queue.rear = queue.queueSize - 1
	queue.currentCount = 0
}

/**
  是否为空队列
*/
func (queue *Queue) IsEmpty() bool {
	if queue.currentCount == 0 {
		return true
	}
	return false
}

/**
  队列是否满了
*/
func (queue *Queue) IsFull() bool {
	if queue.currentCount == queue.queueSize {
		return true
	}
	return false
}

/**
  入队
*/
func (queue *Queue) Offer(e interface{}) error {
	if queue.IsFull() == true {
		return errors.New("the queue is full.")
	}
	queue.rear = queue.ProbeNext(queue.rear)
	queue.elements[queue.rear] = e
	queue.currentCount = queue.currentCount + 1
	return nil
}

/**
  出队一个元素
*/
func (queue *Queue) Poll() (interface{}, error) {
	if queue.IsEmpty() == true {
		return nil, errors.New("the queue is empty.")
	}
	tmp := queue.front
	queue.front = queue.ProbeNext(queue.front)
	queue.currentCount = queue.currentCount - 1
	return queue.elements[tmp], nil
}

//拿当前成员
func (queue *Queue) GetCurrElement() (interface{}, error) {
	if queue.IsEmpty() == true {
		return nil, errors.New("the queue is empty.")
	}
	return queue.elements[queue.front], nil
}

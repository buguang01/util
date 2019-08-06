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

//标识UID的成员接口
type IQueueUID interface {
	GetUID() int        //在队列中的标识ID
	SetUID(uid int) int //设置在队列中的标识ID
}

//加入的成员需要支持IQueueUID接口
type QueueUID struct {
	*Queue
	curruid int //当前标识ID
}

//返回列表,需要成员继承IQueueUID接口
func (queue *QueueUID) GetArray(uidmax int) (*[]interface{}, error) {
	result := make([]interface{}, 0, queue.queueSize)
	for index := 0; index < queue.currentCount; index++ {
		tmp, ok := queue.elements[queue.Probe(queue.front+index)].(IQueueUID)
		if !ok {
			result = append(result, tmp)
		}
		if tmp.GetUID() > uidmax {
			result = append(result, tmp)
		}
	}

	return &result, nil
}

/**
  入队
*/
func (queue *QueueUID) Offer(e IQueueUID) error {
	if queue.IsFull() == true {
		return errors.New("the queue is full.")
	}
	queue.curruid++
	e.SetUID(queue.curruid)
	queue.rear = queue.ProbeNext(queue.rear)
	queue.elements[queue.rear] = e
	queue.currentCount = queue.currentCount + 1
	return nil
}

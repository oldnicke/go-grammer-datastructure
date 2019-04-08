package main

import "fmt"

type SampleQueue interface {
	Size()                    //大小
	Front() interface{}       //第一个元素
	End() interface{}         //最后一个元素
	IsEmpty() bool            //判空
	Enqueue(data interface{}) //入队
	Dequeue() interface{}     //出队
	Clear()                   //清空队列
}

type Queue struct {
	datastore []interface{}
	theSize   int
}

//清空队列
func (queue *Queue) Clear() {
	queue.datastore = make([]interface{}, 0) //开辟内存空间
	queue.theSize = 0
}
func NewQueue() *Queue {
	queue := new(Queue)
	queue.Clear()
	return queue
}
func (queue *Queue) Size() int {
	return queue.theSize //大小
}

func (queue *Queue) Front() interface{} {
	if queue.Size() == 0 {
		return nil
	}
	return queue.datastore[0]
}
func (queue *Queue) End() interface{} {
	if queue.Size() == 0 {
		return nil
	}
	return queue.datastore[queue.theSize-1]
}

func (queue *Queue) IsEmpty() bool {
	return queue.theSize == 0
}
func (queue *Queue) Enqueue(data interface{}) {
	queue.datastore = append(queue.datastore, data)
	queue.theSize++
}
func (queue *Queue) Dequeue() interface{} {
	if queue.Size() == 0 {
		return nil
	}
	data := queue.datastore[0]
	if queue.Size() > 1 {
		queue.datastore = queue.datastore[1:queue.theSize]
	}
	queue.theSize--
	return data
}

func main() {
	myqueue := NewQueue()
	myqueue.Enqueue(1)
	myqueue.Enqueue(2)
	myqueue.Enqueue(3)
	myqueue.Enqueue(4)
	fmt.Println(myqueue.theSize)
	fmt.Println(myqueue.Size())
	fmt.Println(myqueue.Dequeue())
	fmt.Println(myqueue.Dequeue())
	fmt.Println(myqueue.Dequeue())
	fmt.Println(myqueue.Dequeue())
	fmt.Println(myqueue.Size())
	myqueue.Enqueue(1)
	myqueue.Enqueue(2)
	myqueue.Enqueue(3)
	myqueue.Enqueue(4)
	myqueue.Enqueue(5)

}

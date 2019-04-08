package main

import (
	"fmt"
	"github.com/go-errors/errors"
)

type StackArray interface {
	Clear()                //清空
	Size() int             //计算栈大小
	Pop() interface{}      //出栈
	Push(data interface{}) //进栈
	IsEmpty() bool         //判空
	IsFull() bool          //判满
}

//栈结构体
type Stack struct {
	stacklist *ArrayList
	capsize   int
}
type ArrayList struct {
	dataStore []interface{}
	theSize   int
}

//
func New() *ArrayList {
	stack := new(ArrayList)
	stack.dataStore = make([]interface{}, 0, 10) //分配内存10个数组元素
	stack.theSize = 0
	return stack
}

//追加
func (arraylist *ArrayList) Append(newvalue interface{}) {
	arraylist.dataStore = append(arraylist.dataStore, newvalue)
	arraylist.theSize++

}

//移除
func (arraylist *ArrayList) Remove(index int) error {
	if index < 0 || index >= arraylist.theSize {
		return errors.New("索引越界")
	}
	arraylist.dataStore = append(arraylist.dataStore[:index], arraylist.dataStore[index+1:]...)
	arraylist.theSize--
	return nil
}

//清空
func (arraylist *ArrayList) Clear() {
	arraylist.dataStore = make([]interface{}, 0, 10) // 清空
	arraylist.theSize = 0
}

func NewStack() *Stack {

	stack := new(Stack)
	stack.stacklist = New()
	stack.capsize = 10
	return stack
}
func (stack *Stack) Clear() {
	stack.stacklist.dataStore = make([]interface{}, 0, 10)
	stack.stacklist.theSize = 0
}

func (stack *Stack) Size() int {
	return stack.stacklist.theSize
}
func (stack *Stack) IsEmpty() bool {
	if stack.stacklist.theSize == 0 {
		return true
	} else {
		return false
	}
}
func (stack *Stack) IsFull() bool {
	if stack.stacklist.theSize == stack.capsize {
		return true
	} else {
		return false
	}
}
func (stack *Stack) Pop() interface{} {
	if !stack.IsEmpty() {
		last := stack.stacklist.dataStore[stack.stacklist.theSize-1]
		stack.stacklist.Remove(stack.stacklist.theSize - 1)
		return last
	}
	return nil
}
func (stack *Stack) Push(data interface{}) {
	if !stack.IsFull() {
		stack.stacklist.Append(data)
	}

}

func main() {
	mystack := NewStack()
	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)
	mystack.Push(4)
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())
	fmt.Println(mystack.Pop())

}

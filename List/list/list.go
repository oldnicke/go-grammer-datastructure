package main

import (
	"fmt"
	"github.com/btcsuite/goleveldb/leveldb/errors"
)

var Listlength = 10 //定义列表全局长度
type List interface {
	Size() int                                    //返回大小
	Get(index int) (interface{}, error)           //根据索引抓取数据
	Set(index int, newvalue interface{}) error    //设置
	Insert(index int, newvalue interface{}) error //插入数据
	Append(newvalue interface{})                  //追加
	Remove(index int) error                       //删除
	Clear()                                       //清空
	String() string                               //返回字符串
}

type ArrayList struct {
	dataSore []interface{}
	theSize  int
}

//创建新的链表
func New() *ArrayList {
	list := new(ArrayList)
	list.dataSore = make([]interface{}, 0, Listlength) // 分配内存10个数组元素
	list.theSize = 0
	return list
}

//获取链表的长度
func (list *ArrayList) Size() int {
	return list.theSize //返回大小
}
func (list *ArrayList) Append(newvalue interface{}) {
	list.dataSore = append(list.dataSore, newvalue) //数据叠加
	list.theSize++                                  //索引移动
}

//获取索引所在数据
func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.theSize {
		return nil, errors.New("索引越界")
	}
	return list.dataSore[index], nil
}

//根据索引更改数据
func (list *ArrayList) Set(index int, newvalue interface{}) error {
	if index < 0 || index >= list.theSize {
		return errors.New("索引越界")
	}
	list.dataSore[index] = newvalue
	return nil
}

//检查切片是否已满
func (list *ArrayList) checkmemisfull() {
	if list.Size() == cap(list.dataSore) {
		newDataStore := make([]interface{}, 0, 2*list.Size()) //开辟更大内存
		copy(newDataStore, list.dataSore)                     //拷贝
		list.dataSore = newDataStore                          //赋值
	}
}

//插入数据
func (list *ArrayList) Insert(index int, newvalue interface{}) error {
	if index < 0 || index >= list.theSize {
		return errors.New("索引越界")
	}
	list.checkmemisfull()
	list.dataSore = list.dataSore[:list.Size()+1] //开辟内存，延展使用内存
	for i := list.Size(); i > index; i-- {
		list.dataSore[i] = list.dataSore[i-1] //从后往前赋值
	}
	list.dataSore[index] = newvalue
	list.theSize++
	return nil

}

//根据索引移除数据
func (list *ArrayList) Remove(index int) error {
	if index < 0 || index >= list.theSize {
		return errors.New("索引越界")
	}
	list.dataSore = append(list.dataSore[:index], list.dataSore[index+1:]...)
	list.theSize--
	return nil
}

//清空数据
func (list *ArrayList) Clear() {
	list.dataSore = make([]interface{}, 0, 10)
	list.theSize = 0
}

//返回字符串
func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataSore)
}

func main() {
	list := New()
	for i := 0; i < 10; i++ {
		list.Append(i)
	}
	fmt.Printf("链表长度：%d\n输出链表\n%d\n", list.theSize, list.dataSore)

	list.Remove(3)
	fmt.Println("输出链表")
	fmt.Println(list.dataSore)

	list.Insert(3, 100)
	fmt.Println("输出链表")
	fmt.Println(list.dataSore)

	fmt.Println("输出链表")
	fmt.Println(list.String())
}

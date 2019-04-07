package main

import "fmt"

type Item interface {
}

type Node struct {
	Data     Item
	NextNode *Node
}

var head *Node
var current *Node

func (node *Node) CreateHeadNode(data Item) {
	node = new(Node)
	node.Data = data
	node.NextNode = nil
	head = node
	current = node
}

//添加新的节点
func (node *Node) AddNode(data Item) {
	node = new(Node)
	node.Data = data
	node.NextNode = nil

	current.NextNode = node
	current = node

}

//遍历链表
func (node *Node) TraveringNodes() {
	node = head
	for {
		if node.NextNode == nil {
			break
		} else {
			node = node.NextNode
		}
	}
}

//计算节点数量
func (node *Node) NodeCnt() int {
	var cnt int = 1
	node = head
	for {
		if node.NextNode == nil {
			break
		} else {
			node = node.NextNode
			cnt = cnt + 1
		}
	}
	return cnt
}

//插入节点
func (node *Node) InsertNodeByIndex(index int, data Item) {
	if index == 0 {
		//添加为新的头节点
		node.Data = data
		node.NextNode = head
		head = node
	} else if index > node.NodeCnt()-1 {
		//添加节点
		node.AddNode(data)
	} else {
		//中间插入节点
		var n = head
		for i := 0; i < index-1; i++ {
			n = n.NextNode
		}
		var newNode *Node = new(Node)
		newNode.Data = data
		newNode.NextNode = n.NextNode
		n.NextNode = newNode
	}
}

//删除节点
func (node *Node) DeleteNodeByIndex(index int) {
	node = head
	if index == 0 {
		//删除头节点，就是第二个节点为头节点
		head = node.NextNode
	} else {
		for i := 0; i < index-1; i++ {
			node = node.NextNode
		}
		node.NextNode = node.NextNode.NextNode
	}
}

//修改指定节点的下表内容

func (node *Node) UpdateNodeByIndex(index int, data Item) {
	node = head
	if index == 0 {
		head.Data = data
	} else {
		for i := 0; i < index; i++ {
			node = node.NextNode
		}
		node.Data = data
	}
}

func main() {

	//创建头结点
	p := head
	p.CreateHeadNode(0)
	//添加节点
	for i := 1; i < 10; i++ {
		p.AddNode(i)
	}
	fmt.Printf("打印所有节点\n")
	p.TraveringNodes()
	fmt.Printf("一共有%d个节点\n", p.NodeCnt())
	fmt.Printf("当前节点:")
	fmt.Println(p.NodeCnt())

	p.DeleteNodeByIndex(2)
	fmt.Println("删除节点2")
	fmt.Printf("打印所有节点\n")
	p.TraveringNodes()

	fmt.Printf("一共有%d个节点\n", p.NodeCnt())
	fmt.Println("插入节点")
	p.InsertNodeByIndex(2, 2)
	fmt.Printf("打印所有节点\n")
	p.TraveringNodes()
}

package main

import "sync"

type DoubleList struct {
	head *ListNode //指向链表头部
	tail *ListNode	//指向链表尾部
	len int	//列表长度
	lock sync.Locker //为了并发安全pop操作
}

type ListNode struct {
	pre *ListNode //前驱节点
	next *ListNode //后驱节点
	value string //值
}

func (node *ListNode) GetValue() string {
	return node.value
}

func (node *ListNode) GerPre() *ListNode  {
	return node.pre
}
func (node *ListNode) GetNext() *ListNode  {
	return node.next
}

// HasNext 是否存在后驱节点
func (node *ListNode) HasNext() bool {
	return node.next != nil
}

// HasPre 是否存在前驱节点
func (node *ListNode) HasPre() bool {
	return node.pre != nil
}

func (node *ListNode) IsNil() bool {
	return node == nil
}

// AddNodeFromHead 添加节点到链表头部的第N个元素前，N=0表示新节点成为新的头部
func (list *DoubleList) AddNodeFromHead(n int,v string)  {
	list.lock.Lock()
	defer list.lock.Unlock()

	if n > list.len {
		panic("index out")
	}

	node := list.head

	for i := 1;i<n;i++ {
		node = node.next
	}

	newNode := new(ListNode)
	newNode.value = v
	if node.IsNil() {
		list.head = newNode
		list.tail = newNode
	} else {
		pre := node.pre

		if pre.IsNil() {
			newNode.next = node
			node.next = newNode
			list.head = newNode
		} else {
			pre.next = newNode
			newNode.pre = pre
		}
	}
}
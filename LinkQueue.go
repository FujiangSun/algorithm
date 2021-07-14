package main

import "sync"

type LinkQueue struct {
	root *LinkNode //链表起点
	size int //队列的元素数量
	lock sync.Mutex //为了并发安全使用的锁
}

type LinkNode struct {
	Next *LinkNode
	Value string
}

func (queue *LinkQueue) Add(v string)  {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	//如果栈顶为空，那么增加节点
	if queue.root == nil {
		queue.root = new(LinkNode)
		queue.root.Value = v
	} else {//否则新元素插入链表的末尾

		//新节点
		newNode := new(LinkNode)
		newNode.Value = v


		nowNode := queue.root

		//一直遍历到链表尾部
		for nowNode.Next != nil {
			nowNode = nowNode.Next
		}
		//新节点放在链表尾部
		nowNode.Next = newNode
	}
	queue.size = queue.size + 1
}

func (queue *LinkQueue) Remove() string  {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	//队中元素已空
	if queue.size == 0 {
		panic("empty")
	}

	//顶部元素要出队
	topNode := queue.root
	v := topNode.Value

	// 将顶部元素的后继链接链上
	queue.root = topNode.Next

	//队列中元素-1
	queue.size = queue.size -1

	return v

}
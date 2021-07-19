package main

import (
	"fmt"
	"sync"
	"unsafe"
)

// Set 集合结构体
type Set struct {
	m map[int]struct{}
	len int
	sync.RWMutex
}

// NewSet 初始化一个集合
func NewSet(cap int64) *Set {
	temp := make(map[int]struct{},cap)
	return &Set{
		m:temp,
	}
}

func main()  {
	a := struct {}{}
	b := struct {}{}

	if a == b{
		fmt.Printf("right:%p\n",&a)
		fmt.Printf("right:%p\n",&a)
	}
	fmt.Println(unsafe.Sizeof(a))
}

// Add 添加一个元素
func (s *Set) Add(item int)  {
	s.Lock()
	defer s.Unlock()
	s.m[item] = struct{}{}
	s.len = len(s.m)
}

// Remove 删除一个元素
func (s *Set) Remove(item int)  {
	s.Lock()
	defer s.Unlock()

	if s.len == 0 {
		return
	}
	delete(s.m,item)//从字典中删除这个键
	s.len = len(s.m)//重新计算元素数量
}

// Has 查看是否存在元素
func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.Unlock()
	_,ok := s.m[item]
	return ok
}

// Len 查看集合大小
func (s *Set) Len() int  {
	return s.len
}

// IsEmpty 检测几集合是否为空
func (s *Set) IsEmpty() bool {
	if s.len == 0 {
		return true
	}
	return false
}

// Clear 清空集合
func (s *Set) Clear()  {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]struct{}{}//字典重新赋值
	s.len =0 //大小归0
}

// List 将集合转换为列表
func (s *Set) List() []int {
	s.RLock()
	defer s.Unlock()
	list := make([]int,0,s.len)
	for item := range s.m{
		list = append(list,item)
	}
	return list
}
package main

import (
	"fmt"
	"sync"
)

type Array struct {
	array []int  //固定大小的数组，用满容量和满大小的切片来代替
	len int	//实际长度
	cap int //容量
	lock sync.Mutex //为了并发安全使用的锁
}

// Make 新建一个可变长的数组
func Make(len,cap int) *Array {
	s := new(Array)
	if len > cap {
		panic("len large than cap")
	}

	//把切片当数组用
	array := make([]int,cap,cap)

	//元数据
	s.array = array
	s.cap = cap
	s.len = 0
	return s
}

// Append 增加一个元素
func (a *Array) Append(element int)  {
	a.lock.Lock()
	defer a.lock.Unlock()

	//当前长度等于容量，表示已经没有多余的位置了
	if a.len == a.cap {
		//容量不够，需要扩容，扩容到两倍
		newCap := 2 * a.cap

		//如果之前的容量为0，那么新增容量为1
		if a.cap == 0 {
			newCap = 1
		}

		newArray := make([]int,newCap,newCap)

		//把老的数组移动到新的数组
		for k,v := range a.array{
			newArray[k] = v
		}
		//替换数组
		a.array = newArray
		a.cap = newCap
	}
	a.array[a.len] = element
	//实际长度+1
	a.len = a.len+1
}

// AppendMany 增加多个元素
func (a *Array) AppendMany(element ...int)  {
	for _,v := range element{
		a.Append(v)
	}
}

// Get 获取某个下标的元素
func (a *Array) Get(index int) int {
	if a.len == 0 || index > a.len {
		panic("index is out of len")
	}
	return a.array[index]
}

// Len 返回实际长度
func (a *Array) Len() int {
	return a.len
}

// Cap 返回容量
func (a *Array) Cap() int {
	return a.cap
}

func Print(array *Array) (result string)  {
	result = "["

	for i := 0; i < array.Len();i++ {
		if i == 0 {
			result = fmt.Sprintf("%s%d",result,array.Get(i))
			continue
		}
		result = fmt.Sprintf("%s %d",result,array.Get(i))
	}
	result = result + "]"
	return
}

func main()  {
	//创建一个容量为3的动态数组
	a := Make(0,3)
	fmt.Println("cap",a.Cap(),"len",a.Len(),"array:",Print(a))


	// 增加一个元素
	a.Append(10)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))

	// 增加一个元素
	a.Append(9)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))

	// 增加多个元素
	a.AppendMany(8, 7)
	fmt.Println("cap", a.Cap(), "len", a.Len(), "array:", Print(a))
}
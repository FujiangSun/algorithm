package main

import "fmt"

// BubbleSort 冒泡排序
func BubbleSort(list []int)  {
	n := len(list)
	
	hasSwap := false

	for i := n-1;i > 0;i-- {
		for j := 0;j < i;j++ {
			if list[j] > list[j+1] {
				list[j],list[j+1] = list[j+1],list[j]
				hasSwap = true
			}
		}
		if ! hasSwap {
			return
		}
	}

}

func main()  {
	list := []int{1,12,31,24,51,26,77,8}
	BubbleSort(list)
	fmt.Println(list)
}
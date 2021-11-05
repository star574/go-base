package main

import "fmt"

// 切片
// 不允许类似于数组的下标赋值 append delete（每次产生一个新的切片） 可以使用类似于数组的使用下标修改值
// len 长度
// cap 容量 默认以2的倍数方式增长 超过1024以后
// 超过长度自动扩容 返回一个新的切片
// slice的增长在1024之前是双倍增长的，而1024以后则是先增长25%以后再调整这个值为系统需要的最小值，因此这个值是约等于n+n/4
// copy 复制切片 覆盖
// [x:y]原数组也会发生变化
func main() {
	var arr [6]int // [0 0 0 0 0 0]
	fmt.Printf("arr: %v\n", arr)

	var slice []int //slice: []
	fmt.Printf("slice: %v\n", slice)
	//
	slice = append(append(slice, 1), 2)
	// slice = append(slice, 2)
	slice[1] = 3
	fmt.Printf("slice: %v\n", slice) // slice: [1 2]

	// 返回第一个到最后一个元素
	fmt.Printf("slice: %v\n", slice[1:])
	// 所有元素
	fmt.Printf("slice: %v\n", slice[:])

	// 查看容量 cap
	fmt.Printf("slice-cap: %v\n", cap(slice)) // slice-cap: 2

	t := make([]int, 3, 4)

	fmt.Printf("t: %v\n", t)  //[ 0 0 0]
	fmt.Printf("t: %d\n", len(t)) // 3
	fmt.Printf("t: %d\n", cap(t)) // 4
}

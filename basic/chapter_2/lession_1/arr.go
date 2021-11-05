package main

import "fmt"

// 数组
// 数组传递依然是值类型
// [3]int [5]int 是不同的类型
func main() {
	var arr [3]int
	fmt.Printf("arr: %v\n", arr) // [0 0 0]

	arr2 := [3]int{1, 2, 3}
	fmt.Printf("arr2: %v\n", arr2)

	var arr3 [3]int = [3]int{1, 23, 4}
	var arr4 [3]int = [3]int{}
	var arr5 [3]int = [3]int{1}
	fmt.Printf("arr3: %v\n", arr3)
	fmt.Printf("arr4: %v\n", arr4)
	fmt.Printf("arr5: %v\n", arr5)

	arr6 := [3]int{}
	fmt.Printf("arr6: %v\n", arr6)

	arr7 := [3]int{2, 2}
	fmt.Printf("arr7: %v\n", arr7)

	// 长度由编译器推导
	arr8 := [...]int{2, 3, 5, 6}
	fmt.Printf("arr8: %v\n", arr8)   // [2 3 5 6]
	fmt.Printf("arr8类型: %T\n", arr8) // [4]int
	fmt.Printf("arr8长度: %d\n", len(arr8))

	// 多维数组

	t := [3][5]int{
		{1, 2, 34},
		{2, 4, 5},
		{1},
	}
	fmt.Printf("t: %v\n", t) // [[1 2 34 0 0] [2 4 5 0 0] [1 0 0 0 0]]

	for i := 0; i < len(t); i++ {
		for j := 0; j < len(t[i]); j++ {
			fmt.Printf("t[%d][%d]: %v\n", i, j, t[i][j])
		}
	}

	for _, v := range t {
		for _, v := range v {
			fmt.Println(v)
		}
		fmt.Println()
	}
	// 默认返回接受第一个值
	for i := range t {
		fmt.Println(i)
	}

}

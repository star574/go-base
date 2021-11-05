package main

import (
	"fmt"
	"sort"
)

// map
// 键类型不能用map,slice,func 值类型可以
// 删除元素 delete(map3, "-")
// map不存在的键 值为初始化值 int 0 slice nil
func main() {

	map1 := map[string]int{}
	map2 := map[string]int{
		"age": 1,
		"num": 10,
	}
	map4 := make(map[string]int, 5)

	fmt.Printf("map1: %v\n", map1)
	fmt.Printf("map2: %v\n", map2)
	fmt.Printf("map4: %v\n", map4)

	map3 := map[string]func(a, b int) int{
		"+": func(a, b int) int {
			return a + b
		},
		"-": func(a, b int) int {
			return a - b
		},
		"*": func(a, b int) int {
			return a * b
		},
		"/": func(a, b int) int {
			return a / b
		}}

	map3["%"] = func(a, b int) int {
		return a % b
	}

	fmt.Printf("map3: %v\n", map3) // map3: map[%:0x5cdde0 *:0x5cdda0 +:0x5cdd60 -:0x5cdd80 /:0x5cddc0]

	fmt.Printf("map3[\"+\"](1, 5): %v\n", map3["+"](1, 5))

	// 删除元素
	delete(map3, "-")

	fmt.Printf("map3: %v\n", map3) // map[%:0x94e280 *:0x94e220 +:0x94e1e0 /:0x94e240]

	// k: +    v: 7
	// k: *    v: 6
	// k: /    v: 0
	// k: %    v: 1
	// 遍历是无需的
	for k, v := range map3 {
		fmt.Printf("k: %v\t", k)
		fmt.Printf("v: %v\n", v(1, 6))
	}

	// 有序遍历
	keys := make([]string, 0, len(map3))
	for k := range map3 {
		keys = append(keys, k)
	}
	// 1.排序
	sort.Strings(keys)
	// 2.遍历
	// v: %-1
	// v: *-6
	// v: +-7
	// v: /-0
	for _, v := range keys {
		fmt.Printf("v: %v-%v\n", v, map3[v](1, 6))
	}

	// 判断是否处在
	val, ok := map3["+"]
	if ok {
		fmt.Printf("val: %v\n", val(1, 7)) // val: 8
	} else {
		fmt.Println("不存在")
	}

	// 判断是否处在
	if val, ok := map3["-"]; ok {
		fmt.Printf("val: %v %v\n", val(8, 1), ok)
	} else {
		fmt.Printf("不存在: %v \n", ok)
	}

}

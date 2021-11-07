package main

import "fmt"

//结构体
// 结构体嵌套
type user struct {
	id      int
	name    string
	address address
}

type address struct {
	mobile string
	city   string
}

// 匿名成员
/// 类型只允许出现一次
type tercher struct {
	int
	string
	addr address
}

func main() {
	// 这种初始化方式成员数量保持一致
	u1 := user{
		1,
		"张三",
		address{
			"1231231231",
			"成都",
		},
	}

	var u2 user

	var u3 = user{}

	u4 := new(user) //指针

	u4.id = 5

	u5 := user{
		id:   2,
		name: "李四",
	}

	u6 := &user{3, "王五", address{"12346", "贵阳"}} //指针 这种初始化方式成员数量保持一致

	fmt.Printf("u1:%T %v\n", u1, u1)
	fmt.Printf("u2:%T %v\n", u2, u2)
	fmt.Printf("u3:%T %v\n", u3, u3)
	fmt.Printf("u4:%T %v\n", u4, u4)
	fmt.Printf("u5:%T %v\n", u5, u5)
	fmt.Printf("u6:%T %v\n", u6, u6)

	// 匿名结构体
	// 必须使用 var关键字

	var person struct {
		id   int
		name string
	}
	fmt.Printf("u6:%T %v\n", person, person)
	person.id = 4
	person.name = "hello"
	person.id = 10
	// 使用的时候自动解指针
	fmt.Printf("u6:%T %v\n", person, person)

	var student struct {
		id   int
		name string
	}
	student.id = 6
	fmt.Printf("u6:%T %v\n", student, student)

	// 按照类型赋值
	t := tercher{
		1,
		"haha",
		address{
			"12312",
			"asdasd",
		},
	}
	t.int = 5
	t.string = "aaa"
	t.addr = address{
		"hello",
		"askjdhkjas",
	}
	fmt.Printf("t: %v\n", t)

}

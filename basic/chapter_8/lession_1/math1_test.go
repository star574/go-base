package test

import (
	"testing"

	"github.com/star574/go-base/basic/chapter_8/lession_1/demo"
)

// 单元测试
// 文件名与_test结尾
// 包名建议test
// 单元测试 测试用例 必须Test开头
// 运行 go test -v / -v -json / -run
// go test -run 可以指定测试用例名称

func testAdd(t *testing.T) {
	// 准备测试数据
	// 开启并行模式
	t.Parallel()
	// 官方 表格驱动
	data := []struct {
		a, b, res int
	}{
		{1, 2, 3},
		{2, 5, 7},
		{3, 7, 11},
	}
	for _, row := range data {
		actual := demo.Add(row.a, row.b)
		if actual != row.res {
			// 标记测试失败
			// t.Fail()
			// 输入报错信息
			t.Errorf("%d+%d=%d !=%d", row.a, row.b, actual, row.res)
		}
	}
}

// go test -v -run TestMul

func testMul(t *testing.T) {
	// 开启并行模式
	t.Parallel()
	// 准备测试数据
	// 官方 表格驱动
	data := []struct {
		a, b, res int
	}{
		{1, 2, 2},
		{2, 5, 10},
		{3, 7, 21},
	}
	for _, row := range data {
		actual := demo.Mul(row.a, row.b)
		if actual != row.res {
			// 标记测试失败
			// t.Fail()
			// 输入报错信息
			t.Errorf("%d*%d=%d !=%d", row.a, row.b, actual, row.res)
		}
	}
}

// 合并测试用例
// 用例名改成小写
// go test .\math1_test.go -v
func TestAddWithTest(t *testing.T) {
	// 可以做一下初始化工作

	// 用例中未开启并行模式则顺序执行
	t.Run("乘法", testMul)
	t.Run("加法", testAdd)

	// 收尾工作
}

package test

import (
	"testing"

	"github.com/star574/go-base/basic/chapter_8/lession_1/demo"
)

// 性能测试
// 用例名称必须以Bechmark开头
// t.N
// go test -bench=Add  -v .\benchmark_test.go
// op: 操作
// BenchmarkAdd-8          1000000000               0.2828 ns/op
func BenchmarkAdd(t *testing.B) {
	for i := 0; i < t.N; i++ {
		demo.Add(111, 222)
	}
}

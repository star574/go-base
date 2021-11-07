package main

import (
	"fmt"
	"os"
)

// panic 与 recover
func main() {
	defer func() {
		err := recover()
		if err != nil {
            // 向上转播的错误 : 自定义错误码:  1001 错误消息:  文件打开错误
			fmt.Printf("向上转播的错误 : %v\n", err)
		}
	}()
	WriterFile("./basic/chapter_6/lession_2/test.txt", "hello world!\n")
}

// 自定义错误
type MyError struct {
	Code   uint
	Reason string
}

// 实现Error方法
func (e MyError) Error() string {
	s := fmt.Sprintln("自定义错误码: ", e.Code, "错误消息: ", e.Reason)
	return s
}

func WriterFile(fileName string, content string) {
	f, err := os.Open(fileName)
	// 错误处理
	defer func() {
		// 捕获错误
		// recover()只能发到defer里面执行
		if recover() != nil {
			// .(type)只能在 switch 里面使用
			err := recover()
			switch err.(type) {
			case *os.PathError:
				fmt.Println("出现错误: 文件不存在")
			default:
				// main.MyError{Code:0x3e9, Reason:"文件打开错误"}
				// fmt.Printf("%#v\n", MyError{Code: 1001, Reason: "文件打开错误"})

				// 向上传播
				panic(MyError{Code: 1001, Reason: "文件打开错误"})
			}
		}
		f.Close()
	}()
	if err != nil {
		// 抛出错误
		panic("出错了")
	}
	f.WriteString(content)
}

// 包注释 包名叫demo
package demo

import "fmt"

// 测试与文档
// go get golang.org/x/tools/cmd/godoc
// godoc -http="localhost:6060"
// 数据库结构对象
type Database struct {
	Host string
	Port string
}

// 常量注释 最大连接数
const (
	Maxconnections = 100
)

// 全局变量注释 数据库对象实例
var (
	Db *Database
)

// 模拟数据库初始化
// 函数注释 数据库初始化 创建数据库对象实例
func NewDatabase(host, port string) *Database {
	return &Database{host, port}
}

// 方法注释 连接数据库
func (db *Database) Connect(dbName string) *Database {
	fmt.Println("数据库已连接")
	return db
}

// 方法注释 执行sql语句
func (db *Database) ExecuteSql(sql string) {
	fmt.Println("执行语句...", sql)
}

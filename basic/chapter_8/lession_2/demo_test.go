package demo

// 生成示例文档
// Output:是必要的
func ExampleDatabase() {
	db := NewDatabase("127.0.0.1", "3306")
	db.Connect("daName").ExecuteSql("select 1 from table_name")

	// Output:
	// 数据库已连接1
	// 执行语句... select 1 from table_name
}

// 添加需要用到但go.mod中查不到的模块,删除未使用的模块 | 创建依赖目录
module github.com/star574/go-base

go 1.17

require (
	github.com/go-sql-driver/mysql v1.6.0
	// 检查更新 | 升级依赖及传递依赖 | 直接更新
	github.com/wechatpay-apiv3/wechatpay-go v0.2.9
)

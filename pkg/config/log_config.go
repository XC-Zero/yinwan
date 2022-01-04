package config

// LogConfig 日志配置
type LogConfig struct {
}

// MysqlLogMode Mysql日志模式
type MysqlLogMode string

// Console 使用 gorm logger，控制台打印sql
// SlowQuery 打印慢查询sql到日志
// None 不开启日志打印
const (
	Console   MysqlLogMode = "console"
	SlowQuery MysqlLogMode = "slow_query"
	None      MysqlLogMode = "none"
)

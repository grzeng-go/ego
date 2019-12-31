package config

import (
	"fmt"
	"github.com/ebar-go/ego/utils/number"
	"net"
	"strconv"
)

const (
	mysqlDefaultIdleConnections = 10
	mysqlDefaultOpenConnections = 40
	mysqlDefaultPort            = 3306
)

// MysqlConfig mysql配置项
type MysqlConfig struct {
	// 数据库名称
	Name string

	// 地址
	Host string

	// 端口号,默认是3306
	Port int

	// 用户名
	User string

	// 密码
	Password string

	// 是否日志模式，默认不开
	LogMode bool

	// 最大操作连接数
	MaxIdleConnections int

	// 最大打开连接数
	MaxOpenConnections int
}

// complete set default config
func (conf *MysqlConfig) complete() {
	conf.MaxIdleConnections = number.DefaultInt(conf.MaxIdleConnections, mysqlDefaultIdleConnections)
	conf.MaxOpenConnections = number.DefaultInt(conf.MaxOpenConnections, mysqlDefaultOpenConnections)
	conf.Port = number.DefaultInt(conf.Port, mysqlDefaultPort)
}

// Dsn return mysql dsn
func (conf MysqlConfig) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		net.JoinHostPort(conf.Host, strconv.Itoa(conf.Port)),
		conf.Name)
}

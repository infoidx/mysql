package mysql

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DSN Data Source Name
type DSN struct {
	Host     string // 主机
	Port     int64  // 端口
	Database string // 数据库名
	User     string // 用户名
	Password string // 密码
}

func (d *DSN) String() string {
	format := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	return fmt.Sprintf(format, d.User, d.Password, d.Host, d.Port, d.Database)
	//format := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&time_zone=%s"
	//return fmt.Sprintf(format, d.User, d.Password, d.Host, d.Port, d.Database, url.QueryEscape(`'Asia/Shanghai'`))
}

// Cluster 集群配置信息
type Cluster struct {
	Sources  []DSN    // 源信息列表
	Replicas []DSN    // 副本信息列表
	Targets  []string // 所控制的对象，可留空
}

// IsEmpty 判断是否为空
func (c *Cluster) IsEmpty() bool {
	return len(c.Sources)+len(c.Replicas) == 0
}

type Config struct {
	LogLevel int
	Primary  DSN          // gorm 会先使用此配置去链接数据库再去配置
	Clusters []Cluster    // 集群配置
	gConfig  *gorm.Config // 不支持全部的gorm配置
}

type ConfigOption func(cfg *Config)

// Configure configure the config
func (c *Config) Configure(opts ...ConfigOption) *Config {
	for _, fn := range opts {
		fn(c)
	}
	return c
}

var SetLogger = func(writer logger.Writer) ConfigOption {
	return func(cfg *Config) {
		//if cfg.LogLevel == 0 {
		//	cfg.LogLevel = 1
		//}
		cfg.gConfig = &gorm.Config{
			Logger: logger.New(
				writer,
				logger.Config{
					LogLevel: logger.LogLevel(cfg.LogLevel),
				},
			),
		}
	}
}

// SetSkipDefaultTransaction 是否跳过默认事务，此类选项参数要想开放在
var SetSkipDefaultTransaction = func(b bool) ConfigOption {
	return func(cfg *Config) {
		cfg.gConfig.SkipDefaultTransaction = b
	}
}

// Apply 修改Config
func (c *Config) Apply(cfg *Config) {
	if cfg != c {
		*cfg = *c
	}
}

// NewConfig create a *Config
func NewConfig(opts ...ConfigOption) *Config {
	cfg := &Config{
		Primary:  DSN{},
		Clusters: make([]Cluster, 0),
		gConfig:  &gorm.Config{},
	}
	cfg.Configure(opts...)
	return cfg
}

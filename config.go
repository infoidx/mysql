package mysql

import (
	"fmt"
	"net/url"

	"gorm.io/gorm"
)

type DSN struct {
	Host     string
	Port     int64
	DataBase string
	Username string
	Password string
}

type Cluster struct {
	Sources  []DSN
	Replicas []DSN
	Targets  []interface{}
}

func (c Cluster) IsEmpty() bool {
	return len(c.Sources)+len(c.Replicas) == 0
}

func (d DSN) String() string {
	format := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&time_zone=%s"
	return fmt.Sprintf(format, d.Username, d.Password, d.Host, d.Port, d.DataBase, url.QueryEscape(`'Asia/Shanghai'`))
}

type Config struct {
	Primary  DSN       // gorm 会先使用此配置去链接数据库再去配置
	Clusters []Cluster // 集群配置
	*gorm.Config
}

type ConfigOption func(cfg *Config)

// Configure configure the config
func (c *Config) Configure(opts ...ConfigOption) {
	for _, fn := range opts {
		fn(c)
	}
}

func (c *Config) Apply(cfg *Config) {
	if cfg != c {
		*cfg = *c
	}
}

// NewConfig create a *Config
func NewConfig(opts ...ConfigOption) *Config {
	// TODO
	cfg := &Config{
		Primary: DSN{},
		Config:  nil,
	}
	cfg.Configure(opts...)
	return cfg
}

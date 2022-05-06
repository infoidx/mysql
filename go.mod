module github.com/infoidx/mysql

go 1.18

replace github.com/infoidx/logger => /Users/xingshanghe/goworkspaces/src/github.com/infoidx/logger

require (
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.4
	gorm.io/plugin/dbresolver v1.1.0
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
)

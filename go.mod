module github.com/infoidx/mysql

go 1.18

replace github.com/infoidx/logger => /Users/xingshanghe/goworkspaces/src/github.com/infoidx/logger

require (
	github.com/infoidx/logger v0.0.0-00010101000000-000000000000
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.4
	gorm.io/plugin/dbresolver v1.1.0
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect
)

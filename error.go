package mysql

import "errors"

var (
	ErrDBInstanceNotInit = errors.New("the db instance has not been initialized")
)

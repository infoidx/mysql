package mysql

import (
	"time"

	driverMySQL "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type DBOption func(db *gorm.DB) error

var SetMaxIdleConns = func(v int) DBOption {
	return func(db *gorm.DB) error {
		if db == nil {
			return ErrDBInstanceNotInit
		}
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		sqlDB.SetMaxIdleConns(v)
		return nil
	}
}

var SetMaxOpenConns = func(v int) DBOption {
	return func(db *gorm.DB) error {
		if db == nil {
			return ErrDBInstanceNotInit
		}
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		sqlDB.SetMaxOpenConns(v)
		return nil
	}
}

var SetConnMaxLifetime = func(d time.Duration) DBOption {
	return func(db *gorm.DB) error {
		if db == nil {
			return ErrDBInstanceNotInit
		}
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		sqlDB.SetConnMaxLifetime(d)
		return nil
	}
}

var SetConnMaxIdleTime = func(d time.Duration) DBOption {
	return func(db *gorm.DB) error {
		if db == nil {
			return ErrDBInstanceNotInit
		}
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		sqlDB.SetConnMaxIdleTime(d)
		return nil
	}
}

// New initialize a new *gorm.DB
func New(cfg *Config, opts ...DBOption) (*gorm.DB, error) {
	db, err := gorm.Open(driverMySQL.Open(cfg.Primary.String()), cfg.Config)
	if err != nil {
		return nil, err
	}
	// 如果有集群配置
	if len(cfg.Clusters) > 0 {
		var dbResolver *dbresolver.DBResolver
		for _, cfgCluster := range cfg.Clusters {
			dbClusterConfig := dbresolver.Config{
				Sources:  make([]gorm.Dialector, 0),
				Replicas: make([]gorm.Dialector, 0),
				Policy:   dbresolver.RandomPolicy{},
			}
			if !cfgCluster.IsEmpty() {
				if len(cfgCluster.Sources) > 0 {
					for _, source := range cfgCluster.Sources {
						dbClusterConfig.Sources = append(dbClusterConfig.Sources, driverMySQL.Open(source.String()))
					}
				}
				if len(cfgCluster.Replicas) > 0 {
					for _, replica := range cfgCluster.Replicas {
						dbClusterConfig.Sources = append(dbClusterConfig.Sources, driverMySQL.Open(replica.String()))
					}
				}
				dbResolver = dbresolver.Register(dbClusterConfig, cfgCluster.Targets...)
			}
		}
		db.Use(dbResolver)
	}

	for _, opt := range opts {
		err = opt(db)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

package mysql

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var Client *gorm.DB

type ConfInfo struct {
	MaxOpenConns      int  `yaml:"max_open_conns"`
	MaxIdleConns      int  `yaml:"max_idle_conns"`
	ConnMaxLifetime   int  `yaml:"conn_max_lifetime"`
	ConnMaxIdleTime   int  `yaml:"conn_max_idle_time"`
	AllowGlobalUpdate bool `yaml:"allow_global_update"` // gorm 配置，默认不支持
	CreateBatchSize   int  `yaml:"create_batch_size"`   // gorm 批量插入数量
}

type DBConf struct {
	WriteDSN string   `yaml:"write_dsn"`
	ReadsDSN []string `yaml:"reads_dsn"`
	ConfInfo ConfInfo `yaml:"conf"`
	Logger   logger.Interface
}

func InitMysql(_ context.Context, conf *DBConf) (*gorm.DB, error) {
	if conf == nil || conf.Logger == nil {
		return nil, fmt.Errorf("input conf must item nil")
	}

	// 连接到主数据库
	dbWrite, err := gorm.Open(mysql.Open(conf.WriteDSN), &gorm.Config{
		AllowGlobalUpdate: conf.ConfInfo.AllowGlobalUpdate,
		CreateBatchSize:   conf.ConfInfo.CreateBatchSize,
		Logger:            conf.Logger,
	})
	if err != nil {
		return nil, fmt.Errorf("write failed to connect to write database, %v", err)
	}
	if len(conf.ReadsDSN) == 0 {
		sqlDB, err := dbWrite.DB()
		if err != nil {
			return nil, fmt.Errorf("sql db err, %v", err)
		}
		sqlDB.SetConnMaxLifetime(time.Duration(conf.ConfInfo.ConnMaxLifetime) * time.Second)
		sqlDB.SetMaxIdleConns(conf.ConfInfo.MaxIdleConns)
		sqlDB.SetMaxOpenConns(conf.ConfInfo.MaxOpenConns)
	}

	// 读写分离
	replicas := make([]gorm.Dialector, 0, len(conf.ReadsDSN))
	for _, dsn := range conf.ReadsDSN {
		replicas = append(replicas, mysql.Open(dsn))
	}

	err = dbWrite.Use(
		dbresolver.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(conf.WriteDSN)},
			Replicas: replicas,
			Policy:   dbresolver.RandomPolicy{},
		}).
			SetMaxIdleConns(conf.ConfInfo.MaxIdleConns).
			SetConnMaxLifetime(time.Duration(conf.ConfInfo.ConnMaxLifetime) * time.Second).
			SetMaxOpenConns(conf.ConfInfo.MaxOpenConns),
	)

	return dbWrite, nil
}

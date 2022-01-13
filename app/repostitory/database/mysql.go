package database

import (
	"github.com/teleport-network/teleport-relayer/app/config"
	tools "github.com/teleport-network/teleport-relayer/tools"
)

var GormClient *tools.GormDB

func NewMysqlDB(mysqlCfg config.Mysql) *tools.GormDB {
	return tools.InitGormDB(&tools.DBConfig{
		DBAddr:           mysqlCfg.MysqlStr,
		MaxIdleConns:     30,
		LogMode:          true,
		AutoCreateTables: nil,
	})
}

func CreateModel(value interface{}) error {
	if GormClient.Client.NewRecord(value) {
		if mydb := GormClient.Client.Create(value); mydb.Error != nil {
			return mydb.Error
		}
	}
	return nil
}

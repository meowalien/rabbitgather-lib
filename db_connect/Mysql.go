package db_connect

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateGormDBConnection(dbconf MysqlConnectConfiguration) (*gorm.DB, error) {
	dsn := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf(dsn, dbconf.User, dbconf.Password, dbconf.Host, dbconf.Port, dbconf.Database)
	dbconn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}
	rawdb, err := dbconn.DB()
	if err != nil {
		return nil, err
	}
	err = rawdb.Ping()
	if err != nil {
		return nil, err
	}
	return dbconn, err
}

func CloseConn(dbconn *gorm.DB) error {
	rawDbconn, err := dbconn.DB()
	if err != nil {
		return err
	}
	err = rawDbconn.Close()
	if err != nil {
		return err
	}
	return nil
}


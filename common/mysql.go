package common

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var GORMMapPool map[string]*gorm.DB

type MysqlMapConf struct {
	List map[string]*MySQLConf `mapstructure:"list"`
}

type MySQLConf struct {
	DriverName      string `mapstructure:"driver_name"`
	DataSourceName  string `mapstructure:"data_source_name"`
	MaxOpenConn     int    `mapstructure:"max_open_conn"`
	MaxIdleConn     int    `mapstructure:"max_idle_conn"`
	MaxConnLifeTime int    `mapstructure:"max_conn_life_time"`
}

func InitDBPool(env *Env) error {
	DbConfMap := &MysqlMapConf{}
	err := env.ParseConfig(env.GetConfPath("mysql_map"), DbConfMap)
	if err != nil {
		return err
	}
	for confName, dbConf := range DbConfMap.List {
		dbGorm, err := gorm.Open("mysql", dbConf.DataSourceName)
		if err != nil {
			return err
		}
		dbGorm.SingularTable(true)
		err = dbGorm.DB().Ping()
		if err != nil {
			return err
		}
		GORMMapPool = map[string]*gorm.DB{}
		GORMMapPool[confName] = dbGorm
	}
	return nil
}

func GetGormPool(name string) (*gorm.DB, error) {
	if gormPool, ok := GORMMapPool[name]; ok {
		return gormPool, nil
	}
	return nil, errors.New("get pool error")
}

func CloseDB() error {
	//关闭数据库连接
	for confName, dbGorm := range GORMMapPool {
		_ = dbGorm.Close()
		fmt.Printf("close %s", confName)
	}
	return nil
}

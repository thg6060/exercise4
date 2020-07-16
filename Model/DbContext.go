package Model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func DbConn() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("mysql", "root:root@/exercise4")
	if err != nil {
		return nil, err
	}
	return engine, nil
}

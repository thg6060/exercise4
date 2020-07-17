package Model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var db, err = DbConn()

func DbConn() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("mysql", "root:root@/exercise4")
	if err != nil {
		return nil, err
	}
	return engine, nil
}

func CreateTable() error {

	defer db.Close()
	err = db.CreateTables(&User{})
	err = db.Sync2(&Point{})

	if err != nil {
		return err
	}

	return nil
}

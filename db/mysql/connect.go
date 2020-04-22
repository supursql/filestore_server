package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/fileserver?charset=utf8")
	db.SingularTable(true)

	defer db.Close()
	if err != nil {
		panic(err)
	}
}

func DBConn() *gorm.DB {
	return db
}

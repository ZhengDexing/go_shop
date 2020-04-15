package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go_shop/shop/util"
)

var db *sql.DB

func init() {
	config, err := util.GetConfig()
	if err != nil {
		panic(err)
	}
	data := config.Db
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=%s", data.User, data.Password, data.Address,
		data.Port, data.Charset)
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

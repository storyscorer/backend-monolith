package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func createDbConnection(host string, port int, uname string, pwd string, dbname string) (*bun.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", uname, pwd, host, port, dbname)
	sqldb, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, errors.New("failed to initialise connection with DB")
	}

	db := bun.NewDB(sqldb, mysqldialect.New())

	return db, nil
}

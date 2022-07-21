package main

// ORM
// Support MySQL PostgreSQL Sqlite3
//
// go get -u github.com/astaxie/beego/orm

// OR

// go get -u github.com/go-sql-driver/mysql

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type dbInstance struct {
	db *sql.DB
	sync.Once
}

var instance = dbInstance{}

func dial() *sql.DB {

	instance.Do(func() {

		instance.db = func() *sql.DB {

			// Open doesn't open a connection(Mysql may have stopped)
			db, err := sql.Open("mysql", "root:admin123@tcp(127.0.0.1:3306)/test_data?collation=utf8mb4_unicode_ci")
			if err != nil {
				panic(err)
			}
			// defer db.Close()

			// Validate DSN data(testing connect)
			err = db.Ping()
			if err != nil {
				panic(err.Error())
			}

			return db
		}()
	})

	return instance.db
}

func ObtainConn() *sql.DB {
	return dial()
}

func DestroyConn() error {
	return dial().Close()
}

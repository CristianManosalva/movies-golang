package database

import "database/sql"

// MySQLClient will retrieve a url and return the body of the page.
type MySQLClient struct {
	*sql.DB
}

// NewSQLClient will retrieve a url and return the body of the page.
func NewSQLClient() *MySQLClient {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/moviesgolang")
	if err != nil {
		panic(err)

	}
	err = db.Ping()
	if err != nil {

	}
	return &MySQLClient{db}
}

package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Connect("postgres", "user=postgres dbname=todo sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	err = CreateSchema()
	if err != nil {
		log.Fatalln(err)
	}
}

func CreateSchema() error {
	insertQuery := `CREATE TABLE IF NOT EXISTS todotasks (id int, task text, description text)`
	_, err := db.Exec(insertQuery)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

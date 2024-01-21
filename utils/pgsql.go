package utils

import (
	"database/sql"
	"log"
	"machine_svc/config"

	_ "github.com/lib/pq"
)

var db *sql.DB

func PGDB() *sql.DB {
	var err error
	if db == nil {
		db, err = sql.Open("postgres", config.Db_conn)
		if err != nil {
			log.Println(err.Error())
			PGDB()
		} else {
			err = db.Ping()
			if err != nil {
				log.Println(err.Error())
				sql.Open("postgres", config.Db_conn)
			}
		}
	}
	return db
}
func SQLExample() {
	cmd := `select now();`
	row := PGDB().QueryRow(cmd)
	var res string
	row.Scan(&res)
	log.Println("SQLEXAMPLE", res)
}

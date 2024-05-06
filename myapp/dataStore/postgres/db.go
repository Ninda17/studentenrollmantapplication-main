package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// db details
const (
	postgres_host     = "dpg-cos5stq1hbls73felcgg-a.singapore-postgres.render.com"
	postgres_port     = 5432
	postgres_user     = "postgres_admin"
	postgres_password = "meE3sZiHFMqZunof2IDxOgq5UeMOFTeU"
	postgres_dbname   = "my_db_xutd"
)

// create pointer variable db which points to sql driver
var Db *sql.DB

// called before main
func init() {
	// creating connection string
	db_info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", postgres_host, postgres_port, postgres_user, postgres_password, postgres_dbname)

	// fmt.Println(db_info)

	var err error
	//open connection to database or establish a connection to postgresql db server using the driver
	Db, err = sql.Open("postgres", db_info)

	//handle error
	if err != nil {
		panic(err)
	} else {
		log.Println("Database sucessfully connected")
	}
}

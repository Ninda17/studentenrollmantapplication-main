package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// db details
const (
	postgres_host     = "db"
	postgres_port     = 5432
	postgres_user     = "postgres"
	postgres_password = "postgres"
	postgres_dbname   = "my_db"
)

// create pointer variable db which points to sql driver
var Db *sql.DB

// called before main
func init() {
	// creating connection string
	db_info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", postgres_host, postgres_port, postgres_user, postgres_password, postgres_dbname)

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

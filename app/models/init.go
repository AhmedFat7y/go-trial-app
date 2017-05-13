package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// used by go sql package
	_ "github.com/go-sql-driver/mysql"

	"github.com/go-gorp/gorp"
)

// Global database references
var db *sql.DB
var dbmap *gorp.DbMap

// Database settings
var dbName = "testsupermam"
var dbUser = "supermum"
var dbPassword = "hello@mum"
var dbHost = "tcp(127.0.0.1:3301)"

// InitDB Create database connection
func InitDB() {
	var err error
	var connectionString = fmt.Sprintf("%v:%v@%v/%v", dbUser, dbPassword, dbHost, dbName)
	log.Println("Connection String: " + connectionString)
	db, err = sql.Open("mysql", connectionString)
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
	if err != nil {
		log.Println("Failed to connect to database: ")
		log.Panic(err)
	} else {
		err = db.Ping()

		if err != nil {
			log.Println("Failed to ping database: ")
			log.Panic(err)
		} else {
			log.Println("Database connected.")
		}
	}

	_ = dbmap.AddTableWithName(Article{}, "flat_articles").SetKeys(false, "ID")
	dbmap.CreateTablesIfNotExists()
}

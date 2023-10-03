package mysql

import (
	"database/sql"
	"log"
	"time"
)

var DB *Database

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Conn     *sql.DB
}

func InitDB(host, port, user, password, database string) {
	if DB != nil {
		return
	}
	DB = &Database{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Database: database,
	}
	var err error

	DB.Conn, err = sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+database)
	if err != nil {
		panic(err.Error())
	}
	go func() {
		for {
			_, err2 := DB.Conn.Exec("SELECT 1")
			if err2 != nil {
				panic(err2.Error())
			}
			// sleep for 1 minute
			<-time.After(1 * time.Minute)
		}
	}()
	log.Println("connected to database")
}

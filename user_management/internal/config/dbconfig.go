package config

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-mysql-org/go-mysql/driver"
)

var Conn *sql.Conn

func DbInit() {
	dsn := "root:root@localhost:3306?users"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
		return
	}

	ctx, close := context.WithTimeout(context.Background(), 10*time.Second)
	defer close()
	Conn, err = db.Conn(ctx)
	defer Conn.Close()
	if err != nil {
		log.Println(err)
		return
	}
}

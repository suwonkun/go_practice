package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
    // Capture connection properties.
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "db",
        DBName: "go_practice",
		AllowNativePasswords: true,
    }
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }

	var name string
	err = db.QueryRow("SELECT name FROM user WHERE id = ?", 1).Scan(&name)

    fmt.Println("Connected!")
	if err != nil {
    fmt.Printf("Error while querying user: %v\n", err)
	} else {
    fmt.Printf("ユーザーID1の名前は: %s\n", name)
	}
}

package main

import (
	"net/http"
	"database/sql"
	"log"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
)

type User struct{
	Name string `json:"name"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:root@tcp(db:3306)/mysql_dev")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    var name string
    err = db.QueryRow("SELECT name FROM users WHERE id = ?", 1).Scan(&name)
    if err != nil {
        log.Fatal(err)
    }

	user := &User{
		Name: name,
	}
	res, err := json.Marshal(user)

	if err != nil {
        log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
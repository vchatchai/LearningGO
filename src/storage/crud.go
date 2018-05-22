package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var Db *sql.DB

const (
	host     = "finwiz-development.grouplease.co.th"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "postgres"
)

func init() {
	fmt.Println("init chatchai")
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("ERROR:", err)
		panic(err)
		return
	}
	fmt.Println("initted")
}

func (p *Post) Create() (err error) {
	statement := "insert into posts (content, author) values($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		panic(err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(p.Content, p.Author).Scan(&p.Id)
	if err != nil {
		panic(err)
		return
	}
	return
}

func CRUD() {
	fmt.Println("StartDB", Db)
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	fmt.Println(post)
	post.Create()
	fmt.Println(post)
}

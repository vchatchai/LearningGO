package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Book struct {
	Id         int
	Content    string
	AuthorName string `db:"author"`
}

var Dbx *sqlx.DB

func init() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	Dbx, err = sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// err = Dbx.Ping()
	// if err != nil {
	// 	panic(err)
	// }
}

func (book *Book) Create() (err error) {
	query := "insert into book (content, author) values($1, $2) returning id"

	err = Dbx.QueryRow(query, book.Content, book.AuthorName).Scan(&book.Id)

	if err != nil {
		panic(err)
	}

	return

}

func GetBook(id int) (book Book, err error) {
	book = Book{}
	err = Dbx.QueryRowx("select id, content, author from book where id = $1", id).StructScan(&book)

	if err != nil {
		panic(err)
	}
	return

}

func SQLX() {
	book := Book{Content: "Hello World", AuthorName: "Sau Sheong"}
	book.Create()
	fmt.Println(book)

	bookx, _ := GetBook(book.Id)
	fmt.Println(bookx)
}

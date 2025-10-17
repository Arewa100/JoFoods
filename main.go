package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"time"
)

func main() {
	fmt.Println("welcome to my database connection")

	db, err := sql.Open("mysql", "root:Racco1999#@/")
	if err != nil {
		panic("error connecting to database")
	}
	fmt.Println(db)
	//defer db.Close()

}

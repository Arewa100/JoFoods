package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Item struct {
	Name      string
	Price     float64
	Available bool
}

func main() {
	fmt.Println("welcome to my database connection")
	connectionString := "root:Racco1999#@tcp(localhost:3306)/jofood"
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic("error opening database: " + err.Error())
	}
	defer db.Close() //freeing up resources

	err = db.Ping()
	if err != nil {
		panic("error connecting to database: " + err.Error())
	}

	fmt.Println("Database connected successfully!")

	//db.SetConnMaxLifetime(time.Minute * 3)
	//db.SetMaxOpenConns(10)
	//db.SetMaxIdleConns(10)

	CreateItemTable(db)
	//product := Item{"suger", 100, true}
	//itemId := insertItem(db, product)
	//fmt.Println("Item Id is ", itemId)
	//
	//secondProduct := Item{"oil", 20, true}
	//itemId = insertItem(db, secondProduct)
	//
	//theItem := getItemByName(db, "oil")
	//fmt.Println(theItem)
	deleteItem(db, "oil")

}

func CreateItemTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS item (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        price FLOAT NOT NULL,
        available BOOLEAN NOT NULL,
        created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func insertItem(db *sql.DB, item Item) int {
	query := "INSERT INTO item (name, price, available) VALUES (?, ?, ?)"

	result, err := db.Exec(query, item.Name, item.Price, item.Available)
	if err != nil {
		log.Fatal(err)
	}

	itemId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return int(itemId)
}

func getItemByName(db *sql.DB, name string) Item {
	query := "SELECT * FROM item WHERE name = ?"
	rows := db.QueryRow(query, name)
	var item Item
	err := rows.Scan(&item.Name, &item.Price, &item.Available)
	if err != nil {
		log.Fatal(err)
	}
	return item
}
func deleteItem(db *sql.DB, name string) int64 {
	query := "DELETE FROM item WHERE name = ?"
	result, err := db.Exec(query, name)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

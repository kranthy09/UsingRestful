package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	id     int
	name   string
	author string
}

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	log.Println(db)
	if err != nil {
		log.Panicln(err)
	}
	// Create table
	statment, err := db.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64) NULL)")
	if err != nil {
		log.Println(db)
	} else {
		log.Println("Succesfully created table books!")
	}
	statment.Exec()
	// CRUD OPERATIONS
	// Create
	statment, _ = db.Prepare("INSERT INTO books(name, author, isbn) VALUES(?, ?, ?)")
	statment.Exec("A Tale of Two Cities", "Charles Dickens", 140430547)
	log.Panicln("Inserted the book into database")
	// Read
	rows, _ := db.Query("SELECT id, name, author FROM BOOKS")
	var tempBook Book
	for rows.Next() {
		rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
		log.Printf("ID: %d, Book: %s, Author: %s\n", tempBook.id, tempBook.name, tempBook.author)
	}
	// Update
	statment, _ = db.Prepare("update books set name=? where id=?")
	statment.Exec("The Tale of two Cities", 1)
	log.Println("Succesfully Updated the book in database")
	// Delete
	statment, _ = db.Prepare("delete from books where id=?")
	statment.Exec(1)
	log.Println("Successfully deleted the book from database!")
}

package main

import (
	"bookApp/pkg/Base/book"
	postgres "bookApp/pkg/db"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	// Set environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")

	}

	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init")
	}
	log.Println("Postgress connected")

	// Repositories
	bookrepo := book.NewBookRepository(db)
	// bookrepo.Migrations()
	// bookrepo.InsertSampleData()

	// fmt.Println(bookrepo.FindAll())
	// fmt.Println(bookrepo.FindByBookID("2"))
	// fmt.Println(bookrepo.FindByBookName("the"))
	// fmt.Println(bookrepo.FindByAuthorName("ca"))
	// bookrepo.DeleteByID("3")
	fmt.Println(bookrepo.BuyByID("4", 3))
	// bookList.CreateList()
}

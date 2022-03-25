package main

import (
	// "bookApp/pkg/csv_utils"
	"bookApp/pkg/Base/book"
	postgres "bookApp/pkg/db"
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
	bookrepo.Migrations()
	bookrepo.InsertSampleData("books.csv")

	// authorrepo := author.NewAuthorRepository(db)
	// authorrepo.Migrations()
	// authorrepo.InsertSampleData()
	// fmt.Println(authorrepo.FindAuthorsWithBookInfo())
	// fmt.Println(bookrepo.FindByBookID("2"))
	// fmt.Println(bookrepo.FindByBookName("the"))
	// fmt.Println(authorrepo.FindAuthorsWithoutBookInfo())
	// fmt.Println(authorrepo.FindByAuthorID("404"))
	// fmt.Println(authorrepo.FindByAuthorName("J."))
	// fmt.Println(bookrepo.FindByAuthorName("ca"))
	// bookrepo.DeleteByID("4")
	// fmt.Println(bookrepo.BuyByID("4", 3))
	// bookList.CreateList()

	// csv_utils.GetBooksWithWorkerPool("books.csv")
}

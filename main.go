package main

import (
	// "bookApp/pkg/csv_utils"
	postgres "bookApp/common/db"
	"bookApp/domain/author"
	"bookApp/domain/book"
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
	bookRepo := book.NewBookRepository(db)
	// bookRepo.SetupDatabase("books.csv")

	// fmt.Println(bookRepo.FindAll())
	// fmt.Println(bookRepo.FindAllInStock())
	fmt.Println(bookRepo.FindAllBooksUnderPrice(30))
	// fmt.Println(bookRepo.FindByBookID("2"))
	// fmt.Println(bookRepo.FindByBookISBN("9780385093798"))
	// fmt.Println(bookRepo.FindByBookName("the"))
	// fmt.Println(bookRepo.FindByAuthorName("j."))
	// fmt.Println(bookRepo.DeleteByBookID("1"))
	// fmt.Println(bookRepo.BuyByBookID("3", 1))

	authorRepo := author.NewAuthorRepository(db)
	authorRepo.Migrations()
	authorRepo.InsertAuthorData("authors.csv")
	// fmt.Println(authorRepo.FindAuthorsWithBookInfo())
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

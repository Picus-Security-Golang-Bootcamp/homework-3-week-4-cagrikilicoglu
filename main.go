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
		log.Fatal("Error loading .env file.")

	}

	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot be initalized.")
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal("Database connection cannot be closed.")
	}
	defer sqlDb.Close()

	log.Println("Postgress connected")

	// Repositories
	bookRepo := book.NewBookRepository(db)
	bookRepo.SetupDatabase("books.csv")

	authorRepo := author.NewAuthorRepository(db)
	authorRepo.SetupDatabase("authors.csv")

	SampleQueries(*bookRepo, *authorRepo)

	// fmt.Println(bookRepo.FindAll())
	// fmt.Println(bookRepo.FindAllIncludingDeleted())
	// fmt.Println(bookRepo.FindAllInStock())
	// fmt.Println(bookRepo.FindAllBooksUnderPrice(30))
	// fmt.Println(bookRepo.FindByBookID("2"))
	// fmt.Println(bookRepo.FindByBookISBN("9780385093798"))
	// fmt.Println(bookRepo.FindByBookName("the"))
	// fmt.Println(bookRepo.FindByAuthorName("j."))
	// fmt.Println(bookRepo.DeleteByBookID("1"))
	// fmt.Println(bookRepo.BuyByBookID("5", 7))

	// authorRepo.Migrations()
	// authorRepo.InsertAuthorData("authors.csv")
	// fmt.Println(authorRepo.FindAuthorsWithBookInfo())
	// fmt.Println(bookrepo.FindByBookID("2"))
	// fmt.Println(bookrepo.FindByBookName("the"))
	// fmt.Println(authorRepo.FindAuthorsWithoutBookInfo())
	// fmt.Println(authorRepo.FindByAuthorID("404"))
	// fmt.Println(authorRepo.FindByAuthorName("J."))
	// fmt.Println(bookrepo.FindByAuthorName("ca"))
	// bookrepo.DeleteByID("4")
	// fmt.Println(bookrepo.BuyByID("4", 3))
	// bookList.CreateList()

	// csv_utils.GetBooksWithWorkerPool("books.csv")
}

func SampleQueries(bookRepo book.BookRepository, authorRepo author.AuthorRepository) {

	// // Find all books in the book list (excluding soft-deleted)
	// books, _ := bookRepo.FindAll()
	// for _, book := range books {
	// 	fmt.Println(book.ToString())
	// }

	// // Find all books in the book list including soft-deleted
	// books, _ = bookRepo.FindAllIncludingDeleted()
	// for _, book := range books {
	// 	fmt.Println(book.ToString())
	// }

	// // Find all books currently in stock
	// books, _ = bookRepo.FindAllInStock()
	// for _, book := range books {
	// 	fmt.Println(book.ToString())
	// }

	// Find all books under a certain price & currently in stock
	// var samplePriceInput float32 = 16
	// books, _ = bookRepo.FindAllBooksUnderPrice(float32(samplePriceInput))
	// for _, book := range books {
	// 	fmt.Println(book.ToString())
	// }

	// // Find a book by ID
	// var sampleIDInput string = "3"
	// book, _ := bookRepo.FindByBookID(sampleIDInput)
	// fmt.Println(book.ToString())

	// // Find a book by ISBN
	// var sampleISBNInput string = "9780385093798"
	// book, _ = bookRepo.FindByBookISBN(sampleISBNInput)
	// fmt.Println(book.ToString())

	// // Find books by name (elastic search)
	// var sampleBookNameInput string = "the"
	// books, _ = bookRepo.FindByBookName(sampleBookNameInput)
	// for _, book := range books {
	// 	fmt.Println(book.ToString())
	// }

	// // Find books by author's name (elastic search)
	// var sampleAuthorNameInput string = "tolkien"
	// books, _ = bookRepo.FindByAuthorName(sampleAuthorNameInput)
	// for _, book := range books {
	// 	fmt.Println(book.ToString())
	// }

	// // Delete books by ID input
	// var sampleDeleteInputID string = "3"
	// bookRepo.DeleteByBookID(sampleDeleteInputID)

	// // Buy books by ID input and quantity input
	// var sampleBuyInputID string = "5"
	// var sampleQuantity int = 1
	// var sampleQuantityForNotEnoughStock int = 4
	// bookRepo.BuyByBookID(sampleBuyInputID, sampleQuantity)
	// // // To check "not enough stock error"
	// err := bookRepo.BuyByBookID(sampleBuyInputID, sampleQuantityForNotEnoughStock)
	// fmt.Println(err)

	// // Find all authors with their books info
	// authors, _ := authorRepo.FindAuthorsWithBookInfo()
	// for _, author := range authors {
	// 	fmt.Println(author.ToString())
	// }

	// // Find all authors without their books info
	// authors, _ := authorRepo.FindAuthorsWithoutBookInfo()
	// for _, author := range authors {
	// 	fmt.Println(author.ToString())
	// }

	// // Find an author with ID input
	// var sampleAuthorID string = "404"
	// author, _ := authorRepo.FindByAuthorID(sampleAuthorID)
	// fmt.Println(author.ToString())

	// // Find authors by name (elastic search)
	// var sampleAuthorName string = "j."
	// authors, _ = authorRepo.FindByAuthorName(sampleAuthorName)
	// for _, author := range authors {
	// 	fmt.Println(author.ToString())
	// }

	// // Find books of an author by giving name input (of author's name) (elastic search)
	// var sampleAuthorName string = "cao"
	// authors, _ = authorRepo.FindBooksOfAuthorByName(sampleAuthorName)
	// for _, author := range authors {
	// 	fmt.Println(author.ToString())
	// }

}

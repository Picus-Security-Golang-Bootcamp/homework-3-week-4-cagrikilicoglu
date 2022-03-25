package book

import (
	"fmt"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

// SetupDatabase: automatically migrates database of Books with gorm and insert book data to database by the given input path
func (b *BookRepository) SetupDatabase(path string) {
	b.Migrations()
	b.InsertBookData(path)
}

// Migrations: automatically migrates database of Books
func (b *BookRepository) Migrations() {
	b.db.AutoMigrate(&Book{})
}

// InsertBookData: insert book data to database by the given input path
func (b *BookRepository) InsertBookData(path string) {
	books, err := readBooksWithWorkerPool(path)
	if err != nil {
		return
	}
	for _, book := range books {
		b.db.Where(Book{ID: book.ID}).Attrs(Book{ID: book.ID, Name: book.Name, PageNumber: book.PageNumber, StockNumber: book.StockNumber, StockID: book.StockID, Price: book.Price, ISBN: book.ISBN, IsDeleted: book.IsDeleted, AuthorID: book.AuthorID, AuthorName: book.AuthorName}).FirstOrCreate(&book)
	}
}

// FindAll(): return all the books in database
func (b *BookRepository) FindAll() []Book {
	books := []Book{}
	b.db.Find(&books)
	return books
}

// FindAllIncludingDeleted(): return all the books including the deleted ones in database
func (b *BookRepository) FindAllIncludingDeleted() []Book {
	books := []Book{}
	b.db.Unscoped().Find(&books)
	return books
}

// FindAllInStock(): find all books that are currently in stock (stock number > 0).
// Warning: this function is not for showing deleted books, it checks the stock numbers.
func (b *BookRepository) FindAllInStock() []Book {
	books := []Book{}
	b.db.Where("stock_number > ?", 0).Find(&books)
	return books
}

// FindAllUnderPrice(): find all books under a given price input and also that are currently in stock.
func (b *BookRepository) FindAllBooksUnderPrice(price float32) ([]Book, error) {
	books := []Book{}
	b.db.Where("stock_number > ?", 0).Where("price < ?", price).Find(&books)
	if len(books) == 0 {
		return nil, fmt.Errorf("There is no books in the stock under %.2f\n", price)
	}
	return books, nil
}

// FindByBookID: returns the book with given ID input
func (b *BookRepository) FindByBookID(ID string) (*Book, error) {
	book := Book{}
	result := b.db.First(&book, ID)
	if result.Error != nil {
		return nil, result.Error
	}
	// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 	return nil, result.Error
	// }
	return &book, nil
}

// FindByBookISBN: returns the book with given ISBN input
func (b *BookRepository) FindByBookISBN(ISBN string) (*Book, error) {
	book := Book{}
	result := b.db.Where(&Book{ISBN: ISBN}).Find(&book)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

// FindByBookName: returns the book/s with given name input
// the search is elastic and case insensitive
func (b *BookRepository) FindByBookName(name string) []Book {
	books := []Book{}
	nameString := fmt.Sprintf("%%%s%%", name)
	b.db.Where("name ILIKE ?", nameString).Find(&books)
	return books
}

// FindByAuthorName: returns the book/s with given author name input
// the search is elastic and case insensitive
func (b *BookRepository) FindByAuthorName(name string) []Book {
	books := []Book{}
	nameString := fmt.Sprintf("%%%s%%", name)
	b.db.Where("author_name ILIKE ?", nameString).Find(&books)
	return books
}

// DeleteByBookID: soft deletes book from the database
func (b *BookRepository) DeleteByBookID(id string) error {
	// book := Book{}
	book, err := b.FindByBookID(id)
	if err != nil {
		return err
	}
	result := b.db.Delete(&book)
	if result.Error != nil {
		return result.Error
	}
	// buraya return fonksiyonu eklenebilir
	// b.db.Unscoped().Where(&Book{ID: id}).Find(&book)
	// fmt.Println(book)
	return nil
}

// BuyByBookID: orders books that is in the database (not soft deleted) with given id input and requested quantity only if there is enough stock for the order.
func (b *BookRepository) BuyByBookID(id string, num int) error {
	// book := Book{}
	book, err := b.FindByBookID(id)
	if err != nil {
		return err
	}
	if book.StockNumber >= num {
		b.db.Model(&book).Update("stock_number", book.StockNumber-num)
	} else {
		return fmt.Errorf("Not enough stock for %s, please order less than %d book/s.\n", book.Name, book.StockNumber)
	}
	// b.db.Model(&book).Where("stock_number >= ?", num).Update("stock_number", book.StockNumber-num)
	fmt.Println(book.StockNumber)
	return nil
}

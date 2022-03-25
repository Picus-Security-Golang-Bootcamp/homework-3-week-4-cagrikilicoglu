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

// alttaki iki fonksiyon ssetup içine alınabilir
func (b *BookRepository) Migrations() {
	b.db.AutoMigrate(&Book{})
}

func (b *BookRepository) InsertSampleData(path string) {
	// books := []Book{
	// 	{
	// 		ID:          "1",
	// 		Name:        "A Tale of Two Cities",
	// 		PageNumber:  320,
	// 		StockNumber: 10,
	// 		StockID:     "21AC",
	// 		Price:       15.3,
	// 		ISBN:        "9780451530578",
	// 		IsDeleted:   false,
	// 		Author:      Author{ID: "101", Name: "Charles Dickens"},
	// 	},
	// 	{
	// 		ID:          "2",
	// 		Name:        "The Hobbit",
	// 		PageNumber:  376,
	// 		StockNumber: 10,
	// 		StockID:     "44UY",
	// 		Price:       24.0,
	// 		ISBN:        "9780547928227",
	// 		IsDeleted:   false,
	// 		Author:      Author{ID: "202", Name: "J. R. R. Tolkien"},
	// 	},
	// 	{
	// 		ID:          "3",
	// 		Name:        "Harry Potter and the Philosophers Stone",
	// 		PageNumber:  560,
	// 		StockNumber: 10,
	// 		StockID:     "22OL",
	// 		Price:       32.2,
	// 		ISBN:        "9781408855898",
	// 		IsDeleted:   false,
	// 		Author:      Author{ID: "303", Name: "J. K. Rowling"},
	// 	},
	// 	{
	// 		ID:          "4",
	// 		Name:        "The Little Prince",
	// 		PageNumber:  102,
	// 		StockNumber: 10,
	// 		StockID:     "09UJ",
	// 		Price:       7.8,
	// 		ISBN:        "9781853261589",
	// 		IsDeleted:   false,
	// 		Author:      Author{ID: "404", Name: "Antoine de Saint-Exupéry"},
	// 	},
	// 	{
	// 		ID:          "5",
	// 		Name:        "Dream of the Red Chamber",
	// 		PageNumber:  350,
	// 		StockNumber: 10,
	// 		StockID:     "77II",
	// 		Price:       17.0,
	// 		ISBN:        "9780385093798",
	// 		IsDeleted:   false,
	// 		Author:      Author{ID: "505", Name: "Cao Xueqin"},
	// 	}}

	// books := []Book{
	// 	{
	// 		ID:          "1",
	// 		Name:        "A Tale of Two Cities",
	// 		PageNumber:  320,
	// 		StockNumber: 10,
	// 		StockID:     "21AC",
	// 		Price:       15.3,
	// 		ISBN:        "9780451530578",
	// 		IsDeleted:   false,
	// 		AuthorID:    "101",
	// 		AuthorName:  "Charles Dickens",
	// 	},
	// 	{
	// 		ID:          "2",
	// 		Name:        "The Hobbit",
	// 		PageNumber:  376,
	// 		StockNumber: 10,
	// 		StockID:     "44UY",
	// 		Price:       24.0,
	// 		ISBN:        "9780547928227",
	// 		IsDeleted:   false,
	// 		AuthorID:    "202",
	// 		AuthorName:  "J. R. R. Tolkien",
	// 	},
	// 	{
	// 		ID:          "3",
	// 		Name:        "Harry Potter and the Philosophers Stone",
	// 		PageNumber:  560,
	// 		StockNumber: 10,
	// 		StockID:     "22OL",
	// 		Price:       32.2,
	// 		ISBN:        "9781408855898",
	// 		IsDeleted:   false,
	// 		AuthorID:    "303",
	// 		AuthorName:  "J. K. Rowling",
	// 	},
	// 	{
	// 		ID:          "4",
	// 		Name:        "The Little Prince",
	// 		PageNumber:  102,
	// 		StockNumber: 10,
	// 		StockID:     "09UJ",
	// 		Price:       7.8,
	// 		ISBN:        "9781853261589",
	// 		IsDeleted:   false,
	// 		AuthorID:    "404",
	// 		AuthorName:  "Antoine de Saint-Exupéry",
	// 	},
	// 	{
	// 		ID:          "5",
	// 		Name:        "Dream of the Red Chamber",
	// 		PageNumber:  350,
	// 		StockNumber: 10,
	// 		StockID:     "77II",
	// 		Price:       17.0,
	// 		ISBN:        "9780385093798",
	// 		IsDeleted:   false,
	// 		AuthorID:    "505",
	// 		AuthorName:  "Cao Xueqin",
	// 	}}

	books, _ := GetBooksWithWorkerPool(path)
	// fmt.Println(results)

	for _, book := range books {
		b.db.Where(Book{ID: book.ID}).Attrs(Book{ID: book.ID, Name: book.Name, PageNumber: book.PageNumber, StockNumber: book.StockNumber, StockID: book.StockID, Price: book.Price, ISBN: book.ISBN, IsDeleted: book.IsDeleted, AuthorID: book.AuthorID, AuthorName: book.AuthorName}).FirstOrCreate(&book)
	}
}

// func (b *BookRepository) InsertReadData(xxx <-chan Book) {
// 	// book := Book{}
// 	b.db.Where(Book{ID: book.ID}).Attrs(Book{ID: book.ID, Name: book.Name, PageNumber: book.PageNumber, StockNumber: book.StockNumber, StockID: book.StockID, Price: book.Price, ISBN: book.ISBN, IsDeleted: book.IsDeleted, AuthorID: book.AuthorID, AuthorName: book.AuthorName}).FirstOrCreate(&book)
// }

func (b *BookRepository) FindAll() []Book {
	books := []Book{}

	b.db.Find(&books)
	return books
}

func (b *BookRepository) FindByBookID(ID string) Book {
	book := Book{}
	b.db.Where(&Book{ID: ID}).First(&book)
	return book
}

func (b *BookRepository) FindByBookName(name string) []Book {
	books := []Book{}
	nameString := fmt.Sprintf("%%%s%%", name)

	b.db.Where("name ILIKE ?", nameString).Find(&books)
	return books
}

func (b *BookRepository) FindByAuthorName(name string) []Book {
	books := []Book{}
	nameString := fmt.Sprintf("%%%s%%", name)

	b.db.Where("author_name ILIKE ?", nameString).Find(&books)
	return books
}

func (b *BookRepository) DeleteByID(id string) (Book, error) {
	book := Book{}
	result := b.db.Delete(&Book{}, id)
	if result.Error != nil {
		return book, result.Error
	}
	// buraya return fonksiyonu eklenebilir
	b.db.Unscoped().Where(&Book{ID: id}).Find(&book)
	// fmt.Println(book)
	return book, nil
}

func (b *BookRepository) BuyByID(id string, num int) error {
	// book := Book{}
	book := b.FindByBookID(id)
	if book.StockNumber >= num {
		b.db.Model(&book).Update("stock_number", book.StockNumber-num)
	} else {
		return fmt.Errorf("Not enough stock")
	}
	// b.db.Model(&book).Where("stock_number >= ?", num).Update("stock_number", book.StockNumber-num)
	fmt.Println(book.StockNumber)
	return nil
}

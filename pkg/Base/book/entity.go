package book

import (
	"fmt"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	PageNumber  uint    `json:"pageNumber"`
	StockNumber int     `json:"stockNumber"`
	StockID     string  `json:"stockId"`
	Price       float32 `json:"price"`
	ISBN        string  `json:"isbn"`
	IsDeleted   bool    `json:"isdeleted"`
	AuthorID    string  `json:"authorid"`
	AuthorName  string  `json:"authorName"`
}

func (b *Book) ToString() string {
	return fmt.Sprintf("ID: %s, Name: %s, Page Number: %d, Stock Number: %d, StockID: %s, Price: %.2f, ISBN: %s, IsDeleted: %t, Author ID: %s, Author Name: %s", b.ID, b.Name, b.PageNumber, b.StockNumber, b.StockID, b.Price, b.ISBN, b.IsDeleted, b.AuthorID, b.AuthorName)
}

func (b *Book) BeforeDelete(tx *gorm.DB) error {
	fmt.Printf("Book %s is deleting...", b.Name)
	return nil
}
func (b *Book) AfterDelete(tx *gorm.DB) error {
	fmt.Printf("Book %s is deleted...", b.Name)
	return nil
}

// func (b *Book) SetStockNumber(number int) {
// 	b.StockNumber = number
// }
// func (b *Book) SetIsDeleted(isDeleted bool) {
// 	b.IsDeleted = isDeleted
// }

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

// func (b *Book) ToStruct(jobs <-chan []string, results chan<- Book, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for j := range jobs {
// 		pageNumberParsed, _ := strconv.Atoi(j[2])
// 		// if err != nil {
// 		// 	return err
// 		// }
// 		stockNumberParsed, _ := strconv.Atoi(j[3])
// 		// if err != nil {
// 		// 	return err
// 		// }
// 		priceParsed, _ := strconv.ParseFloat(j[5], 0)
// 		// if err != nil {
// 		// 	return err
// 		// }

// 		isDeletedParsed, _ := strconv.ParseBool(j[7])
// 		// if err != nil {
// 		// 	return err
// 		// }
// 		b := Book{ID: j[0],
// 			Name:        j[1],
// 			PageNumber:  uint(pageNumberParsed),
// 			StockNumber: stockNumberParsed,
// 			StockID:     j[4],
// 			Price:       float32(priceParsed),
// 			ISBN:        j[6],
// 			IsDeleted:   isDeletedParsed,
// 			AuthorID:    j[8],
// 			AuthorName:  j[9]}

// 		results <- b
// 	}

// }

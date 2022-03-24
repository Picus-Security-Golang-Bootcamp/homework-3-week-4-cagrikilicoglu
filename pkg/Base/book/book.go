package book

type Book struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	PageNumber  uint    `json:"pageNumber"`
	StockNumber int     `json:"stockNumber"`
	StockID     string  `json:"stockId"`
	Price       float32 `json:"price"`
	ISBN        string  `json:"isbn"`
	IsDeleted   bool    `json:"isdeleted"`
	Author      Author  `json:"author"`
}

type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (b *Book) SetStockNumber(number int) {
	b.StockNumber = number
}
func (b *Book) SetIsDeleted(isDeleted bool) {
	b.IsDeleted = isDeleted
}

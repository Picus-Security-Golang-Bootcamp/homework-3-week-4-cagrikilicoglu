package author

import (
	"bookApp/domain/book"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID    string
	Name  string
	Books []book.Book `gorm:"foreignKey:AuthorID;references:ID"`
}

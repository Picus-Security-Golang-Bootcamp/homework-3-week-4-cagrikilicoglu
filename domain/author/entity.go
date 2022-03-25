package author

import (
	"bookApp/domain/book"
	"fmt"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID    string
	Name  string
	Books []book.Book `gorm:"foreignKey:AuthorID;references:ID"`
}

// ToString: Convert author data into more readable string
func (a *Author) ToString() string {
	return fmt.Sprintf("ID: %s, Name: %s, Books: %v", a.ID, a.Name, a.Books)
}

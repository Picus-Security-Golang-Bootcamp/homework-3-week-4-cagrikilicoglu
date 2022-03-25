package author

import (
	"fmt"

	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (a *AuthorRepository) Migrations() {
	a.db.AutoMigrate(&Author{})
}

func (a *AuthorRepository) InsertAuthorData(path string) {
	// authors := []Author{
	// {
	// 	ID:   "101",
	// 	Name: "Charles Dickens",
	// },
	// {
	// 	ID:   "202",
	// 	Name: "J. R. R. Tolkien",
	// },
	// {
	// 	ID:   "303",
	// 	Name: "J. K. Rowling",
	// },
	// {
	// 	ID:   "404",
	// 	Name: "Antoine de Saint-Exupéry",
	// },
	// {
	// 	ID:   "505",
	// 	Name: "Cao Xueqin",
	// }}
	authors, _ := readAuthorsWithWorkerPool(path)
	for _, author := range authors {
		a.db.Where(Author{ID: author.ID}).Attrs(Author{ID: author.ID, Name: author.Name}).FirstOrCreate(&author)
	}

}
func (a *AuthorRepository) FindAuthorsWithBookInfo() ([]Author, error) {
	authors := []Author{}
	result := a.db.Preload("Books").Find(&authors)
	if result.Error != nil {
		return nil, result.Error
	}
	return authors, nil
}
func (a *AuthorRepository) FindAuthorsWithoutBookInfo() ([]Author, error) {
	authors := []Author{}
	result := a.db.Find(&authors)
	if result.Error != nil {
		return nil, result.Error
	}
	return authors, nil
}

func (a *AuthorRepository) FindByAuthorID(ID string) Author {
	author := Author{}
	a.db.Where(&Author{ID: ID}).First(&author)
	return author
}

func (a *AuthorRepository) FindByAuthorName(name string) []Author {
	authors := []Author{}
	nameString := fmt.Sprintf("%%%s%%", name)

	a.db.Where("name ILIKE ?", nameString).Find(&authors)
	return authors
}

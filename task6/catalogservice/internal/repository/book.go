package repository

import (
	"github.com/google/uuid"
	"time"
)

type Book struct {
	UUID        uuid.UUID   `database:"uuid"        json:"uuid"        example:"01234567-89ab-cdef-0123-456789abcdef"`
	Name        string      `database:"name"        json:"name"        example:"Clean code"`
	Author_uuid uuid.UUID   `database:"author_uuid" json:"author_uuid" example:"01234567-89ab-cdef-0123-456789abcdef"`

	Created_at  *time.Time   `database:"created_at"  json:"created_at" example:"2021-01-01T00:00:00Z"`
	Updated_at  *time.Time   `database:"updated_at"  json:"updated_at" example:"2021-01-01T00:00:00Z"`
	Deleted_at  *time.Time   `database:"deleted_at"  json:"deleted_at" swaggerignore:"true"`

	Categories  []*Category `json:"categories"`
}

type BookFilter struct {
	Page         uint64      `schema:"page"`
	PageSize     uint64      `schema:"per_page"`
	Categories   []uuid.UUID `schema:"categories"`
	Authors      []uuid.UUID `schema:"authors"`
	CategoryName string      `schema:"category_name"`
	AuthorName   string      `schema:"author_name"`
}

type BookRepository interface {
	Create(book *Book) (*Book, error)
	Get(uuid uuid.UUID) (*Book, error)
	List() ([]*Book, error)
	Search(filter *BookFilter) ([]*Book, bool, error)
	Update(book *Book) error
	Delete(book *Book) error

	AddCategory(book *Book, category *Category) error
	RemoveCategory(book *Book, category *Category) error
	FetchCategories(books []*Book) error
}

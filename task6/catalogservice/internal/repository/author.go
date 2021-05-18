package repository

import (
	"github.com/google/uuid"
	"time"
)

type Author struct {
	UUID uuid.UUID `database:"uuid" json:"uuid"`
	Name string    `database:"name" json:"name"`

	Created_at  *time.Time   `database:"created_at"  json:"created_at"`
	Updated_at  *time.Time   `database:"updated_at"  json:"updated_at"`
	Deleted_at  *time.Time   `database:"deleted_at"  json:"deleted_at" swaggerignore:"true"`
}

type AuthorRepository interface {
	Create(author *Author) (*Author, error)
	Get(uuid uuid.UUID) (*Author, error)
	List() ([]*Author, error)
	Update(author *Author) error
	Delete(author *Author) error
}

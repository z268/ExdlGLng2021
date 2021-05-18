package repository

import (
	"github.com/google/uuid"
	"time"
)

type Category struct {
	UUID        uuid.UUID  `database:"uuid" json:"uuid" example:"01234567-89ab-cdef-0123-456789abcdef"`
	Name        string     `database:"name" json:"name" example:"Programming"`
	Parent_uuid *uuid.UUID `database:"parent_uuid" json:"parent_uuid,omitempty" example:"01234567-89ab-cdef-0123-456789abcdef"`

	Created_at *time.Time `database:"created_at"  json:"created_at" example:"2021-01-01T00:00:00Z"`
	Updated_at *time.Time `database:"updated_at"  json:"updated_at" example:"2021-01-01T00:00:00Z"`
	Deleted_at *time.Time `database:"deleted_at"  json:"deleted_at" swaggerignore:"true"`
}

type CategoryRepository interface {
	Create(category *Category) (*Category, error)
	Get(uuid uuid.UUID) (*Category, error)
	List() ([]*Category, error)
	Update(author *Category) error
	Delete(category *Category) error
}
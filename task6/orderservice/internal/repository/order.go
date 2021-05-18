package repository

import (
	"time"
	"github.com/google/uuid"
)

type Order struct {
	UUID        uuid.UUID   `db:"uuid"        json:"uuid"        example:"01234567-89ab-cdef-0123-456789abcdef"`
	Book_uuid   uuid.UUID   `db:"book_uuid"   json:"book_uuid"   example:"01234567-89ab-cdef-0123-456789abcdef"`
	Description string      `db:"description" json:"description"`

	Created_at  *time.Time   `db:"created_at"  json:"created_at"`
	Updated_at  *time.Time   `db:"updated_at"  json:"updated_at"`
	Deleted_at  *time.Time   `db:"deleted_at"  json:"deleted_at" swaggerignore:"true"`
}

type OrderRepository interface {
	Create(order *Order) (*Order, error)
	Get(uuid uuid.UUID) (*Order, error)
	List() ([]*Order, error)
	Update(order *Order) error
	Delete(order *Order) error
}
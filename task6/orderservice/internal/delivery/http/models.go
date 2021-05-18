package http

import "github.com/google/uuid"

type UpdateOrderRequest struct {
	Description string  `db:"description" json:"description" example:"order description"`
}

type CreateOrderRequest struct {
	Book_uuid uuid.UUID `db:"book_uuid"   json:"book_uuid"   example:"01234567-89ab-cdef-0123-456789abcdef"`
	UpdateOrderRequest
}
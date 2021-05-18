package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/z268/ExdlGLng2021/task6/catalogservice/internal/repository"
)


type sqlRepository struct {
	db        *sqlx.DB
	tableName string
}

func NewCatalogRepository(db *sqlx.DB) *repository.CatalogRepository {
	return &repository.CatalogRepository{
		Author:   NewAuthorRepository(db),
		Book:     NewBookRepository(db),
		Category: NewCategoryRepository(db),
	}
}

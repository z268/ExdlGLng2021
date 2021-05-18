package mysql

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/z268/ExdlGLng2021/task6/catalogservice/internal/database"
	"github.com/z268/ExdlGLng2021/task6/catalogservice/internal/repository"
	"time"
)

const categoriesTableName = "categories"

type CategoryRepository struct {
	sqlRepository
}

func (r *CategoryRepository) Create(category *repository.Category) (*repository.Category, error) {
	_, err := sq.Insert(r.tableName).
		Columns("uuid", "name", "parent_uuid").
		Values(category.UUID, category.Name, category.Parent_uuid).
		RunWith(r.db).Query()
	return category, err
}

func (r *CategoryRepository) Get(uuid uuid.UUID) (*repository.Category, error) {
	sql, args, err :=  sq.Select("*").From(r.tableName).
		Where(sq.Eq{"uuid": uuid, "deleted_at": nil}).ToSql()
	if err != nil {
		return nil, err
	}

	category := repository.Category{}
	err = r.db.Get(&category, sql, args...)
	return &category, err
}

func (r *CategoryRepository) List() (categoryList []*repository.Category, err error) {
	sql, args, err :=  sq.Select("*").From(r.tableName).Where(sq.Eq{"deleted_at": nil}).ToSql()
	if err != nil {
		return nil, err
	}

	err = r.db.Select(&categoryList, sql, args...)
	return
}

func (r *CategoryRepository) Update(category *repository.Category) error {
	return database.CheckRowsAffected(
		sq.Update("categories").
			Set("name", category.Name).
			Set("parent_uuid", category.Parent_uuid).
			Where(sq.Eq{"uuid": category.UUID}).
			RunWith(r.db).Exec(),
	)
}

func (r *CategoryRepository) Delete(category *repository.Category) error {
	return database.CheckRowsAffected(
		sq.Update(r.tableName).Set("deleted_at", time.Now()).
			Where(sq.Eq{"uuid": category.UUID, "deleted_at": nil}).
			RunWith(r.db).Exec(),
	)
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{
		sqlRepository{db, categoriesTableName},
	}
}
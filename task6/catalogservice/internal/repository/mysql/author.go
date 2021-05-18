package mysql

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/z268/ExdlGLng2021/task6/catalogservice/internal/database"
	"github.com/z268/ExdlGLng2021/task6/catalogservice/internal/repository"
	"time"
)

const authorsTableName = "authors"

type AuthorRepository struct {
	sqlRepository
}

func (r *AuthorRepository) Create(author *repository.Author) (*repository.Author, error) {
	_, err := sq.
		Insert(r.tableName).Columns("uuid", "name").Values(author.UUID, author.Name).
		RunWith(r.db).Query()
	return author, err
}

func (r *AuthorRepository) Get(uuid uuid.UUID) (*repository.Author, error) {
	sql, args, err :=  sq.Select("*").From(r.tableName).
		Where(sq.Eq{"uuid": uuid, "deleted_at": nil}).ToSql()
	if err != nil {
		return nil, err
	}

	author := repository.Author{}
	err = r.db.Get(&author, sql, args...)
	return &author, err
}

func (r *AuthorRepository) List() ([]*repository.Author, error) {
	sql, args, err :=  sq.Select("*").From(r.tableName).Where(sq.Eq{"deleted_at": nil}).ToSql()
	if err != nil {
		return nil, err
	}

	authorList := []*repository.Author{}
	err = r.db.Select(&authorList, sql, args...)
	return authorList, err
}

func (r *AuthorRepository) Update(author *repository.Author) error {
	return database.CheckRowsAffected(
		sq.Update(r.tableName).Set("name", author.Name).
			Where(sq.Eq{"uuid": author.UUID}).
			RunWith(r.db).Exec(),
	)
}

func (r *AuthorRepository) Delete(author *repository.Author) error {
	return database.CheckRowsAffected(
		sq.Update(r.tableName).Set("deleted_at", time.Now()).
			Where(sq.Eq{"uuid": author.UUID, "deleted_at": nil}).
			RunWith(r.db).Exec(),
	)
}

func NewAuthorRepository(db *sqlx.DB) *AuthorRepository {
	return &AuthorRepository{
		sqlRepository{db, authorsTableName},
	}
}
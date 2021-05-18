package mysql

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/z268/ExdlGLng2021/task6/catalogservice/internal/database"
	"github.com/z268/ExdlGLng2021/task6/catalogservice/internal/repository"
	"time"
)

const booksTableName = "books"

type BookRepository struct {
	sqlRepository
}

func (r *BookRepository) Create(book *repository.Book) (*repository.Book, error) {
	_, err := sq.Insert(r.tableName).
		Columns("uuid", "name", "author_uuid").
		Values(book.UUID, book.Name, book.Author_uuid).
		RunWith(r.db).Query()
	return book, err
}

func (r *BookRepository) Get(uuid uuid.UUID) (*repository.Book, error) {
	sql, args, err :=  sq.Select("*").From(r.tableName).
		Where(sq.Eq{"uuid": uuid, "deleted_at": nil}).ToSql()
	if err != nil {
		return nil, err
	}

	book := repository.Book{}
	err = r.db.Get(&book, sql, args...)
	return &book, err
}

func (r *BookRepository) List() (books []*repository.Book, err error) {
	sql, args, err :=  sq.Select("*").From(r.tableName).Where(sq.Eq{"deleted_at": nil}).ToSql()
	if err != nil {
		return nil, err
	}

	err = r.db.Select(&books, sql, args...)
	return
}

func (r *BookRepository) Search(filter *repository.BookFilter) (books []*repository.Book, hasNext bool, err error) {
	query := sq.
		Select("b.*").
		From(fmt.Sprintf("%v as b", r.tableName)).
		Where(sq.Eq{"b.deleted_at": nil})

	if filter.Page > 1 {
		offset := filter.PageSize * (filter.Page - 1)
		query = query.Offset(offset)
	}

	// Filtering by category
	if len(filter.CategoryName) > 0 || len(filter.Categories) > 0 {
		query = query.Join("book_categories as bc ON (bc.book_uuid = b.uuid)")
	}
	if len(filter.Categories) > 0 {
		query = query.Where(sq.Eq{"bc.category_uuid": filter.Categories})
	}
	if len(filter.CategoryName) > 0 {
		query = query.Join("categories as c ON (c.uuid = bc.category_uuid)").
			Where(sq.Like{"c.name": fmt.Sprint("%", filter.CategoryName, "%")})
	}

	// Filtering by author
	if len(filter.Authors) > 0 {
		query = query.Where(sq.Eq{"b.author_uuid": filter.Authors})
	}
	if len(filter.AuthorName) > 0 {
		query = query.Join("authors as a ON (a.uuid = b.author_uuid)").
			Where(sq.Like{"a.name": fmt.Sprint("%", filter.AuthorName, "%")})
	}

	sql, args, err :=  query.Limit(filter.PageSize + 1).ToSql()
	err = r.db.Select(&books, sql, args...)

	if uint64(len(books)) > filter.PageSize {
		hasNext = true
		books = books[:filter.PageSize]
	}

	return
}

func (r *BookRepository) Update(book *repository.Book) error {
	return database.CheckRowsAffected(
		sq.Update("books").
			Set("name", book.Name).
			Set("author_uuid", book.Author_uuid).
			Where(sq.Eq{"uuid": book.UUID}).
			RunWith(r.db).Exec(),
	)
}

func (r *BookRepository) Delete(book *repository.Book) error {
	return database.CheckRowsAffected(
		sq.Update(r.tableName).Set("deleted_at", time.Now()).
			Where(sq.Eq{"uuid": book.UUID, "deleted_at": nil}).
			RunWith(r.db).Exec(),
	)
}

func NewBookRepository(db *sqlx.DB) *BookRepository {
	return &BookRepository{
		sqlRepository{db, booksTableName},
	}
}
package mysql

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/z268/ExdlGLng2021/task6/catalogservice/internal/database"
	"github.com/z268/ExdlGLng2021/task6/catalogservice/internal/repository"
)

const booksCategoriesTableName = "book_categories"

func (r *BookRepository) AddCategory(book *repository.Book, category *repository.Category) error {
	return database.CheckRowsAffected(
		sq.Insert(booksCategoriesTableName).
			Columns("book_uuid", "category_uuid").
			Values(book.UUID, category.UUID).
			RunWith(r.db).Exec(),
	)
}

func (r *BookRepository) RemoveCategory(book *repository.Book, category *repository.Category) error {
	return database.CheckRowsAffected(
		sq.Delete(booksCategoriesTableName).
			Where(sq.Eq{"book_uuid": book.UUID, "category_uuid": category.UUID}).
			RunWith(r.db).Exec(),
	)
}

func (r *BookRepository) FetchCategories(books []*repository.Book) error {
	booksIds := make([]uuid.UUID, 0, len(books))
	for _, book := range books {
		booksIds = append(booksIds, book.UUID)
	}

	sql, args, err := sq.
		Select("bc.book_uuid, c.*").From("categories as c").
		Join("book_categories as bc ON (bc.category_uuid = c.uuid)").
		Where(sq.Eq{"bc.book_uuid": booksIds}).ToSql()
	if err != nil {
		return err
	}

	rows, err := r.db.Queryx(sql, args...)
	if err != nil {
		return err
	}

	booksByUUID := make(map[uuid.UUID]*repository.Book)
	for _, book := range books {
		booksByUUID[book.UUID] = book
	}

	type m2mRelation struct{
		repository.Category
		Book_uuid uuid.UUID `json:"book_uuid" database:"book_uuid"`
	}
	for rows.Next() {
		rel := m2mRelation{}
		err = rows.StructScan(&rel)
		if book, ok := booksByUUID[rel.Book_uuid]; ok {
			book.Categories = append(book.Categories, &rel.Category)
		}
	}

	return err
}

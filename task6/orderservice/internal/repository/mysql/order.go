package mysql

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/z268/ExdlGLng2021/task6/orderservice/internal/repository"
	"github.com/z268/ExdlGLng2021/task6/orderservice/internal/database"
	"time"
)

const ordersTableName = "orders"

type sqlRepository struct {
	db        *sqlx.DB
	tableName string
}

type OrderRepository struct {
	sqlRepository
}

func (r *OrderRepository) Create(order *repository.Order) (*repository.Order, error) {
	_, err := sq.
		Insert(r.tableName).Columns("uuid", "book_uuid", "description").
		Values(order.UUID, order.Book_uuid, order.Description).
		RunWith(r.db).Query()
	return order, err
}

func (r *OrderRepository) Get(uuid uuid.UUID) (*repository.Order, error) {
	sql, args, err :=  sq.Select("*").From(r.tableName).
		Where(sq.Eq{"uuid": uuid, "deleted_at": nil}).ToSql()
	if err != nil {
		return nil, err
	}

	order := repository.Order{}
	err = r.db.Get(&order, sql, args...)
	return &order, err
}

func (r *OrderRepository) List() ([]*repository.Order, error) {
	sql, args, err :=  sq.Select("*").From(r.tableName).Where(sq.Eq{"deleted_at": nil}).ToSql()
	if err != nil {
		return nil, err
	}

	orderList := []*repository.Order{}
	err = r.db.Select(&orderList, sql, args...)
	return orderList, err
}

func (r *OrderRepository) Update(order *repository.Order) error {
	return database.CheckRowsAffected(
		sq.Update(r.tableName).Set("description", order.Description).
			Where(sq.Eq{"uuid": order.UUID}).
			RunWith(r.db).Exec(),
	)
}

func (r *OrderRepository) Delete(order *repository.Order) error {
	return database.CheckRowsAffected(
		sq.Update(r.tableName).Set("deleted_at", time.Now()).
			Where(sq.Eq{"uuid": order.UUID, "deleted_at": nil}).
			RunWith(r.db).Exec(),
	)
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{
		sqlRepository{db, ordersTableName},
	}
}

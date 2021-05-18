package database

import (
	"fmt"
	"errors"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/z268/ExdlGLng2021/task6/orderservice/internal/config"
)

func Init(cfg *config.Config) (db *sqlx.DB, err error) {
	source := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		cfg.Db.User, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port, cfg.Db.Database)
	return sqlx.Open(cfg.Db.Driver, source)
}

func Close(db *sqlx.DB) error {
	return db.Close()
}

func CheckRowsAffected(res sql.Result, err error) error {
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if rows == 0 {
		return errors.New("No rows affected")
	}

	return err
}

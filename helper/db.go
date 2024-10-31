package helper

import (
	"database/sql"
	"errors"
)

type TxAware[T any] interface {
	WithTx(tx *sql.Tx) *T
}

func Transaction[T any, R TxAware[T]](db *sql.DB, repository R, fn func(qtx *T) error) error {
	tx, err := db.Begin()
	if err != nil {
		return errors.New("创建事务失败")
	}
	defer tx.Rollback()

	qtx := repository.WithTx(tx)

	err = fn(qtx)
	if err != nil {
		return errors.New(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		return errors.New("事务提交失败")
	}

	return err
}

package db

import (
	"database/sql"
	"errors"
	"time"
)

func EasyQuery(query string, args ...any) (*sql.Rows, error) {
	res, err := Db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func EasyExec(query string, args ...any) (sql.Result, error) {
	res, err := Db.Exec(query, args...)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func RetriedDbQuery(retries int, query string, args ...any) (*sql.Rows, error) {
	for range retries {
		res, err := EasyQuery(query, args...)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		} else {
			return res, err
		}
	}

	return nil, errors.New("failed")
}

func RetriedDbExec(retries int, query string, args ...any) (sql.Result, error) {
	for range retries {
		res, err := EasyExec(query, args...)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		} else {
			return res, err
		}
	}

	return nil, errors.New("failed")
}

package db

import (
	"database/sql"
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

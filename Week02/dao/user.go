package dao

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type User struct {
	ID int
}

var ErrNotFound = errors.New("record not found")
var ErrDBServer = errors.New("DB error")

func QueryUserByID(ID int) (*User, error) {

	sqls := "select id from img_ocr_result where id = ?"
	row := MysqlDB.QueryRow(sqls, ID)
	resUser := User{}
	if err := row.Scan(&resUser.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrapf(ErrNotFound, fmt.Sprintf("sql: %s error(%v)", sqls, err))
		}
		return nil, errors.Wrapf(ErrDBServer, fmt.Sprintf("query: %s error(%v)", sqls, err))

	}
	return &resUser, nil
}

package dao

import (
	"aaron/go/learning/homework1/model"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func QueryAll(db *sql.DB) ([]model.User, error) {
	defer db.Close()
	users := make([]model.User, 0)
	rows, err := db.Query("SLELECT id, name, email from user")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func IsDBNotFound(err error) bool {
	_, ok := err.(sql.ErrNoRows)
	return ok
}

func IsDBConnErr(err error) bool {
	_, ok := err.(sql.ErrConnDone)
	return ok
}

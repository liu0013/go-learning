package service

import (
	"aaron/go/learning/homework1/dao"
	"aaron/go/learning/homework1/model"
	"fmt"

	"github.com/pkg/errors"
)

func ListUser(userDao dao.UserDao) ([]model.User, error) {
	users, err := userDao.QueryAll()
	if err != nil {
		if dao.IsDBConnErr(err) {
			fmt.Println("query all is db connection error: ", err.Error())
			// TODO retry to connect db and query again
		} else if dao.IsDBNotFound(err) {
			return nil, nil
		} else {
			return nil, errors.Wrap(err, "query all db error")
		}
	}

	return users, nil
}

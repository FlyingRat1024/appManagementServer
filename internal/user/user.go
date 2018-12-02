package user

import (
	"androidappServer/db"
	"androidappServer/pkg/utils"
)

type ModifyPasswordBody struct {
	UserID   int    `json:"user_id"`
	Password string `json:"password"`
}

func QueryUserInfo(userID int) (string, error) {
	db, err := db.GetDB()
	if err != nil {
		return "", err
	}
	sqlfmt := "select id, employee_id, employee_name, sex, position, create_time from user where id = ? "
	stmt, err := db.Prepare(sqlfmt)
	if err != nil {
		return "", err
	}
	rows, err := stmt.Query(userID)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	jsonStr, err := utils.SqlRows2Json(rows)
	if err != nil {
		return "", err
	}
	return jsonStr, nil
}

func ModifyPassword(param *ModifyPasswordBody) error {
	db, err := db.GetDB()
	if err != nil {
		return err
	}
	sqlfmt := "update user set password = ? where id = ? "
	stmt, err := db.Prepare(sqlfmt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(param.Password, param.UserID)
	return err
}

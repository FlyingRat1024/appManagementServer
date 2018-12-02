package login

import (
	mysql "androidappServer/db"
	"androidappServer/pkg/utils"
)

func Login(employeeID string, password string) (string, error) {
	db, err := mysql.GetDB()
	if err != nil {
		return "", err
	}
	defer db.Close()
	stmt, err := db.Prepare("select id, employee_id, employee_name, sex, position  from user where employee_id = ? and password = ?")
	if err != nil {
		return "", err
	}
	rows, err := stmt.Query(employeeID, password)
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

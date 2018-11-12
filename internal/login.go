package internal

import mysql "androidappServer/db"

func Login(employeeID string, password string) (bool, error) {
	db, err := mysql.GetDB()
	if err != nil {
		return false, err
	}
	defer db.Close()
	stmt, err := db.Prepare("select id from user where employee_id = ? and password = ?")
	if err != nil {
		return false, err
	}
	rows, err := stmt.Query(employeeID, password)
	if rows.Next() {
		return true, nil
	}
	defer rows.Close()
	return false, nil
}

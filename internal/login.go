package internal

import mysql "androidappServer/db"

func Login(user string, password string) (bool, error) {
	db, err := mysql.GetDB()
	if err != nil {
		return false, err
	}
	defer db.Close()
	stmt, err := db.Prepare("select id from user where employee_id = ? and password = ?")
	if err != nil {
		return false, err
	}
	result, err := stmt.Query(user, password)
	if result.Next() {
		return true, nil
	}
	return false, nil
}

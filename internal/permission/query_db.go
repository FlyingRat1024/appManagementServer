package permission

import (
	"androidappServer/db"
)

func QueryPermission(role string) ([]string, error) {
	mysql, err := db.GetDB()
	if err != nil {
		return nil, err
	}
	defer mysql.Close()
	sqlfmt := "select permission from permission where role = ?"
	stmt, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(role)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	permission := ""
	var permissions []string
	for rows.Next() {
		if err := rows.Scan(&permission); err == nil {
			permissions = append(permissions, permission)
		}
	}
	return permissions, nil
}

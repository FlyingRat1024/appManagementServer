package warehouse

import (
	"androidappServer/db"
)

func CreateInWarehouseTable(table *InWarehouseTableBody) error {
	mysql, err := db.GetDB()
	if err != nil {
		return err
	}
	defer mysql.Close()
	//事务
	conn, err := mysql.Begin()
	if err != nil {
		return err
	}
	sqlfmt := "insert into warehouse_in (writer, reissue, create_time, status) values(?,?,now(),'pending')"
	stmt, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(table.Writer, table.ReIssue)
	if err != nil {
		return err
	}
	tableID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	for _, material := range table.Material {
		sqlfmt = "insert into in_material (in_id, material_id, number) values(?,?,?)"
		stmt, err = mysql.Prepare(sqlfmt)
		if err != nil {
			conn.Rollback()
			return err
		}
		_, err = stmt.Exec(tableID, material.MaterialID, material.Num)
		if err != nil {
			conn.Rollback()
			return err
		}
	}
	err = conn.Commit()
	if err != nil {
		return err
	}
	return nil
}

func CreateOutWarehouseTable(table *OutWarehouseTableBody) error {
	mysql, err := db.GetDB()
	if err != nil {
		return err
	}
	defer mysql.Close()
	//事务
	conn, err := mysql.Begin()
	if err != nil {
		return err
	}
	sqlfmt := "insert into warehouse_out (writer, create_time, status) values(?, now(), 'pending')"
	stmt, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(table.Writer)
	if err != nil {
		return err
	}
	tableID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	for _, material := range table.Material {
		sqlfmt = "insert into out_material (out_id, material_id, number) values(?,?,?)"
		stmt, err = mysql.Prepare(sqlfmt)
		if err != nil {
			conn.Rollback()
			return err
		}
		_, err = stmt.Exec(tableID, material.MaterialID, material.Num)
		if err != nil {
			conn.Rollback()
			return err
		}
	}
	err = conn.Commit()
	if err != nil {
		return err
	}
	return nil
}

func ConfirmOutWarehouseTable(body *ConfirmOutWarehouseBody) error {
	mysql, err := db.GetDB()
	if err != nil {
		return err
	}
	defer mysql.Close()
	sqlfmt := "update warehouse_out set verifier = ?, verify_time = now() where id = ?"
	stmt, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(body.ConfirmUser, body.TableID)
	return err
}

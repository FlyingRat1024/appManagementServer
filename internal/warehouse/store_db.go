package warehouse

import (
	"androidappServer/db"
	"fmt"
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
	sqlfmt := "insert into project (name) values(?)"
	stmt, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(table.ProjectName)
	if err != nil {
		return err
	}
	projectID, err := result.LastInsertId()
	if err != nil{
		return err
	}
	sqlfmt = "insert into warehouse_in (project_id, table_num, writer, reissue, create_time, status) values(?,?,?,?,now(),'未确认')"
	stmt, err = mysql.Prepare(sqlfmt)
	if err != nil {
		conn.Rollback()
		return err
	}
	result, err = stmt.Exec(projectID, table.TableNum, table.Writer, table.ReIssue)
	if err != nil {
		conn.Rollback()
		return err
	}
	tableID, err := result.LastInsertId()
	if err != nil {
		conn.Rollback()
		return err
	}
	for _, material := range table.Material {
		if material.MaterialID == 0 || material.Num == 0{
			conn.Rollback()
			return fmt.Errorf("request param error, please check your json")
		}
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
	sqlfmt := "insert into project (name) values(?)"
	stmt, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(table.ProjectName)
	if err != nil {
		return err
	}
	projectID, err := result.LastInsertId()
	if err != nil{
		return err
	}
	sqlfmt = "insert into warehouse_out (project_id, table_num, writer, create_time, status) values(?,?,?, now(), '未确认')"
	stmt, err = mysql.Prepare(sqlfmt)
	if err != nil {
		conn.Rollback()
		return err
	}
	result, err = stmt.Exec(projectID, table.TableNum, table.Writer)
	if err != nil {
		conn.Rollback()
		return err
	}
	tableID, err := result.LastInsertId()
	if err != nil {
		conn.Rollback()
		return err
	}
	for _, material := range table.Material {
		if material.MaterialID == 0 || material.Num == 0{
			conn.Rollback()
			return fmt.Errorf("request param error, please check your json")
		}
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
	sqlfmt := "update warehouse_out set verifier = ?, status = '已确认', verify_time = now() where id = ?"
	stmt, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(body.ConfirmUser, body.TableID)
	return err
}

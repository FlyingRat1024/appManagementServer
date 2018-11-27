package material

import "androidappServer/db"

// 填写申请表
func CreateApplyTable(body *ApplyTableBody) error {
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
	sqlfmt := "insert into material_apply_table (user_id, create_time) values(?, now())"
	stmt, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(body.ApplierID)
	if err != nil {
		return err
	}
	tableID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	sqlfmt = "insert into apply_material (table_id,material_id, num) values(?,?,?)"
	stmt, err = mysql.Prepare(sqlfmt)
	if err != nil {
		conn.Rollback()
		return err
	}
	_, err = stmt.Exec(tableID, body.MaterialID, body.Num)
	if err != nil {
		conn.Rollback()
		return err
	}
	//
	err = conn.Commit()
	if err != nil {
		return err
	}
	return nil
}

// 填写领料表
func CreateReceiveTable(body *RecieveTableBody) error {
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
	sqlfmt := "insert into material_receive_table (receiver, create_time) values(?, now())"
	stmt, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(body.ReciiverID)
	if err != nil {
		return err
	}
	tableID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	sqlfmt = "insert into receive_material (table_id, material_id, receive_num) values(?,?,?)"
	stmt, err = mysql.Prepare(sqlfmt)
	if err != nil {
		conn.Rollback()
		return err
	}
	_, err = stmt.Exec(tableID, body.MaterialID, body.Num)
	if err != nil {
		conn.Rollback()
		return err
	}
	//
	err = conn.Commit()
	if err != nil {
		return err
	}
	return nil
}

// 填写归还表
func CreateBackTable(body *BackTableBody) error {
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
	sqlfmt := "update material_receive_table set back_user = ?, back = 1, back_time = now() where id = ?"
	stmt, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(body.BackerID, body.TableID)
	if err != nil {
		return err
	}
	sqlfmt = "update receive_material set back_num = ? where table_id = ? and material_id = ?"
	stmt, err = mysql.Prepare(sqlfmt)
	if err != nil {
		conn.Rollback()
		return err
	}
	_, err = stmt.Exec(body.Num, body.TableID, body.MaterialID)
	if err != nil {
		conn.Rollback()
		return err
	}
	//
	err = conn.Commit()
	if err != nil {
		return err
	}
	return nil
}

// 填写质检表
func CreateCheckTable(body *CheckTableBody) error {
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
	sqlfmt := "update material_receive_table set checker = ?, `check` = 1, check_time = now() where id = ?"
	stmt, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(body.CheckerID, body.TableID)
	if err != nil {
		return err
	}
	sqlfmt = "update receive_material set check_num = ? where table_id = ? and material_id = ?"
	stmt, err = mysql.Prepare(sqlfmt)
	if err != nil {
		conn.Rollback()
		return err
	}
	_, err = stmt.Exec(body.Num, body.TableID, body.MaterialID)
	if err != nil {
		conn.Rollback()
		return err
	}
	//
	err = conn.Commit()
	if err != nil {
		return err
	}
	return nil
}

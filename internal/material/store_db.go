package material

import (
	"androidappServer/db"
	"fmt"
)

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
	sqlfmt := "insert into project (name) values(?)"
	stmt, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(body.ProjectName)
	if err != nil {
		return err
	}
	projectID, err := result.LastInsertId()
	if err != nil{
		return err
	}
	status := "pending"
	sqlfmt = "insert into material_apply_table (user_id, project_id, table_num, create_time, verify) " +
		"values(?, ?, ?, now(), ?)"
	stmt, err = mysql.Prepare(sqlfmt)
	if err != nil {
		conn.Rollback()
		return err
	}
	result, err = stmt.Exec(body.ApplierID, projectID, body.TableNum , status)
	if err != nil {
		conn.Rollback()
		return err
	}
	tableID, err := result.LastInsertId()
	if err != nil {
		conn.Rollback()
		return err
	}
	for _, material := range body.Material {
		sqlfmt = "insert into apply_material (table_id,material_id, num) values(?,?,?)"
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
	sqlfmt := "insert into project (name) values(?)"
	stmt, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(body.ProjectName)
	if err != nil {
		return err
	}
	projectID, err := result.LastInsertId()
	if err != nil{
		return err
	}
	sqlfmt = "insert into material_receive_table (receiver, table_num, project_id, create_time, verify) values(?, ?, ?, now(), 'pending')"
	stmt, err = mysql.Prepare(sqlfmt)
	if err != nil {
		conn.Rollback()
		return err
	}
	result, err = stmt.Exec(body.ReciiverID, body.TableNum, projectID)
	if err != nil {
		conn.Rollback()
		return err
	}
	tableID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	for _, material := range body.Material {
		if material.MaterialID == 0 || material.Num == 0{
			conn.Rollback()
			return fmt.Errorf("request param error, please check your json")
		}
		sqlfmt = "insert into receive_material (table_id, material_id, receive_num) values(?,?,?)"
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
	sqlfmt := "update material_receive_table set back_table_num = ?, back_user = ?, back = '已归还', back_time = now() " +
		"where id = ?"
	stmt, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(body.TableNum, body.BackerID, body.TableID)
	if err != nil {
		return err
	}
	for _, material := range body.Material {
		if material.MaterialID == 0 || material.Num == 0{
			conn.Rollback()
			return fmt.Errorf("request param error, please check your json")
		}
		sqlfmt = "update receive_material set back_num = ? where table_id = ? and material_id = ?"
		stmt, err = mysql.Prepare(sqlfmt)
		if err != nil {
			conn.Rollback()
			return err
		}
		_, err = stmt.Exec(material.Num, body.TableID, material.MaterialID)
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


// 填写废料归还表
func CreateWasteBackTable(body *BackTableBody) error {
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
	sqlfmt := "update material_receive_table set waste_back_table_num = ?, waste_backer = ?, waste_back = '已归还', " +
		"waste_back_time = now() where id = ?"
	stmt, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(body.TableNum, body.BackerID, body.TableID)
	if err != nil {
		return err
	}
	for _, material := range body.Material {
		if material.MaterialID == 0 || material.Num == 0{
			conn.Rollback()
			return fmt.Errorf("request param error, please check your json")
		}
		sqlfmt = "update receive_material set waste_back_num = ? where table_id = ? and material_id = ?"
		stmt, err = mysql.Prepare(sqlfmt)
		if err != nil {
			conn.Rollback()
			return err
		}
		_, err = stmt.Exec(material.Num, body.TableID, material.MaterialID)
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
	sqlfmt := "update material_receive_table set check_table_num = ?, checker = ?, `check` = '已质检', " +
		"check_time = now() where id = ?"
	stmt, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(body.TableNum, body.CheckerID, body.TableID)
	if err != nil {
		return err
	}
	for _, material := range body.Material {
		if material.MaterialID == 0 || material.Num == 0{
			conn.Rollback()
			return fmt.Errorf("request param error, please check your json")
		}
		sqlfmt = "update receive_material set check_num = ? where table_id = ? and material_id = ?"
		stmt, err = mysql.Prepare(sqlfmt)
		if err != nil {
			conn.Rollback()
			return err
		}
		_, err = stmt.Exec(material.Num, body.TableID, material.MaterialID)
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

// 新增材料
func CreateMaterial(param *Material) error {
	mysql, err := db.GetDB()
	if err != nil {
		return err
	}
	defer mysql.Close()
	sqlfmt := "insert into material(name, size, unit, provider, description) values(?,?,?,?,?)"
	stmt, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(param.Name, param.Size, param.Unit, param.Provider, param.Description)
	if err != nil {
		return err
	}
	return nil
}

// 经理审核申请表, 修改status
func ModifyApplyStatus(param *VerifyBody) error {
	mysql, err := db.GetDB()
	if err != nil {
		return err
	}
	defer mysql.Close()
	sqlfmt := "update material_apply_table set verify = ?, verifier = ?, verify_time = now() where id = ?"
	stmt, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(param.Status, param.Verifier, param.TableID)
	if err != nil {
		return err
	}
	/*
	if ok, _ := result.RowsAffected(); ok != 1{
		return fmt.Errorf("modify database error or has no change")
	}
	*/
	return nil
}

// 经理审核领料表
func ModifyReceiveStatus(param *VerifyBody) error {
	mysql, err := db.GetDB()
	if err != nil {
		return err
	}
	defer mysql.Close()
	sqlfmt := "update material_receive_table set verify = ?, verifier = ?, verify_time= now() where id = ?"
	stmt, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(param.Status, param.Verifier, param.TableID)
	if err != nil {
		return err
	}
	/*
	if ok, _ := result.RowsAffected(); ok != 1{
		return fmt.Errorf("modify database failed or no modify")
	}
	*/
	return nil
}
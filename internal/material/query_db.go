package material

import (
	"androidappServer/db"
	"androidappServer/pkg/utils"
	"fmt"
)

// 查询申请列表
func QueryApplyList(userID string) (string, error) {
	mysql, err := db.GetDB()
	if err != nil {
		return "", err
	}
	defer mysql.Close()
	if userID == "" {
		sqlfmt := "select id as table_id, (select employee_name from user where id = material_apply_table.user_id) as writer, " +
			"verify as status, create_time from material_apply_table"
		stmt, err := mysql.Prepare(sqlfmt)
		if err != nil {
			return "", err
		}
		rows, err := stmt.Query()
		if err != nil {
			return "", err
		}
		defer rows.Close()
		jsonStr, err := utils.SqlRows2JsonList(rows)
		return jsonStr, err
	} else {
		sqlfmt := "select id as table_id, (select employee_name from user where id = material_apply_table.user_id) as writer, " +
			"verify as status, create_time from material_apply_table where user_id = ?"
		stmt, err := mysql.Prepare(sqlfmt)
		if err != nil {
			return "", err
		}
		rows, err := stmt.Query(userID)
		if err != nil {
			return "", err
		}
		defer rows.Close()
		jsonStr, err := utils.SqlRows2JsonList(rows)
		return jsonStr, err
	}

}

// 查询申请单详细信息
func QueryApplyDetail(tableID int) (string, error) {
	mysql, err := db.GetDB()
	if err != nil {
		return "", err
	}
	defer mysql.Close()
	sqlfmt := "select id as table_id, (select employee_name from user where id = material_apply_table.user_id) as writer, " +
		"verify as status, create_time from material_apply_table where id = ?"
	stmt, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return "", err
	}
	rows, err := stmt.Query(tableID)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	resultMap, err := utils.SqlRows2Map(rows)
	if err != nil {
		return "", err
	}
	sqlfmt = "select material.name, material.unit, material.provider, apply.num " +
		"from (select material_id, num from apply_material where id = ?) as apply " +
		"JOIN material ON material.id = apply.material_id"
	stmt, err = mysql.Prepare(sqlfmt)
	if err != nil {
		return "", err
	}
	rows1, err := stmt.Query(tableID)
	if err != nil {
		return "", err
	}
	defer rows1.Close()
	materialJson, err := utils.SqlRows2JsonList(rows1)
	if err != nil {
		return "", err
	}
	if materialJson == "[]"{
		return "", fmt.Errorf("can't find this table")
	}
	resultMap["material"] = materialJson
	jsonStr, err := utils.ToString(resultMap)
	return jsonStr, err
}

// 查询领料单列表
func QueryReceiveTableList(userID string) (string, error) {
	mysql, err := db.GetDB()
	if err != nil {
		return "", err
	}
	defer mysql.Close()
	if userID == "" {
		sqlfmt := "select id as table_id, (select employee_name from user where id = material_receive_table.receiver) as writer, " +
			"verify as status, create_time from material_receive_table"
		stmt, err := mysql.Prepare(sqlfmt)
		if err != nil {
			return "", err
		}
		rows, err := stmt.Query()
		if err != nil {
			return "", err
		}
		defer rows.Close()
		jsonStr, err := utils.SqlRows2JsonList(rows)
		return jsonStr, err
	} else {
		sqlfmt := "select id as table_id, (select employee_name from user where id = material_receive_table.receiver) as writer, " +
			"verify as status, create_time from material_receive_table where receiver = ?"
		stmt, err := mysql.Prepare(sqlfmt)
		if err != nil {
			return "", err
		}
		rows, err := stmt.Query(userID)
		if err != nil {
			return "", err
		}
		defer rows.Close()
		jsonStr, err := utils.SqlRows2JsonList(rows)
		return jsonStr, err
	}

}

// 领料单详细信息
func QueryReceiveDetail(tableID int) (string, error) {
	mysql, err := db.GetDB()
	if err != nil {
		return "", err
	}
	defer mysql.Close()
	sqlfmt := "select id, (select employee_name from user where id = receiver) as writer, create_time, status, " +
		"verify, (select employee_name from user where id = verifier) as verifier, verify_time, " +
		"`check`, (select employee_name from user where id = checker) as checker, check_time, " +
		"back, (select employee_name from user where id = back_user) as backer, back_time " +
		"from material_receive_table where id = ?"
	stmt, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return "", err
	}
	rows, err := stmt.Query(tableID)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	resultMap, err := utils.SqlRows2Map(rows)
	if err != nil {
		return "", err
	}
	sqlfmt = "select material.name, material.unit, material.provider, " +
		"receive.receive_num as num, receive.back_num, receive.check_num " +
		"from (select material_id, receive_num, check_num, back_num from receive_material where id = ?) as receive " +
		"JOIN material ON material.id = receive.material_id"
	stmt1, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return "", err
	}
	rows1, err := stmt1.Query(tableID)
	if err != nil {
		return "", err
	}
	defer rows1.Close()
	materialJson, err := utils.SqlRows2JsonList(rows1)
	if err != nil {
		return "", err
	}
	if materialJson == "[]"{
		return "", fmt.Errorf("can't find this table")
	}
	resultMap["material"] = materialJson
	jsonStr, err := utils.ToString(resultMap)
	return jsonStr, err
}

// 获取材料列表
func GetMaterialList() (string, error) {
	db, err := db.GetDB()
	if err != nil {
		return "", err
	}
	sqlfmt := "select id, name, unit, provider, description from material "
	stmt, err := db.Prepare(sqlfmt)
	if err != nil {
		return "", err
	}
	rows, err := stmt.Query()
	if err != nil {
		return "", err
	}
	defer rows.Close()
	jsonStr, err := utils.SqlRows2JsonList(rows)
	if err != nil {
		return "", err
	}
	return jsonStr, nil
}

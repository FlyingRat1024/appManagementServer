package warehouse

import (
	"androidappServer/db"
	"androidappServer/pkg/utils"
	"fmt"
)

//查询入库单列表
func QueryInWarehouseList(userID string) (string, error) {
	mysql, err := db.GetDB()
	if err != nil {
		return "", err
	}
	defer mysql.Close()
	if userID == ""{
		sqlfmt := "select id as table_id, " +
			"(select employee_name from user where id = warehouse_in.writer) as writer, " +
			"create_time, reissue from warehouse_in"
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
	}else {
		sqlfmt := "select id as table_id, " +
			"(select employee_name from user where id = warehouse_in.writer) as writer, " +
			"create_time, reissue from warehouse_in where writer = ?"
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

//查询入库单详细信息
func QueryInWarehouseDetail(tableID int) (string, error) {
	mysql, err := db.GetDB()
	if err != nil {
		return "", err
	}
	defer mysql.Close()
	sqlfmt := "select id as table_id, " +
		"(select employee_name from user where id = warehouse_in.writer) as writer, " +
		"create_time, reissue, status from warehouse_in where id = ?"
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
	if len(resultMap) == 0{
		return "", fmt.Errorf("can't find this table")
	}
	sqlfmt = "select material.name, material.unit, material.provider, in_material.number " +
		"from (select material_id, number from in_material where in_id = ?) as in_material " +
		"JOIN material ON material.id = in_material.material_id"
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
	resultMap["material"] = materialJson
	jsonStr, err := utils.ToString(resultMap)
	return jsonStr, err
}

//查询出库单列表
func QueryOutWarehouseList(userID string) (string, error) {
	mysql, err := db.GetDB()
	if err != nil {
		return "", err
	}
	defer mysql.Close()
	if userID == ""{
		sqlfmt := "select id as table_id, " +
			"(select employee_name from user where id = warehouse_out.writer) as writer, " +
			"create_time from warehouse_out"
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
	}else{
		sqlfmt := "select id as table_id, " +
			"(select employee_name from user where id = warehouse_out.writer) as writer, " +
			"create_time from warehouse_out where writer = ?"
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

//查询出库单详细信息
func QueryOutWarehouseDetail(tableID int) (string, error) {
	mysql, err := db.GetDB()
	if err != nil {
		return "", err
	}
	defer mysql.Close()
	sqlfmt := "select id as table_id, " +
		"(select employee_name from user where id = warehouse_out.writer) as writer, " +
		"create_time, status from warehouse_out where id = ?"
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
	if len(resultMap) == 0{
		return "", fmt.Errorf("can't find this table")
	}
	sqlfmt = "select material.name, material.unit, material.provider, out_material.number " +
		"from (select material_id, number from out_material where out_id = ?) as out_material " +
		"JOIN material ON material.id = out_material.material_id"
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
	resultMap["material"] = materialJson
	jsonStr, err := utils.ToString(resultMap)
	return jsonStr, err
}

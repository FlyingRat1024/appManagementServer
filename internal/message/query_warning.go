package message

import (
	"androidappServer/db"
	"androidappServer/pkg/utils"
)

// 获取材料列表
func GetWarning() (string, error) {
	db, err := db.GetDB()
	if err != nil {
		return "", err
	}
	sqlfmt := "select table_count.material_name, table_count.receive_num, table_count.check_num, table_count.back_num, " +
		"table_info.writer, table_info.checker, table_info.backer, table_info.create_time " +
		"from " +
		"(select table_id, " +
		"(select name from material where id = material_id) as material_name, " +
		"receive_num, check_num, back_num  from receive_material where receive_num <> check_num + back_num )" +
		"AS table_count " +
		"JOIN " +
		"( select id, create_time, " +
		"(select employee_name from user where id = receiver) as writer, " +
		"(select employee_name from user where id = checker) as checker, " +
		"(select employee_name from user where id = back_user) as backer " +
		"from material_receive_table " +
		"where verify=1 and `check`=1 and back=1 ) " +
		"AS table_info " +
		"ON table_count.table_id = table_info.id"
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

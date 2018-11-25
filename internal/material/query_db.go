package material

import (
	"androidappServer/db"
	"androidappServer/internal/utils"
)

func QueryApplyList() (string, error) {
	mysql, err := db.GetDB()
	if err != nil {
		return "", err
	}
	sqlfmt := "SELECT apply.table_id, apply.create_time, user.employee_name as applier, apply.material_name " +
		"FROM (SELECT material_apply_table.id as table_id,material_apply_table.user_id, " +
		"material_apply_table.create_time, material.name as material_name " +
		"FROM apply_material JOIN material_apply_table ON material_apply_table.id = apply_material.table_id " +
		"JOIN material ON material.id = apply_material.material_id ) " +
		"AS apply JOIN user ON user.id = apply.user_id "
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
}

func QueryApplyDetail(tableID int) (string, error){
	mysql, err := db.GetDB()
	if err != nil {
		return "", err
	}
	sqlfmt := "SELECT apply.table_id, apply.create_time, user.employee_name as applier, apply.material_name, " +
		"apply.material_unit, apply.material_provider, apply.num " +
		"FROM (SELECT apply_table.id as table_id,apply_table.user_id, " +
		"apply_table.create_time, material.name as material_name, apply_material.num, " +
		"material.unit as material_unit, material.provider as material_provider " +
		"FROM apply_material " +
		"JOIN (select id,user_id, create_time from material_apply_table where id = ?) as apply_table " +
		"ON apply_table.id = apply_material.table_id " +
		"JOIN material ON material.id = apply_material.material_id ) " +
		"AS apply JOIN user ON user.id = apply.user_id "
	stmt, err := mysql.Prepare(sqlfmt)
	if err != nil {
		return "", err
	}
	rows, err := stmt.Query(tableID)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	jsonStr, err := utils.SqlRows2Json(rows)
	return jsonStr, err
}
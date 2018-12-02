package warehouse

import "androidappServer/pkg/structs"

type InWarehouseTableBody struct {
	Writer int `json:"writer"`
	Number int `json:"num"`
	structs.Material
}

type OutWarehouseTableBody struct {
	Writer     int `json:"writer"`
	MaterialID int `json:"material_id"`
}

type ConfirmOutWarehouseBody struct {
	ConfirmUser int `json:"confirm_user"`
	TableID     int `json:"table_id"`
}

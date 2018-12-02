package warehouse

type InWarehouseTableBody struct {
	Writer   int      `json:"writer"`
	ReIssue  int      `json:"reissue"`
	Material []Common `json:"material"`
}

type OutWarehouseTableBody struct {
	Writer   int      `json:"writer"`
	Material []Common `json:"material"`
}

type ConfirmOutWarehouseBody struct {
	ConfirmUser int `json:"confirm_user"`
	TableID     int `json:"table_id"`
}

type Common struct {
	MaterialID int `json:"material_id"`
	Num        int `json:"num"`
}

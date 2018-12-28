package warehouse

type InWarehouseTableBody struct {
	TableNum    string   `json:"table_num"`
	ProjectName string   `json:"project_name"`
	Writer      int      `json:"writer"`
	ReIssue     int      `json:"reissue"`
	Material    []Common `json:"material"`
}

type OutWarehouseTableBody struct {
	TableNum    string   `json:"table_num"`
	ProjectName string   `json:"project_name"`
	Writer      int      `json:"writer"`
	Material    []Common `json:"material"`
}

type ConfirmOutWarehouseBody struct {
	ConfirmUser int `json:"confirm_user"`
	TableID     int `json:"table_id"`
}

type Common struct {
	MaterialID int `json:"id"`
	Num        int `json:"num"`
}

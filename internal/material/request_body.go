package material

type ApplyTableBody struct {
	MaterialID int `json:"material_id"`
	ApplierID  int `json:"applier"`
	Num        int `json:"num"`
}

type RecieveTableBody struct {
	ReciiverID int `json:"receiver"`
	MaterialID int `json:"material_id"`
	Num        int `json:"num"`
}

type BackTableBody struct {
	TableID    int `json:"table_id"`
	BackerID   int `json:"backer"`
	MaterialID int `json:"material_id"`
	Num        int `json:"num"`
}

type CheckTableBody struct {
	TableID    int `json:"table_id"`
	CheckerID  int `json:"checker"`
	MaterialID int `json:"material_id"`
	Num        int `json:"num"`
}

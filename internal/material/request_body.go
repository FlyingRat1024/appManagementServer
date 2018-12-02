package material

type ApplyTableBody struct {
	ApplierID int      `json:"applier"`
	Material  []Common `json:"material"`
}

type RecieveTableBody struct {
	ReciiverID int      `json:"receiver"`
	Material   []Common `json:"material"`
}

type BackTableBody struct {
	TableID  int      `json:"table_id"`
	BackerID int      `json:"backer"`
	Material []Common `json:"material"`
}

type CheckTableBody struct {
	TableID   int      `json:"table_id"`
	CheckerID int      `json:"checker"`
	Material  []Common `json:"material"`
}

type Common struct {
	MaterialID int `json:"material_id"`
	Num        int `json:"num"`
}

type Material struct {
	Name        string `json:"name"`
	Unit        string `json:"unit"`
	Provider    string `json:"provider"`
	Description string `json:"description"`
}

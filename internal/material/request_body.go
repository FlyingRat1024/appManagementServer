package material

type ApplyTableBody struct {
	ProjectName string   `json:"project_name"`
	TableNum    string   `json:"table_num"`
	ApplierID   int      `json:"applier"`
	Material    []Common `json:"material"`
}

type RecieveTableBody struct {
	ProjectName string   `json:"project_name"`
	TableNum    string   `json:"table_num"`
	ReciiverID  int      `json:"receiver"`
	Material    []Common `json:"material"`
}

type BackTableBody struct {
	TableNum string   `json:"table_num"`
	TableID  int      `json:"table_id"`
	BackerID int      `json:"backer"`
	Material []Common `json:"material"`
}

type VerifyBody struct {
	TableID  int    `json:"table_id"`
	Verifier int    `json:"verifier"`
	Status   string `json:"status"`
}

type CheckTableBody struct {
	TableNum  string   `json:"table_num"`
	TableID   int      `json:"table_id"`
	CheckerID int      `json:"checker"`
	Material  []Common `json:"material"`
}

type Common struct {
	MaterialID int `json:"id"`
	Num        int `json:"num"`
}

type Material struct {
	Name        string `json:"name"`
	Unit        string `json:"unit"`
	Size        string `json:"size"`
	Provider    string `json:"provider"`
	Description string `json:"description"`
}

package structs

type Material struct {
	ID          string `json:"material"`
	Name        string `json:"name"`
	Unit        string `json:"unit"`
	Provider    string `json:"provider"`
	Description string `json:"description"`
}

type Verify struct {
	Verifier   string `json:"verifier"`
	Verify     string `json:"verify"`
	VerifyTime string `json:"verify_time"`
}

type User struct {
	ID           string
	EmployeeID   string
	EmployeeName string
	Sex          string
	Position     string
	CreateTime   string
}

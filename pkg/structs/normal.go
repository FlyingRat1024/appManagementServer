package structs

type Material struct {
	ID          int    `json:"material"`
	Name        string `json:"name"`
	Unit        string `json:"unit"`
	Provider    string `json:"provider"`
	Description string `json:"description"`
}

type Verify struct {
	Verifier   int    `json:"verifier"`
	Verify     string `json:"verify"`
	VerifyTime string `json:"verify_time"`
}

type User struct {
	ID           int
	EmployeeID   string
	EmployeeName string
	Sex          string
	Position     string
	CreateTime   string
}

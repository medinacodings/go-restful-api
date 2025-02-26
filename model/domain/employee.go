package domain

type Employee struct {
	EmployeeID string `json:"employee_id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
}

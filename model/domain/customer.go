package domain

type Customer struct {
	CustomerID string `json:"customer_id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	Loyaltyts  int    `json:"loyaltyts"`
}

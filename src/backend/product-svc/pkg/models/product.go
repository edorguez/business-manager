package models

type Product struct {
	Id          int64  `json:"id"`
	CompanyId   int64  `json:"companyId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Sku         string `json:"sku"`
	Price       uint64 `json:"price"`
}

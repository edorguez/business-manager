package contracts

type GetProductResponse struct {
	Id            string   `json:"id"`
	CompanyId     int64    `json:"companyId"`
	Name          string   `json:"name"`
	Description   *string  `json:"description"`
	Sku           *string  `json:"sku"`
	Quantity      uint64   `json:"quantity"`
	Price         uint64   `json:"price"`
	Images        []string `json:"images"`
	ProductStatus uint32   `json:"productStatus"`
}

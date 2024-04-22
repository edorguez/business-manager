package models

import "time"

type Product struct {
	CompanyId   int64     `json:"companyId" bson:"companyId"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Sku         string    `json:"sku" bson:"sku"`
	Price       uint64    `json:"price" bson:"price"`
	Images      []string  `json:"images" bson:"images"`
	Status      uint32    `json:"status" bson:"status"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	ModifiedAt  time.Time `json:"modifiedAt" bson:"modifiedAt"`
}

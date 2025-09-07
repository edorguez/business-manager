package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	CompanyId     int64     `json:"companyId" bson:"companyId"`
	Name          string    `json:"name" bson:"name"`
	Description   *string   `json:"description" bson:"description"`
	Sku           *string   `json:"sku" bson:"sku"`
	Quantity      uint64    `json:"quantity" bson:"quantity"`
	Price         uint64    `json:"price" bson:"price"`
	Images        []string  `json:"images" bson:"images"`
	ProductStatus uint32    `json:"productStatus" bson:"productStatus"`
	CreatedAt     time.Time `json:"createdAt" bson:"createdAt"`
	ModifiedAt    time.Time `json:"modifiedAt" bson:"modifiedAt"`
}

type GetProduct struct {
	Id            primitive.ObjectID `json:"_id" bson:"_id"`
	CompanyId     int64              `json:"companyId" bson:"companyId"`
	Name          string             `json:"name" bson:"name"`
	Description   *string            `json:"description" bson:"description"`
	Sku           *string            `json:"sku" bson:"sku"`
	Quantity      uint64             `json:"quantity" bson:"quantity"`
	Price         uint64             `json:"price" bson:"price"`
	Images        []string           `json:"images" bson:"images"`
	ProductStatus uint32             `json:"productStatus" bson:"productStatus"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
	ModifiedAt    time.Time          `json:"modifiedAt" bson:"modifiedAt"`
}

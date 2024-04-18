package repository

import "go.mongodb.org/mongo-driver/mongo"

type ProductRepo struct {
	client *mongo.Client
}

func NewProductRepo(client *mongo.Client) *ProductRepo {
	return &ProductRepo{
		client: client,
	}
}

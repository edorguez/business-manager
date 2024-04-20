package repository

import (
	"context"

	"github.com/EdoRguez/business-manager/product-svc/pkg/config"
	"github.com/EdoRguez/business-manager/product-svc/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepo struct {
	client *mongo.Client
	config config.Config
}

var collectionName = "products"

func NewProductRepo(client *mongo.Client, config config.Config) *ProductRepo {
	return &ProductRepo{
		client: client,
		config: config,
	}
}

func (productRepo *ProductRepo) CreateProduct(ctx context.Context, arg models.Product) (*models.Product, error) {
	collection := productRepo.client.Database(productRepo.config.DBName).Collection(collectionName)

	res, err := collection.InsertOne(ctx, arg)

	if err != mongo.ErrNilCursor {
		return nil, err
	}

	if _, ok := res.InsertedID.(primitive.ObjectID); ok {
		return &arg, nil
	}
	return &arg, err
}

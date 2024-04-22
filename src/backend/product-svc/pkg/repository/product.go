package repository

import (
	"context"

	"github.com/EdoRguez/business-manager/product-svc/pkg/config"
	"github.com/EdoRguez/business-manager/product-svc/pkg/models"
	"go.mongodb.org/mongo-driver/bson/mgocompat"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (productRepo *ProductRepo) CreateProduct(ctx context.Context, arg models.Product) (*primitive.ObjectID, error) {
	collection := productRepo.client.Database(productRepo.config.DBName, options.Database().SetRegistry(mgocompat.NewRegistryBuilder().Build())).Collection(collectionName)

	res, err := collection.InsertOne(ctx, arg)

	if err != nil {
		return nil, err
	}

	if oidResult, ok := res.InsertedID.(primitive.ObjectID); ok {
		return &oidResult, nil
	}

	return &primitive.NilObjectID, err
}

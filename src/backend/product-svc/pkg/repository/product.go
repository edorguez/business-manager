package repository

import (
	"context"
	"time"

	"github.com/EdoRguez/business-manager/product-svc/pkg/config"
	"github.com/EdoRguez/business-manager/product-svc/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
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
	collection := productRepo.client.Database(productRepo.config.DBName).Collection(collectionName)

	res, err := collection.InsertOne(ctx, arg)

	if err != nil {
		return nil, err
	}

	if oidResult, ok := res.InsertedID.(primitive.ObjectID); ok {
		return &oidResult, nil
	}

	return &primitive.NilObjectID, err
}

func (productRepo *ProductRepo) GetProduct(ctx context.Context, id primitive.ObjectID) (*models.GetProduct, error) {
	collection := productRepo.client.Database(productRepo.config.DBName).Collection(collectionName)

	var result models.GetProduct
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type GetProductsParams struct {
	CompanyId int64
	Limit     int32
	Offset    int32
}

func (productRepo *ProductRepo) GetProducts(ctx context.Context, arg GetProductsParams) ([]models.GetProduct, error) {
	collection := productRepo.client.Database(productRepo.config.DBName).Collection(collectionName)

	var result []models.GetProduct
	options := options.Find().SetLimit(int64(arg.Limit)).SetSkip(int64(arg.Offset))
	cursor, err := collection.Find(ctx, bson.M{"companyId": arg.CompanyId}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(context.Background()) {
		var product models.GetProduct
		cursor.Decode(&product)
		result = append(result, product)
	}

	return result, nil
}

func (productRepo *ProductRepo) UpdateProduct(ctx context.Context, id primitive.ObjectID, arg models.Product) error {
	collection := productRepo.client.Database(productRepo.config.DBName).Collection(collectionName)

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"name":          arg.Name,
		"description":   arg.Description,
		"sku":           arg.Sku,
		"price":         arg.Price,
		"images":        arg.Images,
		"productStatus": arg.ProductStatus,
		"modifiedAt":    time.Now(),
	},
	}

	_, err := collection.UpdateOne(ctx, filter, update, opts)
	return err
}

func (productRepo *ProductRepo) DeleteProduct(ctx context.Context, id primitive.ObjectID) error {
	collection := productRepo.client.Database(productRepo.config.DBName).Collection(collectionName)

	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(ctx, filter)
	return err
}

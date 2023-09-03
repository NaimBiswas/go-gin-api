package productServices

import (
	Models "NaimBiswas/go-gin-api/models"
	CommonServices "NaimBiswas/go-gin-api/services"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllProduct(limit, page int, collection *mongo.Collection) ([]Models.ProductModel, int64, error) {
	dataCursor, dataCount, err := CommonServices.GetValues(collection, limit, page)
	if err != nil {
		return nil, 0, err
	}
	var products []Models.ProductModel
	defer dataCursor.Close(context.Background())
	err = dataCursor.All(context.Background(), &products)

	if err != nil {
		return nil, 0, err
	}
	return products, dataCount, nil
}

func GetAllImportedProducts(limit int, page int, collection *mongo.Collection) ([]Models.ProductModel, int64, error) {
	dataCursor, dataCount, err := CommonServices.GetValues(collection, limit, page)
	if err != nil {
		return nil, 0, err
	}
	var products []Models.ProductModel
	defer dataCursor.Close(context.Background())
	err = dataCursor.All(context.Background(), &products)

	if err != nil {
		return nil, 0, err
	}
	return products, dataCount, nil
	return nil, 0, nil
}

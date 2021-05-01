package service

import (
	"context"
	"github.com/maulanay85/go-micro-product/pkg/config"
	"github.com/maulanay85/go-micro-product/pkg/db"
	"github.com/maulanay85/go-micro-product/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func GetAllCategory(opt *util.Paging) (data *[]bson.M, err error) {
	database, err := config.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer database.Client().Disconnect(context.TODO())
	opts := options.Find()
	opts.SetSkip(int64(opt.Skip))
	opts.SetLimit(int64(opt.Limit))
	bsonM := util.Filter(opt.Filter)
	cursor, err := database.Collection("categories").Find(context.TODO(), bsonM, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	var tmp []bson.M
	if err = cursor.All(context.TODO(), &tmp); err != nil {
		return nil, err
	}
	return &tmp, nil
}

func CreateCategory(data *db.Category) (*mongo.InsertOneResult, error) {
	database, err := config.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer database.Client().Disconnect(context.Background())

	result, err := database.Collection("categories").InsertOne(context.Background(), &data)
	return result, err
}

func FindByIdCategory(id string) (*db.Category, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	database, err := config.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer database.Client().Disconnect(context.Background())
	category := db.Category{}
	findOneResult := database.Collection("categories").FindOne(context.Background(), bson.M{"_id": objId})
	if err := findOneResult.Decode(&category); err != nil {
		return nil, err
	}
	return &category, nil
}

func DeleteCategoryById(id string) (*mongo.UpdateResult, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	database, err := config.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer database.Client().Disconnect(context.Background())
	var category bson.M
	findOneResult := database.Collection("categories").FindOne(context.Background(), bson.M{"_id": objId})
	if err := findOneResult.Decode(&category); err != nil {
		return nil, err
	}
	t, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	result, err := database.Collection("categories").UpdateOne(context.Background(), bson.M{"_id": objId},
		bson.D{
			{"$set", bson.D{{"isDeleted", true}}},
			{"$set", bson.D{{"updatedAt", t}}},
		})
	return result, err
}

func UpdateCategoryById(id string, category *db.Category) (*mongo.UpdateResult, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	database, err := config.Connect()
	if err != nil {
		return nil, err
	}
	defer database.Client().Disconnect(context.Background())
	var tmp bson.M
	findOneResult := database.Collection("categories").FindOne(context.Background(), bson.M{"_id": objId})
	if err := findOneResult.Decode(&tmp); err != nil {
		return nil, err
	}
	t, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	result, err := database.Collection("categories").UpdateOne(context.Background(), bson.M{"_id": objId},
		bson.D{
			{"$set", bson.D{{"categoryName", &category.CategoryName}}},
			{"$set", bson.D{{"categoryCode", &category.CategoryCode}}},
			{"$set", bson.D{{"updatedAt", t}}},
		})
	return result, err
}

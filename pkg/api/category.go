package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/maulanay85/go-micro-product/pkg/config"
	"github.com/maulanay85/go-micro-product/pkg/db"
	"github.com/maulanay85/go-micro-product/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
	//"time"
)

func CreateCategory(c *gin.Context) {
	var input db.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, config.ResponseBadRequest(err.Error()))
		return

	}
	t, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	data := db.Category{
		ID:           primitive.NewObjectID(),
		CategoryCode: input.CategoryCode,
		CategoryName: input.CategoryName,
		IsDeleted:    false,
		CreatedAt:    t,
		UpdatedAt:    t,
	}
	database, err := config.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer database.Client().Disconnect(context.TODO())

	result, err := database.Collection("categories").InsertOne(context.TODO(), &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err.Error()})
		return
	}
	c.JSON(http.StatusOK, config.ResponseSuccess(result))
}

func GetAll(c *gin.Context) {
	pageReq := c.DefaultQuery("page", "1")
	limitReq := c.DefaultQuery("limit", "30")
	skip, limit := util.Pagination(pageReq, limitReq)
	filter := make(map[string]interface{})
	categoryName := c.Query("categoryName")
	if categoryName != "" {
		filter["categoryName"] = categoryName
	}
	filter["isDeleted"] = false
	database, err := config.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer database.Client().Disconnect(context.TODO())
	opts := options.Find()
	opts.SetSkip(int64(skip))
	opts.SetLimit(int64(limit))
	bsonM := util.Filter(filter)
	cursor, err := database.Collection("categories").Find(context.TODO(), bsonM, opts)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer cursor.Close(context.TODO())
	//fmt.Println(cursor)
	var tmp []bson.M
	if err = cursor.All(context.TODO(), &tmp); err != nil {
		log.Fatal(err.Error())

	}
	c.JSON(http.StatusOK, config.ResponseSuccess(tmp))

}

func FindById(c *gin.Context) {
	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)
	database, err := config.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer database.Client().Disconnect(context.Background())
	var category bson.M
	findOneResult := database.Collection("categories").FindOne(context.Background(), bson.M{"_id": objId})
	if err := findOneResult.Decode(&category); err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, config.ResponseNotFound(""))
			return
		} else {
			c.JSON(http.StatusInternalServerError, config.ResponseServerError(err.Error()))
			return
		}
	}
	c.JSON(http.StatusOK, config.ResponseSuccess(category))
}

func UpdateById(c *gin.Context) {
	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)
	var input db.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, config.ResponseBadRequest(err.Error()))
		return

	}
	database, err := config.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer database.Client().Disconnect(context.Background())
	var category bson.M
	findOneResult := database.Collection("categories").FindOne(context.Background(), bson.M{"_id": objId})
	if err := findOneResult.Decode(&category); err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, config.ResponseNotFound(""))
			return
		} else {
			c.JSON(http.StatusInternalServerError, config.ResponseServerError(err.Error()))
			return
		}
	}
	t, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	result, err := database.Collection("categories").UpdateOne(context.Background(), bson.M{"_id": objId},
		bson.D{
			{"$set", bson.D{{"categoryName", input.CategoryName}}},
			{"$set", bson.D{{"categoryCode", input.CategoryCode}}},
			{"$set", bson.D{{"updatedAt", t}}},
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, config.ResponseServerError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, config.ResponseSuccess(result))
}

func DeleteById(c *gin.Context) {
	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)
	database, err := config.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer database.Client().Disconnect(context.Background())
	var category bson.M
	findOneResult := database.Collection("categories").FindOne(context.Background(), bson.M{"_id": objId})
	if err := findOneResult.Decode(&category); err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, config.ResponseNotFound(""))
			return
		} else {
			c.JSON(http.StatusInternalServerError, config.ResponseServerError(err.Error()))
			return
		}
	}
	t, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	result, err := database.Collection("categories").UpdateOne(context.Background(), bson.M{"_id": objId},
		bson.D{
			{"$set", bson.D{{"isDeleted", true}}},
			{"$set", bson.D{{"updatedAt", t}}},
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, config.ResponseServerError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, config.ResponseSuccess(result))
}

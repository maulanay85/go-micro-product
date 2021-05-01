package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maulanay85/go-micro-product/pkg/config"
	"github.com/maulanay85/go-micro-product/pkg/db"
	"github.com/maulanay85/go-micro-product/pkg/service"
	"github.com/maulanay85/go-micro-product/pkg/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	result, err := service.CreateCategory(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, config.ResponseServerError(""))
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
	opt := util.Paging{Skip: skip, Limit: limit, Filter: filter}
	if data, err := service.GetAllCategory(&opt); err != nil {
		c.JSON(http.StatusInternalServerError, config.ResponseServerError(""))
		return
	} else {
		c.JSON(http.StatusOK, config.ResponseSuccess(data))
	}
}

func FindById(c *gin.Context) {
	id := c.Param("id")
	data, err := service.FindByIdCategory(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, config.ResponseNotFound(""))
			return
		} else {
			c.JSON(http.StatusInternalServerError, config.ResponseServerError(err.Error()))
			return
		}
	}
	c.JSON(http.StatusOK, config.ResponseSuccess(data))
}

func UpdateById(c *gin.Context) {
	id := c.Param("id")
	var input db.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, config.ResponseBadRequest(err.Error()))
		return

	}
	result, err := service.UpdateCategoryById(id, &input)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, config.ResponseNotFound(""))
			return
		} else {
			c.JSON(http.StatusInternalServerError, config.ResponseServerError(err.Error()))
			return
		}
	}
	c.JSON(http.StatusOK, config.ResponseSuccess(result))
}

func DeleteById(c *gin.Context) {
	id := c.Param("id")
	result, err := service.DeleteCategoryById(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, config.ResponseNotFound(""))
			return
		} else {
			c.JSON(http.StatusInternalServerError, config.ResponseServerError(err.Error()))
			return
		}
	}
	c.JSON(http.StatusOK, config.ResponseSuccess(result))
}

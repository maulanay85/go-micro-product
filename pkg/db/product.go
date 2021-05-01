package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Product struct {
	ID                 primitive.ObjectID `bson:"_id"`
	ProductName        string             `bson:"productName" binding:"required"`
	ProductDescription string             `bson:"productDescription" binding:"required"`
	ProductPrice       float64            `bson:"productPrice" binding:"required"`
	ProductStock       float64            `bson:"productStock"`
	CreatedAt          time.Time          `bson:"createdAt"`
	UpdatedAt          time.Time          `bson:"updatedAt"`
	IsDeleted          bool               `bson:"isDeleted"`
	Category           Category           `bson:"category" binding:"required"`
	MerchantId         int32              `bson:"merchantId" binding:"required"`
	MerchantName       string             `bson:"merchantName" binding:"required"`
	HasSell            int64              `bson:"hasSell"`
}

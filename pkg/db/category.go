package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Category struct {
	ID           primitive.ObjectID `bson:"_id",omitempty`
	CategoryName string             `bson:"categoryName" binding:"required"`
	CategoryCode string             `bson:"categoryCode" binding:"required"`
	CreatedAt    time.Time          `bson:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt"`
	IsDeleted    bool               `bson:"isDeleted"`
}

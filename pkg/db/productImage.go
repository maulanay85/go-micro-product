package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ProductImage struct {
	ID         primitive.ObjectID `bson:"_id"`
	Url        string             `bson:"url"`
	ImageName  string             `bson:"imageName"`
	MerchantId primitive.ObjectID `bson:"merchantId"`
	CreatedAt  time.Time          `bson:"createdAt"`
}

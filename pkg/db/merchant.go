package db

import "time"

type Merchant struct {
	MerchantName    string    `bson:"merchantName"`
	MerchantLogo    string    `bson:"merchantLogo"`
	MerchantAddress string    `bson:"merchantAddress"`
	Status          int32     `bson:"status"`
	CreatedAt       time.Time `bson:"createdAt"`
	UpdatedAt       time.Time `bson:"updatedAt"`
	OwnerId         int64     `bson:"ownerId"`
}

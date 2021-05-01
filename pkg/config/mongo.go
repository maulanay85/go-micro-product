package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

//const urlString string

func Connect() (*mongo.Database, error) {
	urlString := GetString("database.url")
	client, err := mongo.NewClient(options.Client().ApplyURI(urlString))
	if err != nil {
		return nil, err
	}

	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	//defer client.Disconnect(context.TODO())
	return client.Database("product-micro"), nil

}

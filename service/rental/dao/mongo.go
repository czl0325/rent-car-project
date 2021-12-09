package dao

import "go.mongodb.org/mongo-driver/mongo"

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("rental"),
	}
}


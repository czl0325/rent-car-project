package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017/rentcar?readPreference=primary&ssl=false"))
	if err != nil {
		panic(err)
	}
	defer mc.Disconnect(c)
	db := mc.Database("rentcar").Collection("account")
	findRow(c, db)
}

func InsertData(c context.Context, db *mongo.Collection)  {
	res, err := db.InsertMany(c, []interface{}{
		bson.M{
			"phone": "111",
		},
		bson.M{
			"phone": "222",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", res)
}

func findRow(c context.Context, db *mongo.Collection) {
	cur, err := db.Find(c, bson.M{})
	if err != nil {
		panic(err)
	}
	for cur.Next(c) {
		var row struct {
			ID    primitive.ObjectID `bson:"_id"`
			Phone string             `bson:"phone"`
		}
		err := cur.Decode(&row)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", row)
	}
}

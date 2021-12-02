package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	authpb "service/auth/api/gen/v1"
	"service/shared/token"
)

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("account"),
	}
}

func (m *Mongo) LoginWithRegister(c context.Context, user *authpb.LoginRequest) (*authpb.UserInfo, error) {
	res := m.col.FindOne(c, bson.M{
		"phone": user.Phone,
	})
	var row struct {
		Id       primitive.ObjectID `bson:"_id"`
		Phone    string             `bson:"phone"`
		Password string             `bson:"password"`
	}
	err := res.Decode(&row)
	if err != nil {
		insert, err := m.col.InsertOne(c, bson.M{
			"phone":    user.Phone,
			"password": user.Password,
		})
		if err != nil {
			return nil, err
		}
		tk, err := token.GenerateToken(row.Phone)
		if err != nil {
			return nil, err
		}
		newUser := &authpb.UserInfo{Id: insert.InsertedID.(primitive.ObjectID).Hex(), Phone: row.Phone, Token: tk}
		return newUser, nil
	} else {
		if row.Password != user.Password {
			return nil, fmt.Errorf("密码错误")
		}
		tk, err := token.GenerateToken(row.Phone)
		if err != nil {
			return nil, err
		}
		newUser := &authpb.UserInfo{Id: row.Id.Hex(), Phone: row.Phone, Token: tk}
		return newUser, nil
	}
}

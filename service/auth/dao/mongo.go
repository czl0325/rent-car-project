package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	authpb "service/auth/api/gen/v1"
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
	res := m.col.FindOneAndUpdate(c, bson.M{
		"phone": user.Phone,
	}, bson.M{
		"$set": bson.M{
			"phone":    user.Phone,
			"password": user.Password,
		},
	}, options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After))
	if res.Err() != nil {
		return nil, fmt.Errorf("登录注册失败,错误=%+v\n", res.Err())
	}
	var row struct {
		Id    primitive.ObjectID `bson:"_id"`
		Phone string             `bson:"phone"`
	}
	err := res.Decode(&row)
	if err != nil {
		return nil, fmt.Errorf("解析数据库错误,错误=%+v\n", err)
	}
	newUser := &authpb.UserInfo{Id:row.Id.String(), Phone:row.Phone}
	return newUser, nil
}

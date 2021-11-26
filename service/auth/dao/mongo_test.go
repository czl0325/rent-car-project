package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	authpb "service/auth/api/gen/v1"
	"testing"
)

func TestLoginWithRegister(t *testing.T) {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017/rentcar?readPreference=primary&ssl=false"))
	if err != nil {
		t.Errorf("连接mongodb失败，错误=%v\n", err)
	}
	m := NewMongo(mc.Database("rentcar"))
	res, err := m.LoginWithRegister(c, &authpb.LoginRequest{
		Phone:    "123",
		Password: "123",
	})
	if err != nil {
		t.Errorf("LoginWithRegister执行失败，错误=%v\n", err)
	}
	fmt.Printf("返回数据=%+v\n", res)
}

package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	authpb "service/auth/api/gen/v1"
	mongo_testing "service/shared/mongo/testing"
	"testing"
)

func TestLoginWithRegister(t *testing.T) {
	c := context.Background()
	mc, err := mongo_testing.NewClient(c)
	if err != nil {
		t.Errorf("连接mongodb失败，错误=%v\n", err)
	}
	m := NewMongo(mc.Database("rentcar"))
	_, err = m.col.InsertOne(c, bson.M{
		"phone":    "123",
		"password": "123",
	})
	if err != nil {
		t.Errorf("插入数据错误，错误=%+v\n", err)
	}
	user, err := m.LoginWithRegister(c, &authpb.LoginRequest{
		Phone:    "123",
		Password: "123",
	})
	if err != nil {
		t.Errorf("LoginWithRegister错误，错误=%+v\n", err)
	}
	fmt.Printf("插入后的数据=%+v\n", user)
	if user.Phone != "123" {
		t.Errorf("测试失败")
	}
}

func TestMain(m *testing.M) {
	os.Exit(mongo_testing.RunWithMongoInDocker(m))
}

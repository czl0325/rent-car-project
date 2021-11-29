package testing

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

const (
	image         = "mongo"
	containerPort = "27017/tcp"
)
var mongoURI string

func RunWithMongoInDocker(m *testing.M) int {
	c, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	res, err := c.ContainerCreate(ctx, &container.Config{
		Image: image,
		ExposedPorts: nat.PortSet{
			containerPort: {},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			containerPort: []nat.PortBinding {
				{
					HostIP: "127.0.0.1",
					HostPort: "0",
				},
			},
		},
	}, nil, nil, "")
	containerId := res.ID
	defer func() {
		err := c.ContainerRemove(ctx, containerId, types.ContainerRemoveOptions{Force: true})
		if err != nil {
			panic(err)
		}
	}()
	err = c.ContainerStart(ctx, containerId, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}
	inspectRes, err := c.ContainerInspect(ctx, containerId)
	if err != nil {
		panic(err)
	}
	host := inspectRes.NetworkSettings.Ports[containerPort][0]
	mongoURI = fmt.Sprintf("mongodb://%s:%s", host.HostIP, host.HostPort)
	return m.Run()
}

func NewClient(c context.Context) (*mongo.Client, error)  {
	if mongoURI == "" {
		return nil, fmt.Errorf("mongodb的URL为空，请补充TestMain函数")
	}
	return mongo.Connect(c, options.Client().ApplyURI(mongoURI))
}


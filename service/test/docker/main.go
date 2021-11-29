package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"time"
)

func main() {
	c, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	res, err := c.ContainerCreate(ctx, &container.Config{
		Image: "mongo",
		ExposedPorts: nat.PortSet{
			"27017/tcp": {},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			"27017/tcp": []nat.PortBinding{
				{
					HostIP:   "127.0.0.1",
					HostPort: "0",
				},
			},
		},
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}
	err = c.ContainerStart(ctx, res.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("容器启动")
	my, err := c.ContainerInspect(ctx, res.ID)
	fmt.Printf("监听开始，%+v\n", my.NetworkSettings.Ports["27017/tcp"][0])
	time.Sleep(10 * time.Second)
	err = c.ContainerRemove(ctx, res.ID, types.ContainerRemoveOptions{Force: true})
	if err != nil {
		panic(err)
	}
	fmt.Println("容器停止并删除")
}
